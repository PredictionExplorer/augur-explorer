#!/bin/bash

echo "========================================"
echo "Cosmic Signature FAQ Bot - Backend Setup"
echo "========================================"
echo

# Check if Python is installed
if ! command -v python3 &> /dev/null; then
    echo "Error: Python 3 is not installed or not in PATH"
    echo "Please install Python 3.9 or higher"
    exit 1
fi

echo "Python found!"
echo

# Create virtual environment
if [ ! -d "venv" ]; then
    echo "Creating virtual environment..."
    python3 -m venv venv
    echo "Virtual environment created!"
else
    echo "Virtual environment already exists."
fi

echo

# Activate virtual environment
echo "Activating virtual environment..."
source venv/bin/activate

echo

# Upgrade pip
echo "Upgrading pip..."
python -m pip install --upgrade pip

echo

# Install dependencies
echo "Installing dependencies..."
pip install -r requirements.txt

echo

# Create .env file if it doesn't exist
if [ ! -f ".env" ]; then
    echo "Creating .env file..."
    cp .env.example .env
    echo ".env file created!"
else
    echo ".env file already exists."
fi

echo
echo "========================================"
echo "Setup Complete!"
echo "========================================"
echo
echo "To start the backend server:"
echo "  1. Run: source venv/bin/activate"
echo "  2. Run: python app.py"
echo
echo "Or simply run: ./start.sh"
echo

# Make start.sh executable
chmod +x start.sh
