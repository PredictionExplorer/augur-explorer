# Architecture - Cosmic Signature FAQ Bot

This document explains the technical architecture and design decisions of the FAQ bot.

## System Overview

```
┌─────────────────────────────────────────────────────────────┐
│                         User Browser                         │
│                    (http://localhost:3000)                   │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            │ HTTP/REST
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                      Next.js Frontend                        │
│  ┌────────────────────────────────────────────────────────┐ │
│  │  • React Components (TypeScript)                       │ │
│  │  • Chat Interface                                      │ │
│  │  • State Management                                    │ │
│  └────────────────────────────────────────────────────────┘ │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            │ POST /api/query
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                   FastAPI Backend (Python)                   │
│  ┌────────────────────────────────────────────────────────┐ │
│  │  • REST API Endpoints                                  │ │
│  │  • Request/Response Handling                           │ │
│  │  • CORS Configuration                                  │ │
│  └────────────────────────────────────────────────────────┘ │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Haystack RAG Pipeline                     │
│  ┌────────────────────────────────────────────────────────┐ │
│  │                                                         │ │
│  │  1. Query Input                                        │ │
│  │         │                                               │ │
│  │         ▼                                               │ │
│  │  2. Document Retriever (BM25)                          │ │
│  │         │                                               │ │
│  │         ▼                                               │ │
│  │  3. Answer Generator                                   │ │
│  │         │                                               │ │
│  │         ▼                                               │ │
│  │  4. Formatted Answer                                   │ │
│  │                                                         │ │
│  └────────────────────────────────────────────────────────┘ │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                  In-Memory Document Store                    │
│  ┌────────────────────────────────────────────────────────┐ │
│  │  • Indexed Code Documents                              │ │
│  │  • Metadata (file path, repo, type)                    │ │
│  │  • BM25 Index                                          │ │
│  └────────────────────────────────────────────────────────┘ │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    GitHub Repository Cache                   │
│  ┌────────────────────────────────────────────────────────┐ │
│  │  • Smart Contracts (Cosmic-Signature)                  │ │
│  │  • Backend API (augur-explorer)                        │ │
│  │  • Frontend (cosmic-front-alternate)                   │ │
│  └────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

## Component Details

### Frontend Layer

**Technology:** Next.js 15 with React 19, TypeScript

**Key Files:**
- `app/page.tsx` - Main page component
- `app/components/ChatInterface.tsx` - Chat UI component
- `app/layout.tsx` - Root layout
- `app/globals.css` - Global styles

**Responsibilities:**
1. Render chat interface
2. Handle user input
3. Make API calls to backend
4. Display responses with sources
5. Manage UI state (messages, loading)

**Design Patterns:**
- Client-side rendering for interactive components
- React hooks for state management
- CSS modules for scoped styling
- Responsive design (mobile-first)

### Backend API Layer

**Technology:** FastAPI (Python), Uvicorn server

**Key Files:**
- `backend/app.py` - Main FastAPI application
- `backend/.env` - Environment configuration

**Endpoints:**
- `GET /` - Root endpoint
- `GET /health` - Health check
- `POST /api/query` - Query the FAQ bot
- `POST /api/reindex` - Force re-indexing

**Responsibilities:**
1. Serve REST API
2. Handle CORS for frontend
3. Initialize RAG pipeline on startup
4. Route queries to pipeline
5. Format responses

**Features:**
- Automatic startup initialization
- Error handling and logging
- Async support for better performance

### RAG Pipeline Layer

**Technology:** Haystack AI Framework

**Key Files:**
- `backend/rag_pipeline.py` - Main pipeline logic
- `backend/answer_generator.py` - Answer generation
- `backend/github_fetcher.py` - Repository fetcher

**Components:**

1. **Document Store (InMemoryDocumentStore)**
   - Stores all indexed documents
   - Fast in-memory access
   - Supports metadata filtering

2. **Retriever (InMemoryBM25Retriever)**
   - BM25 algorithm for keyword matching
   - Returns top-k relevant documents
   - Considers term frequency and document length

3. **Answer Generator**
   - Custom logic for answer formatting
   - Question type detection
   - Smart contract analysis
   - Code snippet extraction

**Pipeline Flow:**
```
User Query → Retriever → Retrieved Documents → Answer Generator → Formatted Answer
```

### Data Processing Layer

**GitHub Fetcher:**

**Process:**
1. Clone or update repository (shallow clone, depth=1)
2. Scan all files recursively
3. Filter by file extension and directory
4. Read file contents
5. Create Haystack Document objects
6. Store in document store

**Supported File Types:**
- `.sol` - Solidity smart contracts
- `.py` - Python code
- `.js, .jsx` - JavaScript code
- `.ts, .tsx` - TypeScript code
- `.json` - Configuration files
- `.md` - Documentation
- `.yml, .yaml` - YAML configs

**Excluded:**
- `node_modules/`, `.git/`, `build/`, `dist/`
- Lock files (package-lock.json, yarn.lock)
- Binary files
- Files > 500KB

## Data Flow

### Initialization Flow

```
1. Backend starts
2. Check if documents are indexed
3. If not indexed:
   a. Clone 3 GitHub repositories
   b. Scan all files
   c. Filter and process files
   d. Create Document objects
   e. Write to document store
4. Initialize RAG pipeline
5. Ready to serve queries
```

### Query Flow

```
1. User types question in frontend
2. Frontend sends POST to /api/query
3. Backend receives request
4. RAG pipeline processes query:
   a. Retriever searches document store
   b. Returns top 5 relevant documents
   c. Answer generator analyzes documents
   d. Formats structured answer
5. Backend returns response with sources
6. Frontend displays answer
```

## Design Decisions

### Why BM25 Instead of Embeddings?

**Chosen:** BM25 (keyword-based retrieval)

**Rationale:**
- Faster initialization (no embedding generation needed)
- Lower resource requirements
- Good performance for code search
- Exact keyword matching works well for technical queries

**Future Enhancement:**
- Can add embedding-based retrieval later
- Hybrid approach (BM25 + embeddings) for best results

### Why In-Memory Document Store?

**Chosen:** InMemoryDocumentStore

**Rationale:**
- Fast query performance
- Simple setup
- Sufficient for small-medium codebases
- No external dependencies

**Production Alternative:**
- FAISS for vector similarity
- Elasticsearch for large-scale
- Weaviate for hybrid search

### Why Shallow Clone?

**Chosen:** Git shallow clone (depth=1)

**Rationale:**
- Faster cloning (only latest commit)
- Less disk space
- We only need current code, not history

### Why Custom Answer Generator?

**Chosen:** Custom logic instead of LLM

**Rationale:**
- No API keys required
- Lower latency
- Cost-free
- Transparent logic

**Production Enhancement:**
- Add OpenAI/Anthropic for better answers
- Use prompt engineering
- Implement streaming responses

## Performance Characteristics

### Initialization (First Run)

**Time:** 2-5 minutes
**Factors:**
- Network speed (cloning repos)
- File count
- CPU speed

**Optimization:**
- Parallel repository cloning
- Incremental updates
- Cache management

### Query Performance

**Typical:** 100-500ms
**Breakdown:**
- Retrieval: 50-200ms
- Answer generation: 50-300ms

**Optimization:**
- Caching common queries
- Pre-computing common patterns
- Using faster retrievers

### Resource Usage

**Memory:**
- Frontend: ~50MB
- Backend: ~200-500MB
- Document store: ~100-300MB

**Disk:**
- Repositories: ~100-500MB
- Node modules: ~200MB
- Python packages: ~500MB

## Scalability Considerations

### Current Limitations

1. **In-memory storage** - Limited by RAM
2. **Single process** - No horizontal scaling
3. **No caching** - Repeated queries re-computed

### Scaling Options

**Vertical Scaling:**
- More RAM for larger codebases
- Faster CPU for quicker processing

**Horizontal Scaling:**
- Load balancer + multiple backend instances
- Shared document store (Redis, Elasticsearch)
- Caching layer (Redis, Memcached)

**Optimization:**
- Query result caching
- Pre-computed embeddings
- CDN for frontend
- Database for document storage

## Security Considerations

### Current Implementation

- No authentication required
- Public GitHub repositories only
- Read-only operations
- CORS limited to localhost

### Production Recommendations

1. **Authentication:** Add user auth (JWT, OAuth)
2. **Rate Limiting:** Prevent abuse
3. **API Keys:** Secure access to API
4. **HTTPS:** Encrypt traffic
5. **Input Validation:** Prevent injection attacks
6. **Secrets Management:** Use environment variables

## Future Enhancements

### High Priority

1. **LLM Integration**
   - OpenAI GPT-4 or Anthropic Claude
   - Better answer quality
   - Natural language understanding

2. **Embedding-based Retrieval**
   - Semantic search
   - Better context understanding
   - Hybrid search (BM25 + embeddings)

3. **Persistent Storage**
   - Save indexed documents
   - Faster startup
   - Version control

### Medium Priority

1. **Caching Layer**
   - Cache common queries
   - Reduce latency
   - Lower resource usage

2. **Real-time Updates**
   - WebSocket support
   - Streaming responses
   - Live code updates

3. **Analytics**
   - Query tracking
   - Popular questions
   - Performance metrics

### Nice to Have

1. **Multi-repository Support**
   - User can add repos
   - Dynamic indexing
   - Repository management UI

2. **Code Execution**
   - Run code snippets
   - Validate suggestions
   - Interactive examples

3. **Visualization**
   - Code dependency graphs
   - Architecture diagrams
   - Flow charts

## Technology Stack Summary

| Layer | Technology | Purpose |
|-------|-----------|---------|
| Frontend | Next.js 15 + React 19 | UI framework |
| Styling | CSS Modules | Component styling |
| Language | TypeScript | Type safety |
| Backend | FastAPI | REST API server |
| RAG Framework | Haystack AI | Document retrieval |
| Retrieval | BM25 | Keyword-based search |
| Document Store | InMemoryDocumentStore | Document storage |
| Version Control | Git | Repository management |
| Server | Uvicorn | ASGI server |
| Language (Backend) | Python 3.9+ | Backend logic |

## Development vs Production

### Development

- In-memory storage
- No caching
- Single instance
- Local repositories
- Detailed logging

### Production

- Persistent storage (PostgreSQL + FAISS)
- Redis caching
- Load balanced instances
- CDN for repos
- Structured logging (ELK stack)
- Monitoring (Prometheus + Grafana)
- Error tracking (Sentry)

---

This architecture provides a solid foundation while allowing for future enhancements and scaling.
