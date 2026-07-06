# Cosmic Signature FAQ Bot

An AI-powered FAQ bot that helps users understand the Cosmic Signature project by analyzing source code from GitHub repositories using Haystack AI framework.

## Features

- 🤖 **AI-Powered Answers**: Uses Haystack's RAG (Retrieval Augmented Generation) pipeline
- 🧠 **LLM Integration**: Optional OpenAI/Anthropic support for intelligent, user-friendly answers
- 📚 **Code Analysis**: Analyzes smart contracts, backend API, and frontend code
- 🔍 **Intelligent Search**: Keyword-based retrieval for relevant code snippets
- 💬 **Interactive Chat**: Beautiful Next.js chat interface
- 🚀 **Fast & Efficient**: Caches repositories locally for quick responses

## Architecture

### Frontend (Next.js)
- Modern React-based chat interface
- TypeScript for type safety
- Responsive design with beautiful UI
- Real-time message streaming

### Backend (Python + Haystack)
- FastAPI for high-performance API
- Haystack AI for RAG pipeline
- In-memory document store with BM25 retriever
- Automatic GitHub repository fetching and indexing

## Analyzed Repositories

The bot analyzes code from three repositories:

1. **Smart Contracts**: https://github.com/PredictionExplorer/Cosmic-Signature
2. **Backend API**: https://github.com/PredictionExplorer/augur-explorer
3. **Frontend**: https://github.com/PredictionExplorer/cosmic-front-alternate

## Prerequisites

- Node.js 18+ and npm
- Python 3.9+
- Git

## Installation

### 1. Clone this repository

```bash
git clone <your-repo-url>
cd faq
```

### 2. Install Frontend Dependencies

```bash
npm install
```

### 3. Install Backend Dependencies

```bash
cd backend
python -m venv venv

# On Windows
venv\Scripts\activate

# On macOS/Linux
source venv/bin/activate

pip install -r requirements.txt
```

### 4. Configure Environment Variables

```bash
# Copy example env files
cp .env.example .env
cp backend/.env.example backend/.env

# Edit .env files if needed
```

## Running the Application

### Option 1: Run Both Services Manually

**Terminal 1 - Backend:**
```bash
cd backend
# Activate virtual environment first
python app.py
```

The backend will:
- Start on http://localhost:8000
- Automatically clone and index repositories on first run (takes 2-5 minutes)
- Create a cache in `backend/data/repos/`

**Terminal 2 - Frontend:**
```bash
npm run dev
```

The frontend will start on http://localhost:3000

### Option 2: Use the Start Script (Windows)

```bash
# Run both frontend and backend
npm run start:all
```

## Usage

1. Open http://localhost:3000 in your browser
2. Wait for the backend to finish indexing (first run only)
3. **(Optional but Recommended)** Set up OpenAI or Anthropic API key for better answers - see [LLM_SETUP.md](LLM_SETUP.md)
4. Ask questions about the Cosmic Signature project!

### Example Questions

- "How can I bid on the website?"
- "Where is information about the distribution of prizes?"
- "How many smart contracts are there?"
- "What is the Cosmic Signature project?"
- "How does the bidding system work?"
- "What tokens are used in the smart contracts?"

## API Endpoints

### `POST /api/query`
Query the FAQ bot

**Request:**
```json
{
  "question": "How can I bid on the website?"
}
```

**Response:**
```json
{
  "answer": "Based on the codebase...",
  "sources": [
    "smart_contracts: contracts/Bidding.sol",
    "frontend: components/BidForm.tsx"
  ]
}
```

### `GET /health`
Health check endpoint

### `POST /api/reindex`
Force re-indexing of all repositories

## Project Structure

```
faq/
├── app/                          # Next.js frontend
│   ├── components/              # React components
│   │   └── ChatInterface.tsx   # Main chat component
│   ├── layout.tsx              # Root layout
│   ├── page.tsx                # Home page
│   └── globals.css             # Global styles
├── backend/                     # Python backend
│   ├── app.py                  # FastAPI application
│   ├── rag_pipeline.py         # Haystack RAG pipeline
│   ├── github_fetcher.py       # GitHub repository fetcher
│   ├── requirements.txt        # Python dependencies
│   └── .env.example           # Environment variables template
├── package.json                # Node.js dependencies
├── tsconfig.json              # TypeScript configuration
├── next.config.js             # Next.js configuration
└── README.md                  # This file
```

## Customization

### Adding More Repositories

Edit `backend/rag_pipeline.py` and add to the `REPOSITORIES` dict:

```python
REPOSITORIES = {
    "smart_contracts": "https://github.com/...",
    "backend_api": "https://github.com/...",
    "frontend": "https://github.com/...",
    "your_new_repo": "https://github.com/your/repo",  # Add here
}
```

### Changing Indexed File Types

Edit `backend/github_fetcher.py` and modify `SUPPORTED_EXTENSIONS`:

```python
SUPPORTED_EXTENSIONS = {
    '.sol', '.py', '.js', '.ts',
    '.your_extension',  # Add here
}
```

### Enhancing Answer Quality

For production use, consider integrating a language model:

1. Add OpenAI/Anthropic API integration
2. Use embedding-based retrieval instead of BM25
3. Implement re-ranking for better results

Example using OpenAI:
```python
from haystack.components.generators import OpenAIGenerator

generator = OpenAIGenerator(
    api_key=os.getenv("OPENAI_API_KEY"),
    model="gpt-4"
)
```

## Performance Tips

1. **First Run**: Initial indexing takes 2-5 minutes. Subsequent runs use cached data.
2. **Update Cache**: Run `POST /api/reindex` to fetch latest code from GitHub
3. **Memory**: The in-memory document store is fast but uses RAM. For production, consider FAISS or Elasticsearch.

## Troubleshooting

### Backend won't start
- Ensure Python 3.9+ is installed
- Activate virtual environment
- Check all dependencies are installed: `pip install -r requirements.txt`

### Frontend can't connect to backend
- Verify backend is running on port 8000
- Check CORS settings in `backend/app.py`
- Ensure `.env` has correct `NEXT_PUBLIC_BACKEND_URL`

### Indexing fails
- Check internet connection (needs to clone from GitHub)
- Ensure git is installed and accessible
- Check disk space (repos can be 100MB+)

### Slow responses
- First query after startup may be slow
- Consider adding embedding-based retrieval
- Increase `top_k` parameter for more context

## Development

### Adding New Features

1. **Frontend**: Add components in `app/components/`
2. **Backend**: Add endpoints in `backend/app.py`
3. **RAG Logic**: Modify `backend/rag_pipeline.py`

### Testing

```bash
# Frontend
npm run lint

# Backend
cd backend
pytest
```

## Production Deployment

### Frontend (Vercel/Netlify)
```bash
npm run build
npm start
```

### Backend (Docker)
```dockerfile
FROM python:3.9-slim
WORKDIR /app
COPY backend/requirements.txt .
RUN pip install -r requirements.txt
COPY backend/ .
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

MIT License - feel free to use for your projects!

## Acknowledgments

- Built with [Haystack](https://haystack.deepset.ai/) by deepset
- Powered by [Next.js](https://nextjs.org/)
- UI inspired by modern chat interfaces

## Support

For issues or questions:
- Open an issue on GitHub
- Check the troubleshooting section
- Review Haystack documentation

---

Made with ❤️ for the Cosmic Signature community
