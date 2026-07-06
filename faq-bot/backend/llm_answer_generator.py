"""
Enhanced Answer Generator using LLMs
Generates user-friendly, actionable answers from code context
"""
import os
import logging
from typing import List
from haystack import Document

from code_understander import CodeUnderstander

logger = logging.getLogger(__name__)

OPENAI_MODELS = ["gpt-4o-mini", "gpt-3.5-turbo", "gpt-4o"]


class LLMAnswerGenerator:
    def __init__(self):
        self.openai_api_key = os.getenv("OPENAI_API_KEY", "")
        self.gemini_api_key = os.getenv("GEMINI_API_KEY", "")
        self.anthropic_api_key = os.getenv("ANTHROPIC_API_KEY", "")
        self.code_understander = CodeUnderstander()
        self.llm_type = None

        if self.openai_api_key and not self.openai_api_key.startswith("your_"):
            self.llm_type = "openai"
            logger.info("Using OpenAI for answer generation")
        elif self.gemini_api_key and not self.gemini_api_key.startswith("your_"):
            self.llm_type = "gemini"
            logger.info("Using Google Gemini for answer generation (free!)")
        elif self.anthropic_api_key and not self.anthropic_api_key.startswith("your_"):
            self.llm_type = "anthropic"
            logger.info("Using Anthropic Claude for answer generation")
        else:
            logger.info("No LLM API key configured - using Smart Code Understander")

    def generate_answer(self, question: str, documents: List[Document]) -> str:
        if not documents:
            return self.code_understander.generate_user_answer(question, [])
        if self.llm_type == "openai":
            return self._generate_with_openai(question, documents)
        elif self.llm_type == "gemini":
            return self._generate_with_gemini(question, documents)
        elif self.llm_type == "anthropic":
            return self._generate_with_anthropic(question, documents)
        else:
            return self.code_understander.generate_user_answer(question, documents)

    def _generate_with_openai(self, question: str, documents: List[Document]) -> str:
        try:
            from openai import OpenAI
            client = OpenAI(api_key=self.openai_api_key)
            context = self._build_context(documents)
            prompt = self._build_prompt(question, context)
            last_error = None
            for model in OPENAI_MODELS:
                try:
                    logger.info(f"Trying OpenAI model: {model}")
                    response = client.chat.completions.create(
                        model=model,
                        messages=[
                            {
                                "role": "system",
                                "content": (
                                    "You are a helpful assistant for the Cosmic Signature website - "
                                    "a blockchain-based bidding game. Explain things clearly for new users. "
                                    "Always give step-by-step instructions."
                                ),
                            },
                            {"role": "user", "content": prompt},
                        ],
                        temperature=0.7,
                        max_tokens=900,
                    )
                    logger.info(f"OpenAI answered using model: {model}")
                    return response.choices[0].message.content
                except Exception as model_err:
                    logger.warning(f"Model {model} failed: {model_err}")
                    last_error = model_err
                    continue
            logger.error(f"All OpenAI models failed: {last_error}")
        except Exception as e:
            logger.error(f"OpenAI client error: {e}")
        logger.info("Falling back to Smart Code Understander")
        return self.code_understander.generate_user_answer(question, documents)

    def _generate_with_gemini(self, question: str, documents: List[Document]) -> str:
        try:
            import google.generativeai as genai
            genai.configure(api_key=self.gemini_api_key)
            model = genai.GenerativeModel("gemini-2.0-flash")
            context = self._build_context(documents)
            prompt = self._build_prompt(question, context)
            response = model.generate_content(prompt)
            logger.info("Gemini answered successfully")
            return response.text
        except Exception as e:
            logger.error(f"Gemini error: {e}")
        logger.info("Falling back to Smart Code Understander")
        return self.code_understander.generate_user_answer(question, documents)

    def _generate_with_anthropic(self, question: str, documents: List[Document]) -> str:
        try:
            from anthropic import Anthropic
            client = Anthropic(api_key=self.anthropic_api_key)
            context = self._build_context(documents)
            prompt = self._build_prompt(question, context)
            response = client.messages.create(
                model="claude-3-5-haiku-20241022",
                max_tokens=900,
                messages=[{"role": "user", "content": prompt}],
            )
            logger.info("Anthropic answered successfully")
            return response.content[0].text
        except Exception as e:
            logger.error(f"Anthropic error: {e}")
        logger.info("Falling back to Smart Code Understander")
        return self.code_understander.generate_user_answer(question, documents)

    def _build_prompt(self, question: str, context: str) -> str:
        return f"""You are a helpful assistant for the Cosmic Signature website - a blockchain-based bidding game.

The user asked: "{question}"

Here is relevant code from the project (smart contracts, backend API, frontend):

{context}

Write a clear, friendly step-by-step answer for a new user who has never used the site before.
Include:
1. Which page or URL to visit (e.g. /game/play)
2. What buttons, toggles, or forms to interact with
3. What actions to take in order
4. Requirements (wallet connected, token balance, etc.)
5. What to expect after completing the action

Be specific and friendly. Do NOT show raw code or imports."""

    def _build_context(self, documents: List[Document], max_chars: int = 6000) -> str:
        def priority(doc: Document) -> int:
            fp = doc.meta.get("file_path", "")
            if "test" in fp.lower():
                return 2
            if fp.endswith((".tsx", ".jsx", ".ts", ".js")) and "page" in fp:
                return 0
            if fp.endswith(".sol"):
                return 1
            return 2

        context_parts = []
        total_chars = 0
        for i, doc in enumerate(sorted(documents, key=priority), 1):
            repo = doc.meta.get("repository", "unknown")
            file_path = doc.meta.get("file_path", "unknown")
            content = doc.content[:1500] + "\n...(truncated)" if len(doc.content) > 1500 else doc.content
            part = f"--- [{i}] {repo}/{file_path} ---\n{content}\n\n"
            if total_chars + len(part) > max_chars:
                break
            context_parts.append(part)
            total_chars += len(part)
        return "".join(context_parts)
