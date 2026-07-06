# Quick Start Guide

Get the Cosmic Signature FAQ Bot running in 5 minutes!

## Prerequisites

- Node.js 18+ ([Download](https://nodejs.org/))
- Python 3.9+ ([Download](https://www.python.org/downloads/))
- Git ([Download](https://git-scm.com/))

## Installation

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd faq
```

### 2. Backend Setup

**Windows:**
```bash
cd backend
setup.bat
```

**macOS/Linux:**
```bash
cd backend
chmod +x setup.sh
./setup.sh
```

### 3. Frontend Setup

```bash
# Go back to root directory
cd ..

# Install dependencies
npm install

# Create environment file
copy .env.example .env  # Windows
cp .env.example .env    # macOS/Linux
```

## Running the Application

### Option 1: Using Scripts (Recommended)

**Terminal 1 - Backend:**
```bash
cd backend
start.bat     # Windows
./start.sh    # macOS/Linux
```

**Terminal 2 - Frontend:**
```bash
npm run dev
```

### Option 2: Docker (Easiest)

```bash
docker-compose up
```

That's it! Docker will build and start both services.

### Option 3: Manual

**Terminal 1 - Backend:**
```bash
cd backend
venv\Scripts\activate     # Windows
source venv/bin/activate  # macOS/Linux
python app.py
```

**Terminal 2 - Frontend:**
```bash
npm run dev
```

## Access the Application

1. **Open your browser:** http://localhost:3000
2. **Wait for indexing:** First run takes 2-5 minutes
3. **Start chatting!** Ask questions about Cosmic Signature

## Example Questions

Try asking:
- "How can I bid on the website?"
- "How many smart contracts are there?"
- "Where is information about prize distribution?"
- "What is the Cosmic Signature project?"

## Verification

### Check Backend Status

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

Open http://localhost:3000 - you should see the chat interface.

## Troubleshooting

### Backend won't start
- Make sure Python 3.9+ is installed: `python --version`
- Check if virtual environment is activated
- Try: `pip install -r requirements.txt`

### Frontend won't start
- Make sure Node.js 18+ is installed: `node --version`
- Try: `npm install` again
- Delete `node_modules` and `.next`, then reinstall

### Can't connect to backend
- Verify backend is running on port 8000
- Check `.env` file: `NEXT_PUBLIC_BACKEND_URL=http://localhost:8000`
- Check firewall settings

### Port already in use
**Backend (port 8000):**
- Kill the process or change port in `backend/app.py`

**Frontend (port 3000):**
- Use a different port: `npm run dev -- -p 3001`

## What's Next?

- **Learn More:** Read [README.md](README.md) for detailed documentation
- **Architecture:** See [ARCHITECTURE.md](ARCHITECTURE.md) for technical details
- **Setup Details:** Check [SETUP.md](SETUP.md) for comprehensive setup guide
- **Contributing:** Read [CONTRIBUTING.md](CONTRIBUTING.md) to contribute

## Project Structure

```
faq/
├── app/                  # Frontend (Next.js)
├── backend/             # Backend (Python + Haystack)
│   ├── app.py          # Main API server
│   ├── rag_pipeline.py # RAG logic
│   └── ...
├── package.json        # Node dependencies
└── README.md          # Full documentation
```

## Common Commands

```bash
# Start backend
cd backend && python app.py

# Start frontend
npm run dev

# Build frontend for production
npm run build

# Re-index repositories
curl -X POST http://localhost:8000/api/reindex

# Check backend logs
# Watch the terminal where backend is running
```

## Getting Help

- **Issues:** Check existing issues or create a new one
- **Documentation:** Read README.md and SETUP.md
- **Questions:** Open a discussion on GitHub

---

Happy chatting! 🚀
