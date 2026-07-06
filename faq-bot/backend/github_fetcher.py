"""
GitHub Repository Fetcher for Cosmic Signature FAQ Bot
Fetches and processes code files from GitHub repositories
"""
import os
import logging
from typing import List, Dict, Optional
from pathlib import Path
import git
from haystack import Document

logger = logging.getLogger(__name__)


class GitHubFetcher:
    """
    Fetches code from GitHub repositories and converts to Haystack Documents
    """
    
    # File extensions to index
    SUPPORTED_EXTENSIONS = {
        '.sol',      # Solidity smart contracts
        '.py',       # Python
        '.js',       # JavaScript
        '.ts',       # TypeScript
        '.jsx',      # React JSX
        '.tsx',      # React TSX
        '.json',     # JSON config files
        '.md',       # Markdown documentation
        '.txt',      # Text files
        '.yml',      # YAML config
        '.yaml',     # YAML config
    }
    
    # Directories to skip
    SKIP_DIRECTORIES = {
        'node_modules',
        '.git',
        'dist',
        'build',
        '.next',
        '__pycache__',
        'venv',
        'env',
        '.venv',
        'coverage',
        '.pytest_cache',
    }
    
    # Files to skip
    SKIP_FILES = {
        'package-lock.json',
        'yarn.lock',
        'pnpm-lock.yaml',
        '.DS_Store',
    }
    
    def __init__(self, cache_dir: str = "./repos_cache"):
        """
        Initialize GitHub fetcher
        
        Args:
            cache_dir: Directory to cache cloned repositories
        """
        self.cache_dir = Path(cache_dir)
        self.cache_dir.mkdir(exist_ok=True, parents=True)
    
    def fetch_repository(self, repo_url: str, repo_name: str) -> List[Document]:
        """
        Fetch a GitHub repository and convert files to Haystack Documents
        
        Args:
            repo_url: GitHub repository URL
            repo_name: Name identifier for the repository
            
        Returns:
            List of Haystack Document objects
        """
        logger.info(f"Fetching repository: {repo_url}")
        
        # Clone or update repository
        repo_path = self._clone_or_update_repo(repo_url, repo_name)
        
        # Process files and create documents
        documents = self._process_repository(repo_path, repo_name)
        
        logger.info(f"Created {len(documents)} documents from {repo_name}")
        return documents
    
    def _clone_or_update_repo(self, repo_url: str, repo_name: str) -> Path:
        """
        Clone repository or update if it already exists
        
        Args:
            repo_url: GitHub repository URL
            repo_name: Name for the local directory
            
        Returns:
            Path to the local repository
        """
        repo_path = self.cache_dir / repo_name
        
        try:
            if repo_path.exists():
                logger.info(f"Repository {repo_name} already exists, pulling latest changes...")
                repo = git.Repo(repo_path)
                origin = repo.remotes.origin
                origin.pull()
            else:
                logger.info(f"Cloning repository {repo_name}...")
                git.Repo.clone_from(repo_url, repo_path, depth=1)  # Shallow clone
            
            return repo_path
            
        except Exception as e:
            logger.error(f"Error cloning/updating repository: {e}")
            raise
    
    def _process_repository(self, repo_path: Path, repo_name: str) -> List[Document]:
        """
        Process all supported files in the repository
        
        Args:
            repo_path: Path to the local repository
            repo_name: Name of the repository
            
        Returns:
            List of Document objects
        """
        documents = []
        
        for file_path in repo_path.rglob('*'):
            # Skip if it's a directory
            if file_path.is_dir():
                continue
            
            # Skip if in a blacklisted directory
            if any(skip_dir in file_path.parts for skip_dir in self.SKIP_DIRECTORIES):
                continue
            
            # Skip if filename is blacklisted
            if file_path.name in self.SKIP_FILES:
                continue
            
            # Skip if extension not supported
            if file_path.suffix not in self.SUPPORTED_EXTENSIONS:
                continue
            
            # Read and process file
            try:
                content = self._read_file(file_path)
                if content:
                    # Create relative path from repo root
                    rel_path = file_path.relative_to(repo_path)
                    
                    # Create Haystack Document
                    doc = Document(
                        content=content,
                        meta={
                            "repository": repo_name,
                            "file_path": str(rel_path),
                            "file_type": file_path.suffix,
                            "file_name": file_path.name,
                        }
                    )
                    documents.append(doc)
                    
            except Exception as e:
                logger.warning(f"Error processing file {file_path}: {e}")
                continue
        
        return documents
    
    def _read_file(self, file_path: Path, max_size_kb: int = 500) -> Optional[str]:
        """
        Read file content with size limit
        
        Args:
            file_path: Path to file
            max_size_kb: Maximum file size in KB
            
        Returns:
            File content as string or None
        """
        try:
            # Check file size
            file_size_kb = file_path.stat().st_size / 1024
            if file_size_kb > max_size_kb:
                logger.debug(f"Skipping large file {file_path} ({file_size_kb:.1f} KB)")
                return None
            
            # Read file with different encodings
            encodings = ['utf-8', 'latin-1', 'cp1252']
            
            for encoding in encodings:
                try:
                    with open(file_path, 'r', encoding=encoding) as f:
                        content = f.read()
                    
                    # Skip empty files
                    if not content.strip():
                        return None
                    
                    return content
                    
                except UnicodeDecodeError:
                    continue
            
            logger.debug(f"Could not decode file {file_path}")
            return None
            
        except Exception as e:
            logger.debug(f"Error reading file {file_path}: {e}")
            return None
