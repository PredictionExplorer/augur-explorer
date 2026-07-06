"""
Enhanced Answer Generator for Cosmic Signature FAQ Bot
Provides intelligent answer generation based on retrieved code context
"""
import re
from typing import List, Dict, Any
from haystack import Document


class AnswerGenerator:
    """
    Generates intelligent answers from retrieved code documents
    """
    
    def __init__(self):
        """Initialize the answer generator"""
        self.question_patterns = self._init_question_patterns()
    
    def _init_question_patterns(self) -> Dict[str, List[str]]:
        """
        Initialize patterns for different types of questions
        Returns a dictionary mapping question types to relevant keywords
        """
        return {
            "bidding": ["bid", "bidding", "auction", "offer", "price"],
            "prize": ["prize", "reward", "distribution", "payout", "winning"],
            "contract": ["contract", "smart contract", "solidity", "blockchain"],
            "token": ["token", "erc", "cryptocurrency", "coin"],
            "staking": ["stake", "staking", "lock", "deposit"],
            "user": ["user", "account", "profile", "wallet"],
            "api": ["api", "endpoint", "route", "request"],
            "frontend": ["ui", "interface", "component", "page", "frontend"],
        }
    
    def generate_answer(self, question: str, documents: List[Document]) -> str:
        """
        Generate a comprehensive answer based on the question and retrieved documents
        
        Args:
            question: The user's question
            documents: List of relevant documents retrieved from the codebase
            
        Returns:
            A formatted answer string
        """
        if not documents:
            return self._no_documents_answer(question)
        
        # Identify question type
        question_type = self._identify_question_type(question)
        
        # Extract relevant information from documents
        relevant_info = self._extract_relevant_info(documents, question_type)
        
        # Build structured answer
        answer = self._build_answer(question, question_type, relevant_info)
        
        return answer
    
    def _identify_question_type(self, question: str) -> str:
        """Identify the type of question being asked"""
        question_lower = question.lower()
        
        # Count matches for each question type
        type_scores = {}
        for q_type, keywords in self.question_patterns.items():
            score = sum(1 for keyword in keywords if keyword in question_lower)
            type_scores[q_type] = score
        
        # Return the type with highest score, or "general" if no matches
        if max(type_scores.values()) > 0:
            return max(type_scores, key=type_scores.get)
        return "general"
    
    def _extract_relevant_info(self, documents: List[Document], question_type: str) -> Dict[str, Any]:
        """Extract and organize relevant information from documents"""
        info = {
            "smart_contracts": [],
            "backend_code": [],
            "frontend_code": [],
            "documentation": [],
            "count": {"contracts": 0, "files": 0, "repos": set()}
        }
        
        for doc in documents:
            repo = doc.meta.get("repository", "unknown")
            file_path = doc.meta.get("file_path", "unknown")
            file_type = doc.meta.get("file_type", "")
            content = doc.content[:1000]  # First 1000 chars
            
            # Track counts
            info["count"]["files"] += 1
            info["count"]["repos"].add(repo)
            
            # Categorize by file type and repository
            if file_type == ".sol":
                info["count"]["contracts"] += 1
                info["smart_contracts"].append({
                    "file": file_path,
                    "repo": repo,
                    "content": content,
                    "summary": self._extract_contract_info(content)
                })
            elif file_type == ".md":
                info["documentation"].append({
                    "file": file_path,
                    "repo": repo,
                    "content": content
                })
            elif "backend" in repo.lower() or "api" in repo.lower():
                info["backend_code"].append({
                    "file": file_path,
                    "repo": repo,
                    "content": content
                })
            else:
                info["frontend_code"].append({
                    "file": file_path,
                    "repo": repo,
                    "content": content
                })
        
        return info
    
    def _extract_contract_info(self, content: str) -> Dict[str, Any]:
        """Extract key information from smart contract code"""
        info = {
            "contract_names": [],
            "functions": [],
            "events": []
        }
        
        # Extract contract names
        contract_matches = re.findall(r'contract\s+(\w+)', content)
        info["contract_names"] = contract_matches
        
        # Extract function names
        function_matches = re.findall(r'function\s+(\w+)\s*\(', content)
        info["functions"] = function_matches[:5]  # First 5 functions
        
        # Extract events
        event_matches = re.findall(r'event\s+(\w+)\s*\(', content)
        info["events"] = event_matches[:5]  # First 5 events
        
        return info
    
    def _build_answer(self, question: str, question_type: str, info: Dict[str, Any]) -> str:
        """Build a comprehensive answer based on extracted information"""
        answer_parts = []
        
        # Start with a context-aware introduction
        intro = self._generate_intro(question, question_type, info)
        answer_parts.append(intro)
        
        # Add smart contract information if relevant
        if info["smart_contracts"]:
            contracts_info = self._format_contracts_info(info["smart_contracts"])
            answer_parts.append(contracts_info)
        
        # Add backend/API information if relevant
        if info["backend_code"]:
            backend_info = self._format_backend_info(info["backend_code"])
            answer_parts.append(backend_info)
        
        # Add frontend information if relevant
        if info["frontend_code"]:
            frontend_info = self._format_frontend_info(info["frontend_code"])
            answer_parts.append(frontend_info)
        
        # Add summary
        summary = self._generate_summary(question_type, info)
        answer_parts.append(summary)
        
        return "\n\n".join(answer_parts)
    
    def _generate_intro(self, question: str, question_type: str, info: Dict[str, Any]) -> str:
        """Generate an appropriate introduction based on the question"""
        count = info["count"]
        repos_str = ", ".join(count["repos"])
        
        if question_type == "bidding":
            return f"Based on the codebase analysis, I found information about bidding in {count['files']} files across the {repos_str} repositories."
        elif question_type == "prize":
            return f"Here's what I found about prize distribution in the Cosmic Signature project (analyzed {count['files']} files):"
        elif question_type == "contract":
            return f"I found {count['contracts']} smart contract(s) in the codebase. Here's the breakdown:"
        else:
            return f"Based on analyzing {count['files']} files from {len(count['repos'])} repositories, here's what I found:"
    
    def _format_contracts_info(self, contracts: List[Dict]) -> str:
        """Format smart contract information"""
        parts = ["**Smart Contracts:**"]
        
        for i, contract in enumerate(contracts[:3], 1):  # Show top 3
            summary = contract["summary"]
            file = contract["file"]
            
            if summary["contract_names"]:
                contracts_list = ", ".join(summary["contract_names"])
                parts.append(f"{i}. `{file}` - Contracts: {contracts_list}")
                
                if summary["functions"]:
                    functions_list = ", ".join(summary["functions"][:3])
                    parts.append(f"   - Key functions: {functions_list}")
        
        return "\n".join(parts)
    
    def _format_backend_info(self, backend_files: List[Dict]) -> str:
        """Format backend/API information"""
        parts = ["**Backend/API:**"]
        
        for i, file_info in enumerate(backend_files[:2], 1):  # Show top 2
            file = file_info["file"]
            content = file_info["content"]
            
            # Extract API routes or key functions
            routes = re.findall(r'@app\.(get|post|put|delete)\(["\']([^"\']+)', content)
            if routes:
                routes_str = ", ".join([route[1] for route in routes[:3]])
                parts.append(f"{i}. `{file}` - Routes: {routes_str}")
            else:
                parts.append(f"{i}. `{file}`")
        
        return "\n".join(parts)
    
    def _format_frontend_info(self, frontend_files: List[Dict]) -> str:
        """Format frontend information"""
        parts = ["**Frontend:**"]
        
        for i, file_info in enumerate(frontend_files[:2], 1):  # Show top 2
            file = file_info["file"]
            parts.append(f"{i}. `{file}`")
        
        return "\n".join(parts)
    
    def _generate_summary(self, question_type: str, info: Dict[str, Any]) -> str:
        """Generate a summary based on the question type"""
        if question_type == "bidding":
            return ("To place a bid, you'll need to interact with the smart contract functions. "
                   "Check the contract code for specific parameters and requirements.")
        elif question_type == "prize":
            return ("Prize distribution is typically handled by smart contract logic. "
                   "Review the contract code for the exact distribution mechanism and calculations.")
        elif question_type == "contract":
            return f"The project uses {info['count']['contracts']} smart contract(s) to manage blockchain operations."
        else:
            return "For more specific details, please refer to the source files mentioned above."
    
    def _no_documents_answer(self, question: str) -> str:
        """Generate an answer when no relevant documents are found"""
        return (
            "I couldn't find specific information in the codebase to answer your question. "
            "This could mean:\n\n"
            "1. The information might be in a file type that isn't indexed\n"
            "2. The question might need to be rephrased\n"
            "3. The feature might not be implemented yet\n\n"
            "Try rephrasing your question or asking about specific aspects like:\n"
            "- Smart contract functionality\n"
            "- API endpoints\n"
            "- Frontend components\n"
            "- Specific file or function names"
        )
