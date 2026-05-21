# Setup Guide - Cosmic Signature FAQ Bot

This guide will help you set up the Cosmic Signature FAQ Bot on your local machine.

## Quick Start (Recommended)

### Step 1: Install Prerequisites

1. **Install Node.js** (version 18 or higher)
   - Download from: https://nodejs.org/
   - Verify installation: `node --version`

2. **Install Python** (version 3.9 or higher)
   - Download from: https://www.python.org/downloads/
   - Verify installation: `python --version` or `python3 --version`

3. **Install Git**
   - Download from: https://git-scm.com/
   - Verify installation: `git --version`

### Step 2: Setup Backend

**On Windows:**
```bash
cd backend
setup.bat
```

**On macOS/Linux:**
```bash
cd backend
chmod +x setup.sh
./setup.sh
```

This will:
- Create a Python virtual environment
- Install all required Python packages
- Create a `.env` file from the template

### Step 3: Setup Frontend

```bash
# In the root directory
npm install
cp .env.example .env
```

### Step 4: Run the Application

**Option A: Run manually in separate terminals**

Terminal 1 (Backend):
```bash
cd backend
# On Windows: venv\Scripts\activate
# On macOS/Linux: source venv/bin/activate
python app.py
```

Terminal 2 (Frontend):
```bash
npm run dev
```

**Option B: Use start scripts**

On Windows (Backend):
```bash
cd backend
start.bat
```

On macOS/Linux (Backend):
```bash
cd backend
./start.sh
```

Frontend (in new terminal):
```bash
npm run dev
```

### Step 5: Access the Application

1. Open your browser and go to: http://localhost:3000
2. Wait for the backend to finish indexing repositories (2-5 minutes on first run)
3. Start asking questions!

## Detailed Installation Steps

### Backend Setup (Manual)

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Create virtual environment:**
   ```bash
   # Windows
   python -m venv venv
   venv\Scripts\activate

   # macOS/Linux
   python3 -m venv venv
   source venv/bin/activate
   ```

3. **Install dependencies:**
   ```bash
   pip install --upgrade pip
   pip install -r requirements.txt
   ```

4. **Create environment file:**
   ```bash
   # Windows
   copy .env.example .env

   # macOS/Linux
   cp .env.example .env
   ```

5. **Start the server:**
   ```bash
   python app.py
   ```

   The backend will:
   - Start on http://localhost:8000
   - Clone the three GitHub repositories
   - Index all relevant files
   - Be ready to answer questions

### Frontend Setup (Manual)

1. **Navigate to root directory:**
   ```bash
   cd ..  # If you're in backend directory
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

3. **Create environment file:**
   ```bash
   # Windows
   copy .env.example .env

   # macOS/Linux
   cp .env.example .env
   ```

4. **Start development server:**
   ```bash
   npm run dev
   ```

   The frontend will start on http://localhost:3000

## First Run

### What Happens on First Run?

When you start the backend for the first time:

1. **Repository Cloning** (1-2 minutes)
   - Clones Cosmic Signature smart contracts repo
   - Clones Augur Explorer backend API repo
   - Clones Cosmic Front Alternate frontend repo

2. **Code Indexing** (1-3 minutes)
   - Scans all files in the repositories
   - Filters supported file types (.sol, .py, .js, .ts, etc.)
   - Creates searchable document index
   - Stores in memory for fast retrieval

3. **Ready to Use!**
   - Backend API available at http://localhost:8000
   - Frontend available at http://localhost:3000

### Monitoring Progress

Watch the backend terminal for status messages:

```
INFO:     Initializing RAG pipeline...
INFO:     No existing index found. Starting indexing process...
INFO:     Fetching smart_contracts from https://github.com/...
INFO:     Cloning repository smart_contracts...
INFO:     Created 245 documents from smart_contracts
INFO:     Fetching backend_api from https://github.com/...
...
INFO:     Indexing complete!
INFO:     RAG pipeline ready!
```

## Verifying Installation

### Check Backend Health

Visit http://localhost:8000/health or run:

```bash
curl http://localhost:8000/health
```

Expected response:
```json
{
  "status": "healthy",
  "pipeline_ready": true,
  "indexed": true
}
```

### Check Frontend

Visit http://localhost:3000 - you should see the chat interface.

## Troubleshooting

### Common Issues

#### 1. "Python not found"
**Solution:** Install Python 3.9+ and add to PATH

#### 2. "pip not found"
**Solution:** 
```bash
python -m ensurepip --upgrade
```

#### 3. "Virtual environment activation failed"
**Windows Solution:**
```bash
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

#### 4. "Cannot connect to backend"
**Solution:**
- Ensure backend is running on port 8000
- Check firewall settings
- Verify `.env` has `NEXT_PUBLIC_BACKEND_URL=http://localhost:8000`

#### 5. "Git clone failed"
**Solution:**
- Check internet connection
- Ensure git is installed
- Try manual clone to test access

#### 6. "Port already in use"
**Solution:**
```bash
# Change backend port in backend/app.py
uvicorn.run(app, host="0.0.0.0", port=8001)  # Use 8001 instead

# Update frontend .env
NEXT_PUBLIC_BACKEND_URL=http://localhost:8001
```

### Still Having Issues?

1. Check the logs in both terminals
2. Ensure all prerequisites are installed
3. Try deleting and recreating the virtual environment
4. Check GitHub repository access

## Environment Variables

### Backend (.env)
```bash
PORT=8000
HOST=0.0.0.0
DATA_DIR=./data
LOG_LEVEL=INFO
```

### Frontend (.env)
```bash
NEXT_PUBLIC_BACKEND_URL=http://localhost:8000
```

## Next Steps

After successful setup:

1. ✅ Try the example questions
2. ✅ Explore the code structure
3. ✅ Customize for your needs
4. ✅ Deploy to production (see README.md)

## Development Mode

### Hot Reload

Both frontend and backend support hot reload:

- **Frontend:** Next.js automatically reloads on file changes
- **Backend:** Use `uvicorn app:app --reload` for auto-reload

### Debug Mode

Enable debug logging:

```bash
# In backend/.env
LOG_LEVEL=DEBUG
```

## Production Setup

See README.md for production deployment instructions.

---

Need help? Open an issue on GitHub!
