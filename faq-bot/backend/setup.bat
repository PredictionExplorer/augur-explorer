@echo off
echo ========================================
echo Cosmic Signature FAQ Bot - Backend Setup
echo ========================================
echo.

REM Check if Python is installed
python --version >nul 2>&1
if errorlevel 1 (
    echo Error: Python is not installed or not in PATH
    echo Please install Python 3.9 or higher
    pause
    exit /b 1
)

echo Python found!
echo.

REM Create virtual environment
if not exist venv (
    echo Creating virtual environment...
    python -m venv venv
    echo Virtual environment created!
) else (
    echo Virtual environment already exists.
)

echo.

REM Activate virtual environment
echo Activating virtual environment...
call venv\Scripts\activate.bat

echo.

REM Upgrade pip
echo Upgrading pip...
python -m pip install --upgrade pip

echo.

REM Install dependencies
echo Installing dependencies...
pip install -r requirements.txt

echo.

REM Create .env file if it doesn't exist
if not exist .env (
    echo Creating .env file...
    copy .env.example .env
    echo .env file created!
) else (
    echo .env file already exists.
)

echo.
echo ========================================
echo Setup Complete!
echo ========================================
echo.
echo To start the backend server:
echo   1. Run: venv\Scripts\activate.bat
echo   2. Run: python app.py
echo.
echo Or simply run: start.bat
echo.
pause
