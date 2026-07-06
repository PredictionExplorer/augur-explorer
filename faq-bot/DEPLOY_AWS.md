# AWS EC2 Deployment Guide

## Architecture

```
Internet
   │
   ▼
EC2 Instance (Ubuntu 22.04)
   │
   ├── Nginx (port 80/443) ← reverse proxy
   │     ├── / → Next.js frontend (port 3000)
   │     └── /api → FastAPI backend (port 8000)
   │
   ├── Next.js (PM2, port 3000)
   └── FastAPI (PM2, port 8000)
```

---

## Step 1 – Launch EC2 Instance

1. Go to **AWS Console → EC2 → Launch Instance**
2. Choose:
   - **AMI:** Ubuntu Server 22.04 LTS
   - **Instance type:** `t3.medium` (2 vCPU, 4GB RAM — minimum for indexing repos)
   - **Storage:** 20GB+ (repos + dependencies)
3. **Key pair:** Create or select an existing `.pem` key
4. **Security Group — open these ports:**

| Port | Protocol | Source    | Purpose         |
|------|----------|-----------|-----------------|
| 22   | TCP      | Your IP   | SSH access      |
| 80   | TCP      | 0.0.0.0/0 | HTTP            |
| 443  | TCP      | 0.0.0.0/0 | HTTPS           |

> Do NOT expose ports 3000 or 8000 publicly — Nginx will handle routing.

5. Click **Launch Instance**

---

## Step 2 – Connect to the Instance

```bash
# Replace with your .pem file path and EC2 public IP
chmod 400 your-key.pem
ssh -i your-key.pem ubuntu@YOUR_EC2_PUBLIC_IP
```

---

## Step 3 – Install System Dependencies

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Git
sudo apt install -y git curl build-essential

# Install Python 3.10 + venv
sudo apt install -y python3.10 python3.10-venv python3-pip

# Install Node.js 20 via nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
source ~/.bashrc
nvm install 20
nvm use 20
node --version   # should print v20.x.x

# Install PM2 (process manager — keeps app running)
npm install -g pm2

# Install Nginx
sudo apt install -y nginx
```

---

## Step 4 – Clone & Configure the Project

```bash
# Clone your repo (or upload files via scp)
git clone https://github.com/YOUR_USERNAME/YOUR_REPO.git
cd YOUR_REPO

# Or upload from local machine:
# scp -i your-key.pem -r ./faq ubuntu@YOUR_EC2_PUBLIC_IP:~/faq
```

---

## Step 5 – Set Up the Backend

```bash
cd ~/faq/backend

# Create virtual environment
python3 -m venv venv
source venv/bin/activate

# Install dependencies
pip install --upgrade pip
pip install -r requirements.txt

# Create .env file
cp .env.example .env
nano .env
```

Edit `.env` with your values:

```env
PORT=8000
HOST=0.0.0.0
DATA_DIR=./data
LOG_LEVEL=INFO

# Add your LLM key (Gemini is free)
GEMINI_API_KEY=your_gemini_key_here
```

Test the backend runs:

```bash
source venv/bin/activate
python app.py
# Press Ctrl+C after confirming it starts
```

---

## Step 6 – Set Up the Frontend

```bash
cd ~/faq

# Create .env.local
cat > .env.local << 'EOF'
NEXT_PUBLIC_BACKEND_URL=http://YOUR_EC2_PUBLIC_IP/api-backend
EOF

# Install dependencies
npm install

# Build for production
npm run build
```

---

## Step 7 – Create PM2 Config

```bash
cd ~/faq
cat > ecosystem.config.js << 'EOF'
module.exports = {
  apps: [
    {
      name: 'faq-backend',
      cwd: '/home/ubuntu/faq/backend',
      script: '/home/ubuntu/faq/backend/venv/bin/python',
      args: 'app.py',
      interpreter: 'none',
      env: {
        PATH: '/home/ubuntu/faq/backend/venv/bin:' + process.env.PATH,
      },
      watch: false,
      autorestart: true,
      max_restarts: 5,
    },
    {
      name: 'faq-frontend',
      cwd: '/home/ubuntu/faq',
      script: 'node_modules/.bin/next',
      args: 'start -p 3000',
      watch: false,
      autorestart: true,
      max_restarts: 5,
    },
  ],
}
EOF
```

Start both services:

```bash
pm2 start ecosystem.config.js

# Check status
pm2 status

# Check logs
pm2 logs faq-backend
pm2 logs faq-frontend

# Save to restart on reboot
pm2 save
pm2 startup
# Run the command PM2 outputs
```

---

## Step 8 – Configure Nginx

```bash
sudo nano /etc/nginx/sites-available/faq
```

Paste this config (replace `YOUR_DOMAIN_OR_IP`):

```nginx
server {
    listen 80;
    server_name YOUR_DOMAIN_OR_IP;

    # Frontend
    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_cache_bypass $http_upgrade;
    }

    # Backend API
    location /api/ {
        proxy_pass http://localhost:8000/api/;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_read_timeout 120s;
    }

    # Backend health check
    location /health {
        proxy_pass http://localhost:8000/health;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
    }
}
```

Enable and restart:

```bash
sudo ln -s /etc/nginx/sites-available/faq /etc/nginx/sites-enabled/
sudo nginx -t          # test config — must say "ok"
sudo systemctl restart nginx
sudo systemctl enable nginx
```

---

## Step 9 – Update Frontend Backend URL

Since Nginx proxies `/api/` to the backend, update `.env.local`:

```bash
cd ~/faq
echo "NEXT_PUBLIC_BACKEND_URL=http://YOUR_EC2_PUBLIC_IP" > .env.local

# Rebuild the frontend
npm run build

# Restart PM2
pm2 restart faq-frontend
```

---

## Step 10 – Verify Everything Works

```bash
# Check PM2 processes
pm2 status

# Test backend health
curl http://localhost:8000/health

# Test Nginx routing
curl http://YOUR_EC2_PUBLIC_IP/health

# Check indexing status
curl http://YOUR_EC2_PUBLIC_IP/debug/stats
```

Open your browser: **http://YOUR_EC2_PUBLIC_IP**

You should see the FAQ bot chat interface!

---

## Step 11 – (Optional) Add a Domain + HTTPS

If you have a domain name:

```bash
# Point your domain's A record to the EC2 public IP
# Then install Certbot for free SSL

sudo apt install -y certbot python3-certbot-nginx
sudo certbot --nginx -d yourdomain.com

# Auto-renew SSL
sudo systemctl enable certbot.timer
```

Update Nginx `server_name`:

```nginx
server_name yourdomain.com www.yourdomain.com;
```

---

## Useful PM2 Commands

```bash
pm2 status                    # see all processes
pm2 logs faq-backend          # backend logs (live)
pm2 logs faq-frontend         # frontend logs (live)
pm2 restart faq-backend       # restart backend
pm2 restart faq-frontend      # restart frontend
pm2 restart all               # restart everything
pm2 stop all                  # stop everything
pm2 monit                     # live monitoring dashboard
```

## Useful Nginx Commands

```bash
sudo nginx -t                       # test config
sudo systemctl restart nginx        # restart
sudo tail -f /var/log/nginx/error.log   # error logs
sudo tail -f /var/log/nginx/access.log  # access logs
```

---

## Troubleshooting

### Backend won't start
```bash
pm2 logs faq-backend --lines 50
# Usually: missing package → pip install -r requirements.txt
```

### Frontend 502 Bad Gateway
```bash
pm2 status   # is faq-frontend running?
pm2 logs faq-frontend
```

### First run takes too long
The backend clones 3 GitHub repos and indexes ~3000 files on first start.
This can take **5-10 minutes** on a t3.medium. Monitor progress:
```bash
pm2 logs faq-backend
```

### Out of memory during indexing
Upgrade to `t3.large` (8GB RAM) or add swap:
```bash
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

### CORS errors in browser
The backend CORS config only allows `localhost`. Update `backend/app.py`:
```python
allow_origins=["http://YOUR_EC2_PUBLIC_IP", "https://yourdomain.com"]
```
Then `pm2 restart faq-backend`.

---

## Cost Estimate

| Resource       | Type        | Monthly Cost |
|----------------|-------------|--------------|
| EC2 instance   | t3.medium   | ~$30         |
| EC2 instance   | t3.small    | ~$15 (tight) |
| Storage (EBS)  | 20GB gp3    | ~$1.60       |
| Data transfer  | Outbound    | ~$1-5        |
| **Total**      |             | **~$17-36**  |

> 💡 Use a **t3.small** for low traffic. Upgrade to **t3.medium** if indexing is slow.
