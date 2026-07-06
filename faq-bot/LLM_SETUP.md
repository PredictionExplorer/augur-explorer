# LLM Setup Guide - Get Better Answers!

The FAQ bot now supports **OpenAI GPT-4** and **Anthropic Claude** to generate clear, step-by-step answers instead of just showing code snippets.

## Why Use an LLM?

**Without LLM (basic mode):**
```
Based on the codebase, here's what I found:

**frontend:**
- In `components/BidForm.tsx`:
  function BidForm() { const handleSubmit = () => { ... }
```

**With LLM (intelligent mode):**
```
To place a bid on the Cosmic Signature website:

1. Visit the main website and navigate to the active game page
2. Look for the "Place Bid" button in the main interface
3. Click the button to open the bidding form
4. Enter your bid amount (must be higher than the current minimum)
5. Confirm the transaction in your wallet
6. Wait for the transaction to be confirmed on the blockchain

Note: You'll need to have your wallet connected and sufficient funds available.
```

## Setup Instructions

### Option 1: OpenAI (Recommended)

1. **Get an API Key**
   - Go to https://platform.openai.com/api-keys
   - Sign up or log in
   - Create a new API key
   - Copy the key (starts with `sk-`)

2. **Add to Environment File**
   ```bash
   cd backend
   nano .env  # or use any text editor
   ```

3. **Add this line:**
   ```
   OPENAI_API_KEY=sk-your-actual-key-here
   ```

4. **Install OpenAI Package**
   ```bash
   # Activate venv first
   source venv/bin/activate  # Linux/macOS
   # or
   venv\Scripts\activate     # Windows
   
   # Install
   pip install openai
   ```

5. **Restart Backend**
   ```bash
   python app.py
   ```

   You should see:
   ```
   INFO:llm_answer_generator:Using OpenAI for answer generation
   ```

### Option 2: Anthropic Claude

1. **Get an API Key**
   - Go to https://console.anthropic.com/
   - Sign up or log in
   - Create a new API key
   - Copy the key

2. **Add to Environment File**
   ```bash
   cd backend
   nano .env
   ```

3. **Add this line:**
   ```
   ANTHROPIC_API_KEY=your-key-here
   ```

4. **Install Anthropic Package**
   ```bash
   pip install anthropic
   ```

5. **Restart Backend**
   ```bash
   python app.py
   ```

### Option 3: No API Key (Basic Mode)

The bot will still work without an API key, but answers will be basic code snippets. You'll see:

```
INFO:llm_answer_generator:No LLM API key found. Using basic answer generation.
INFO:llm_answer_generator:For better answers, set OPENAI_API_KEY or ANTHROPIC_API_KEY in .env
```

## Cost Considerations

### OpenAI Pricing (GPT-4 Turbo)
- **Input:** ~$0.01 per 1K tokens
- **Output:** ~$0.03 per 1K tokens
- **Typical query:** $0.02 - $0.05 per question

**Example:** 100 questions ≈ $2-5

### Anthropic Pricing (Claude 3 Sonnet)
- **Input:** ~$0.003 per 1K tokens
- **Output:** ~$0.015 per 1K tokens
- **Typical query:** $0.01 - $0.03 per question

**Example:** 100 questions ≈ $1-3

### Free Tier
- OpenAI: $5 free credits for new accounts
- Anthropic: Limited free tier

## Testing

After setup, test with these questions:

```
"How can I bid on the website?"
"Where is information about prize distribution?"
"How do I connect my wallet?"
"What tokens can I use?"
```

You should get detailed, step-by-step answers!

## Troubleshooting

### "No LLM API key found"
- Check that you added the key to `backend/.env`
- Make sure there are no extra spaces
- Restart the backend server

### "Error using OpenAI: ..."
- Verify your API key is correct
- Check your OpenAI account has credits
- Ensure you have internet connection

### "Module not found: openai"
- Install: `pip install openai`
- Make sure venv is activated

## Example .env File

```env
PORT=8000
HOST=0.0.0.0
DATA_DIR=./data
LOG_LEVEL=INFO

# Add your OpenAI key here
OPENAI_API_KEY=sk-proj-xxxxxxxxxxxxxxxxxxxxx

# OR use Anthropic instead
# ANTHROPIC_API_KEY=sk-ant-xxxxxxxxxxxxxxxxxxxxx
```

## Advanced: Using Local Models

For free, local LLM options (coming soon):
- Ollama integration
- Llama 3 support
- Mistral support

Stay tuned for updates!

---

**Recommendation:** Start with OpenAI GPT-4 for best results. The improved answer quality is worth the small cost for most use cases.
