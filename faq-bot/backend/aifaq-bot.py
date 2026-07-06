"""Cosmic Signature AI FAQ bot — HTTP server entrypoint.

Use this module to start the backend so the process is visible in ps(1) as
``python aifaq-bot.py`` rather than ``python app.py``.
"""
from __future__ import annotations

if __name__ == "__main__":
    import os

    import uvicorn

    from app import app

    uvicorn.run(
        app,
        host=os.getenv("HOST", "0.0.0.0"),
        port=int(os.getenv("PORT", "8000")),
    )
