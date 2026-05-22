#!/usr/bin/env bash
# Create the FAQ bot Python venv at $FAQ_BOT_VENV (see ~/configs/faq-bot.env).
set -euo pipefail

CONFIG="${FAQ_BOT_ENV:-${HOME}/configs/faq-bot.env}"
if [[ -f "${CONFIG}" ]]; then
  # shellcheck disable=SC1090
  source "${CONFIG}"
fi

if [[ -z "${FAQ_BOT_VENV:-}" ]]; then
  echo "ERROR: FAQ_BOT_VENV is not set." >&2
  echo "Add to ${CONFIG}:" >&2
  echo "  export FAQ_BOT_VENV=/path/to/venv" >&2
  exit 1
fi

if [[ -d "${FAQ_BOT_VENV}" ]]; then
  echo "Venv already exists: ${FAQ_BOT_VENV}"
  echo "To recreate: rm -rf \"${FAQ_BOT_VENV}\" && $0"
  exit 0
fi

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BACKEND="${SCRIPT_DIR}/../backend"

echo "Creating venv at ${FAQ_BOT_VENV} ..."
python3 -m venv "${FAQ_BOT_VENV}"
# shellcheck disable=SC1091
source "${FAQ_BOT_VENV}/bin/activate"
pip install --upgrade pip
pip install -r "${BACKEND}/requirements.txt"
echo "Done. Activate with: source ${FAQ_BOT_VENV}/bin/activate"
