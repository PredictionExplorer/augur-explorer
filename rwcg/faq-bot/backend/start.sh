#!/bin/bash

echo "Starting Cosmic Signature FAQ Bot Backend..."
echo

# Activate virtual environment
if [ -f "venv/bin/activate" ]; then
    source venv/bin/activate
else
    echo "Error: Virtual environment not found!"
    echo "Please run: python -m venv venv"
    echo "Then run: source venv/bin/activate"
    echo "Then run: pip install -r requirements.txt"
    exit 1
fi

echo "Virtual environment activated"
echo

# Run the FastAPI server
echo "Starting FastAPI server on http://localhost:8000"
echo
python app.py
