@echo off
echo Starting Cosmic Signature FAQ Bot Backend...
echo.

REM Activate virtual environment
if exist venv\Scripts\activate.bat (
    call venv\Scripts\activate.bat
) else (
    echo Error: Virtual environment not found!
    echo Please run: python -m venv venv
    echo Then run: venv\Scripts\activate
    echo Then run: pip install -r requirements.txt
    pause
    exit /b 1
)

echo Virtual environment activated
echo.

REM Run the FastAPI server
echo Starting FastAPI server on http://localhost:8000
echo.
python app.py
