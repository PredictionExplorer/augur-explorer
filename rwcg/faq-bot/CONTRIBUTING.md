# Contributing to Cosmic Signature FAQ Bot

Thank you for your interest in contributing! This document provides guidelines for contributing to the project.

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue with:
- Clear description of the bug
- Steps to reproduce
- Expected vs actual behavior
- Environment details (OS, Python version, Node version)
- Error messages or logs

### Suggesting Enhancements

Enhancement suggestions are welcome! Please include:
- Clear description of the enhancement
- Use case or problem it solves
- Proposed implementation (if you have ideas)

### Pull Requests

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Follow the code style guidelines
   - Add tests if applicable
   - Update documentation

4. **Test your changes**
   ```bash
   # Backend tests
   cd backend
   pytest

   # Frontend lint
   npm run lint
   ```

5. **Commit your changes**
   ```bash
   git commit -m "Add: brief description of your changes"
   ```

6. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Open a Pull Request**
   - Provide clear description
   - Reference any related issues
   - Explain what you changed and why

## Development Guidelines

### Code Style

**Python (Backend):**
- Follow PEP 8
- Use type hints
- Add docstrings to functions and classes
- Maximum line length: 100 characters

**TypeScript/React (Frontend):**
- Use TypeScript for type safety
- Follow React best practices
- Use functional components with hooks
- Use meaningful variable names

### Project Structure

```
faq/
├── app/                    # Next.js frontend
│   ├── components/        # React components
│   └── ...
├── backend/               # Python backend
│   ├── app.py            # FastAPI app
│   ├── rag_pipeline.py   # Haystack RAG
│   └── ...
└── ...
```

### Adding New Features

#### Adding a New Frontend Component

1. Create component in `app/components/`
2. Create corresponding CSS module
3. Import and use in pages
4. Add TypeScript types

Example:
```typescript
// app/components/NewComponent.tsx
'use client'

import styles from './NewComponent.module.css'

interface NewComponentProps {
  title: string
}

export default function NewComponent({ title }: NewComponentProps) {
  return <div className={styles.container}>{title}</div>
}
```

#### Adding a New API Endpoint

1. Add endpoint in `backend/app.py`
2. Add corresponding types/models
3. Update documentation
4. Add error handling

Example:
```python
@app.get("/api/new-endpoint")
async def new_endpoint():
    """Description of what this endpoint does"""
    try:
        # Your logic here
        return {"status": "success"}
    except Exception as e:
        logger.error(f"Error: {e}")
        raise HTTPException(status_code=500, detail=str(e))
```

#### Extending the RAG Pipeline

To add new document types or improve retrieval:

1. Update `SUPPORTED_EXTENSIONS` in `github_fetcher.py`
2. Modify `_process_repository()` if needed
3. Update answer generation logic in `answer_generator.py`
4. Test with sample queries

### Testing

#### Backend Testing

Create tests in `backend/tests/`:

```python
import pytest
from app import app

def test_health_endpoint():
    # Your test here
    pass
```

Run tests:
```bash
cd backend
pytest
```

#### Frontend Testing

```bash
npm run lint
```

### Documentation

- Update README.md for major changes
- Update SETUP.md for installation changes
- Add inline comments for complex logic
- Update API documentation

### Commit Messages

Use clear, descriptive commit messages:

- `Add: new feature or file`
- `Fix: bug fix`
- `Update: changes to existing feature`
- `Refactor: code restructuring`
- `Docs: documentation updates`
- `Style: formatting changes`

Examples:
```
Add: support for .cpp files in indexer
Fix: CORS error on frontend API calls
Update: improve answer generation for bidding questions
Docs: add troubleshooting section to README
```

## Areas for Contribution

### High Priority

- [ ] Add LLM integration (OpenAI, Anthropic, etc.)
- [ ] Implement embedding-based retrieval
- [ ] Add caching for common queries
- [ ] Improve error handling
- [ ] Add unit tests

### Medium Priority

- [ ] Add user authentication
- [ ] Implement chat history
- [ ] Add export chat functionality
- [ ] Improve mobile UI
- [ ] Add dark mode

### Nice to Have

- [ ] Add more customization options
- [ ] Support for more file types
- [ ] Real-time updates when repos change
- [ ] Analytics dashboard
- [ ] Multi-language support

## Questions?

Feel free to:
- Open an issue for discussion
- Reach out to maintainers
- Check existing issues and PRs

## License

By contributing, you agree that your contributions will be licensed under the project's MIT License.

---

Thank you for contributing to make this project better! 🚀
