#!/bin/bash
# ============================================================
# Cosmic Signature FAQ Bot — EC2 Auto-Setup Script
# Run this on a fresh Ubuntu 22.04 EC2 instance
# Usage: bash setup_ec2.sh
# ============================================================

set -e  # Exit on any error

# ── Auto-detect project directory ────────────────────────
# The script looks for the project relative to its own location
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

echo "=================================================="
echo "  Cosmic Signature FAQ Bot — EC2 Setup"
echo "  Project directory: $PROJECT_DIR"
echo "=================================================="
echo ""

# ── 1. System dependencies ───────────────────────────────
echo "[1/8] Installing system dependencies..."
sudo apt update -y

# Detect available Python 3 version
if apt-cache show python3.12 &>/dev/null; then
    PY_PKG="python3.12 python3.12-venv"
    PY_BIN="python3.12"
elif apt-cache show python3.11 &>/dev/null; then
    PY_PKG="python3.11 python3.11-venv"
    PY_BIN="python3.11"
else
    PY_PKG="python3 python3-venv"
    PY_BIN="python3"
fi
echo "  Using Python: $PY_PKG"
sudo apt install -y git curl build-essential $PY_PKG python3-pip nginx

# ── 2. Node.js via NodeSource (more reliable in scripts than nvm) ────────────
echo "[2/8] Installing Node.js 20..."
if ! command -v node &>/dev/null; then
    curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
    sudo apt install -y nodejs
else
    echo "  Node.js already installed: $(node --version)"
fi
echo "  Node: $(node --version), npm: $(npm --version)"

# ── 3. PM2 ────────────────────────────────────────────────
echo "[3/8] Installing PM2..."
npm install -g pm2 2>/dev/null || true

# ── 4. Python backend ─────────────────────────────────────
echo "[4/8] Setting up Python backend..."
cd "$PROJECT_DIR/backend"
$PY_BIN -m venv venv
source venv/bin/activate
pip install --upgrade pip -q
pip install -r requirements.txt -q
deactivate

# ── 5. .env files ─────────────────────────────────────────
echo "[5/8] Setting up environment files..."

# Backend .env
if [ ! -f "$PROJECT_DIR/backend/.env" ]; then
    cp "$PROJECT_DIR/backend/.env.example" "$PROJECT_DIR/backend/.env"
    echo "  Created backend/.env — please add your GEMINI_API_KEY!"
fi

# Frontend .env.local
PUBLIC_IP=$(curl -s http://checkip.amazonaws.com/ 2>/dev/null || echo "localhost")
echo "NEXT_PUBLIC_BACKEND_URL=http://${PUBLIC_IP}" > "$PROJECT_DIR/.env.local"
echo "  Frontend will connect to: http://${PUBLIC_IP}"

# ── 6. Build frontend ─────────────────────────────────────
echo "[6/8] Building Next.js frontend..."
cd "$PROJECT_DIR"
npm install -q
npm run build

# ── 7. PM2 ecosystem config ───────────────────────────────
echo "[7/8] Creating PM2 config..."
BACKEND_PYTHON="$PROJECT_DIR/backend/venv/bin/python"

cat > "$PROJECT_DIR/ecosystem.config.js" << PMEOF
module.exports = {
  apps: [
    {
      name: 'faq-backend',
      cwd: '${PROJECT_DIR}/backend',
      script: '${BACKEND_PYTHON}',
      args: 'app.py',
      interpreter: 'none',
      env: {
        PATH: '${PROJECT_DIR}/backend/venv/bin:' + process.env.PATH,
      },
      watch: false,
      autorestart: true,
      max_restarts: 5,
    },
    {
      name: 'faq-frontend',
      cwd: '${PROJECT_DIR}',
      script: 'node_modules/.bin/next',
      args: 'start -p 3000',
      watch: false,
      autorestart: true,
      max_restarts: 5,
    },
  ],
}
PMEOF

# Start with PM2
pm2 start "$PROJECT_DIR/ecosystem.config.js"
pm2 save
pm2 startup systemd -u $USER --hp $HOME | tail -1 | sudo bash || true

# ── 8. Nginx config ───────────────────────────────────────
echo "[8/8] Configuring Nginx..."
sudo tee /etc/nginx/sites-available/faq > /dev/null << NGEOF
server {
    listen 80;
    server_name _;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_cache_bypass \$http_upgrade;
    }

    location /api/ {
        proxy_pass http://localhost:8000/api/;
        proxy_http_version 1.1;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_read_timeout 120s;
    }

    location /health {
        proxy_pass http://localhost:8000/health;
    }

    location /debug/ {
        proxy_pass http://localhost:8000/debug/;
    }
}
NGEOF

sudo rm -f /etc/nginx/sites-enabled/default
sudo ln -sf /etc/nginx/sites-available/faq /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl restart nginx && sudo systemctl enable nginx

echo ""
echo "=================================================="
echo "  Setup Complete!"
echo "=================================================="
echo ""
echo "  Public IP:    http://${PUBLIC_IP}"
echo "  Health check: http://${PUBLIC_IP}/health"
echo "  Debug stats:  http://${PUBLIC_IP}/debug/stats"
echo ""
echo "  IMPORTANT: Add your Gemini API key:"
echo "    nano $PROJECT_DIR/backend/.env"
echo "    → Set GEMINI_API_KEY=your_key"
echo "    → Then: pm2 restart faq-backend"
echo ""
echo "  PM2 commands:"
echo "    pm2 status"
echo "    pm2 logs faq-backend"
echo "    pm2 logs faq-frontend"
echo ""
echo "  NOTE: First startup will take 5-10 minutes"
echo "  (cloning and indexing GitHub repositories)"
echo "=================================================="
