"""
RAG Pipeline for Cosmic Signature FAQ Bot using Haystack
"""
import os
import logging
from typing import List, Dict, Any, Optional
from pathlib import Path

from haystack import Pipeline, Document
from haystack.components.retrievers.in_memory import InMemoryBM25Retriever
from haystack.components.builders import PromptBuilder
from haystack.components.generators import HuggingFaceLocalGenerator
from haystack.document_stores.in_memory import InMemoryDocumentStore
from haystack.components.embedders import SentenceTransformersDocumentEmbedder, SentenceTransformersTextEmbedder
from haystack.components.retrievers.in_memory import InMemoryEmbeddingRetriever

from github_fetcher import GitHubFetcher
from llm_answer_generator import LLMAnswerGenerator

logger = logging.getLogger(__name__)


class CosmicRAGPipeline:
    """
    RAG Pipeline for answering questions about Cosmic Signature project
    """
    
    # Repository URLs
    REPOSITORIES = {
        "smart_contracts": "https://github.com/PredictionExplorer/Cosmic-Signature",
        "backend_api": "https://github.com/PredictionExplorer/augur-explorer",
        "frontend": "https://github.com/PredictionExplorer/cosmic-front-alternate"
    }
    
    def __init__(self, data_dir: str = "./data"):
        """Initialize the RAG pipeline"""
        self.data_dir = Path(data_dir)
        self.data_dir.mkdir(exist_ok=True)
        
        # Initialize document store
        self.document_store = InMemoryDocumentStore()
        
        # Initialize GitHub fetcher
        self.github_fetcher = GitHubFetcher(cache_dir=str(self.data_dir / "repos"))
        
        # Initialize answer generator (with LLM support)
        self.answer_generator = LLMAnswerGenerator()
        
        # Initialize pipeline components
        self._init_pipeline()
        
    def _init_pipeline(self):
        """Initialize Haystack pipeline components"""
        logger.info("Initializing pipeline components...")
        
        # Note: Retriever will be created after documents are indexed
        # to ensure the BM25 index is properly built
        self.retriever = None
        self.query_pipeline = None
        
        logger.info("Pipeline components initialized (retriever will be built after indexing)")
    
    def _build_retriever(self):
        """Build the retriever and query pipeline after documents are indexed"""
        logger.info("Building retriever and query pipeline...")
        
        # Create BM25 retriever with the populated document store
        self.retriever = InMemoryBM25Retriever(document_store=self.document_store)
        
        # Prompt template for question answering
        prompt_template = """
You are an AI assistant that helps users understand the Cosmic Signature project.
Answer the question based on the provided context from the codebase.

Context:
{% for document in documents %}
    File: {{ document.meta.file_path }}
    Repository: {{ document.meta.repository }}
    Content: {{ document.content }}
    ---
{% endfor %}

Question: {{ question }}

Please provide a clear, concise answer based on the code context above.
If you can't find the answer in the context, say so honestly.
Include references to specific files or functions when relevant.

Answer:
"""
        
        self.prompt_builder = PromptBuilder(template=prompt_template)
        
        # Build the query pipeline
        self.query_pipeline = Pipeline()
        self.query_pipeline.add_component("retriever", self.retriever)
        self.query_pipeline.add_component("prompt_builder", self.prompt_builder)
        
        # Connect components
        self.query_pipeline.connect("retriever.documents", "prompt_builder.documents")
        
        logger.info("Retriever and query pipeline built successfully")
    
    def is_indexed(self) -> bool:
        """Check if documents are already indexed"""
        return self.document_store.count_documents() > 0
    
    async def index_repositories(self, force: bool = False):
        """
        Fetch and index all repositories
        
        Args:
            force: If True, re-index even if documents exist
        """
        if not force and self.is_indexed():
            logger.info("Documents already indexed, skipping")
            return
        
        logger.info("Starting repository indexing...")
        all_documents = []
        
        for repo_name, repo_url in self.REPOSITORIES.items():
            logger.info(f"Fetching {repo_name} from {repo_url}")
            
            try:
                documents = self.github_fetcher.fetch_repository(repo_url, repo_name)
                logger.info(f"Fetched {len(documents)} documents from {repo_name}")
                all_documents.extend(documents)
            except Exception as e:
                logger.error(f"Error fetching {repo_name}: {e}")
                continue
        
        if all_documents:
            logger.info(f"Writing {len(all_documents)} documents to store...")
            self.document_store.write_documents(all_documents)
            logger.info("Indexing complete!")
            
            # Build the retriever AFTER documents are indexed
            self._build_retriever()
        else:
            logger.warning("No documents were indexed!")
    
    async def query(self, question: str, top_k: int = 5) -> Dict[str, Any]:
        """
        Query the RAG pipeline with a question
        
        Args:
            question: User's question
            top_k: Number of documents to retrieve
            
        Returns:
            Dictionary with answer and sources
        """
        if not self.is_indexed():
            return {
                "answer": "The knowledge base is not yet indexed. Please wait for the indexing to complete.",
                "sources": []
            }
        
        try:
            logger.info(f"Document store has {self.document_store.count_documents()} documents")
            
            # Use simple keyword search as fallback
            documents = self._simple_keyword_search(question, top_k)
            logger.info(f"Retrieved {len(documents)} documents for query: {question}")
            
            # Build a simple answer from retrieved documents
            answer = self._build_answer(question, documents)
            
            # Extract sources
            sources = [
                f"{doc.meta.get('repository', 'unknown')}: {doc.meta.get('file_path', 'unknown')}"
                for doc in documents[:3]  # Top 3 sources
            ]
            
            return {
                "answer": answer,
                "sources": sources
            }
            
        except Exception as e:
            logger.error(f"Error during query: {e}", exc_info=True)
            return {
                "answer": f"I encountered an error while processing your question: {str(e)}",
                "sources": []
            }
    
    def _simple_keyword_search(self, query: str, top_k: int = 5) -> List[Document]:
        """
        Simple keyword-based search as a reliable fallback
        """
        # Get all documents
        all_docs = self.document_store.filter_documents()
        
        # Extract keywords from query (lowercase, remove common words)
        stop_words = {'the', 'a', 'an', 'is', 'are', 'was', 'were', 'how', 'what', 'where', 'when', 'why', 'can', 'do', 'does', 'did'}
        query_words = [w.lower() for w in query.split() if w.lower() not in stop_words and len(w) > 2]
        
        logger.info(f"Search keywords: {query_words}")
        
        if not query_words:
            query_words = query.lower().split()
        
        # Score each document
        scored_docs = []
        for doc in all_docs:
            content_lower = doc.content.lower()
            file_path_lower = doc.meta.get('file_path', '').lower()
            
            # Calculate score based on keyword matches
            score = 0
            for word in query_words:
                # Count occurrences in content
                score += content_lower.count(word) * 2
                # Boost if in file path
                score += file_path_lower.count(word) * 5
            
            if score > 0:
                scored_docs.append((score, doc))
        
        # Sort by score and return top_k
        scored_docs.sort(key=lambda x: x[0], reverse=True)
        top_docs = [doc for score, doc in scored_docs[:top_k]]
        
        logger.info(f"Top scores: {[score for score, _ in scored_docs[:5]]}")
        
        return top_docs
    
    def _build_answer(self, question: str, documents: List[Document]) -> str:
        """
        Build an answer from retrieved documents using the enhanced answer generator
        """
        return self.answer_generator.generate_answer(question, documents)
