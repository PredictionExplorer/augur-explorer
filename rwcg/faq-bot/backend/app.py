"""
FastAPI backend for Cosmic Signature FAQ Bot.

Architecture: generated knowledge base + Haystack retrieval + Codex MCP synthesis.
Both Haystack and Codex MCP must be healthy for /api/query to succeed.
"""
from __future__ import annotations

import logging
import os
from contextlib import asynccontextmanager
from typing import Optional

from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import FileResponse
from pydantic import BaseModel, Field

from llm.codex_mcp_client import CodexMCPClient, CodexMCPError
from orchestrator import HaystackUnavailableError, Orchestrator
from retrieval.pipeline import KnowledgeRetriever
from sessions.store import SessionStore

load_dotenv()

logging.basicConfig(level=os.getenv("LOG_LEVEL", "INFO"))
logger = logging.getLogger(__name__)

_raw_origins = os.getenv("CORS_ORIGINS", "http://localhost:3000,http://localhost:3001")
_allowed_origins = [o.strip() for o in _raw_origins.split(",") if o.strip()]

retriever: Optional[KnowledgeRetriever] = None
codex_client: Optional[CodexMCPClient] = None
orchestrator: Optional[Orchestrator] = None
session_store = SessionStore(ttl_seconds=int(os.getenv("SESSION_TTL_SECONDS", "86400")))


class QueryRequest(BaseModel):
    question: str
    session_id: str | None = None


class QueryResponse(BaseModel):
    answer: str
    sources: list[str]
    session_id: str


class ErrorResponse(BaseModel):
    error: str
    component: str = Field(description="haystack | codex_mcp | validation")


@asynccontextmanager
async def lifespan(app: FastAPI):
    global retriever, codex_client, orchestrator
    logger.info("Starting FAQ bot services...")

    retriever = KnowledgeRetriever(
        include_code_fallback=os.getenv("INCLUDE_CODE_FALLBACK", "true").lower() == "true"
    )
    try:
        count = retriever.index_knowledge_base()
        logger.info("Haystack indexed %d documents", count)
    except FileNotFoundError as exc:
        logger.error("Haystack indexing failed: %s", exc)
        raise

    codex_client = CodexMCPClient()
    await codex_client.start()

    orchestrator = Orchestrator(retriever, codex_client, session_store)
    health = await orchestrator.health()
    logger.info("Startup health: %s", health)
    if health["status"] != "healthy":
        logger.warning("System started in degraded mode: %s", health)

    yield

    if codex_client:
        await codex_client.stop()


app = FastAPI(
    title="Cosmic Signature FAQ Bot API",
    description="Generated knowledge base + Haystack + Codex MCP",
    version="2.0.0",
    lifespan=lifespan,
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=_allowed_origins,
    allow_origin_regex=r"http://.*\.compute\.amazonaws\.com(:\d+)?|http://\d+\.\d+\.\d+\.\d+(:\d+)?|https?://.*",
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

_STATIC_DIR = os.path.join(os.path.dirname(os.path.abspath(__file__)), "static")


@app.get("/test-ui")
async def test_ui():
    """Minimal browser UI for manual FAQ testing."""
    return FileResponse(os.path.join(_STATIC_DIR, "test-ui.html"))


@app.get("/")
async def root():
    indexed = retriever.is_ready if retriever else False
    codex_ready = codex_client.is_ready if codex_client else False
    return {
        "message": "Cosmic Signature FAQ Bot API",
        "status": "running",
        "haystack_indexed": indexed,
        "codex_mcp_ready": codex_ready,
    }


@app.get("/health")
async def health_check():
    if not orchestrator:
        raise HTTPException(status_code=503, detail="Orchestrator not initialized")
    status = await orchestrator.health()
    code = 200 if status["status"] == "healthy" else 503
    if code != 200:
        raise HTTPException(status_code=code, detail=status)
    return status


@app.get("/debug/stats")
async def debug_stats():
    if not retriever or not retriever.is_ready:
        raise HTTPException(status_code=503, detail="Haystack not ready")
    docs = retriever.document_store.filter_documents()
    sample = [
        {
            "source": d.meta.get("source"),
            "file_path": d.meta.get("file_path"),
            "tier": d.meta.get("tier"),
            "content_type": d.meta.get("content_type"),
        }
        for d in docs[:8]
    ]
    return {
        "total_documents": retriever.document_store.count_documents(),
        "knowledge_base": str(retriever.knowledge_dir),
        "facts_keys": list((retriever.load_facts() or {}).keys()),
        "sample_documents": sample,
        "sessions": session_store.stats(),
    }


@app.post("/api/query", response_model=QueryResponse, responses={503: {"model": ErrorResponse}})
async def query(request: QueryRequest):
    if not orchestrator:
        raise HTTPException(
            status_code=503,
            detail={"error": "Orchestrator not initialized", "component": "haystack"},
        )
    if not request.question or not request.question.strip():
        raise HTTPException(
            status_code=400,
            detail={"error": "Question cannot be empty", "component": "validation"},
        )

    try:
        result = await orchestrator.query(request.question, request.session_id)
        return QueryResponse(**result)
    except HaystackUnavailableError as exc:
        logger.error("Haystack error: %s", exc)
        raise HTTPException(
            status_code=503,
            detail={"error": str(exc), "component": "haystack"},
        )
    except CodexMCPError as exc:
        logger.error("Codex MCP error: %s", exc)
        raise HTTPException(
            status_code=503,
            detail={"error": str(exc), "component": "codex_mcp"},
        )


@app.post("/api/reindex")
async def reindex():
    if not retriever:
        raise HTTPException(status_code=503, detail="Retriever not initialized")
    try:
        count = retriever.index_knowledge_base(force=True)
        return {"status": "success", "documents_indexed": count}
    except Exception as exc:
        raise HTTPException(status_code=500, detail=str(exc))


@app.post("/api/generate-knowledge")
async def generate_knowledge():
    """Run deterministic knowledge generation scripts."""
    import subprocess
    import sys

    backend_root = os.path.dirname(os.path.abspath(__file__))
    proc = subprocess.run(
        [sys.executable, "-m", "knowledge.generate.run_all"],
        cwd=backend_root,
        capture_output=True,
        text=True,
    )
    if proc.returncode != 0:
        raise HTTPException(status_code=500, detail=proc.stderr or proc.stdout)
    if retriever:
        count = retriever.index_knowledge_base(force=True)
        return {"status": "success", "stdout": proc.stdout, "documents_indexed": count}
    return {"status": "success", "stdout": proc.stdout}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host=os.getenv("HOST", "0.0.0.0"), port=int(os.getenv("PORT", "8000")))
