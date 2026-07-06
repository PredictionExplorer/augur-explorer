#!/usr/bin/env bash
# Create or refresh the FAQ bot Python venv at $FAQ_BOT_VENV (see ~/configs/faq-bot.env).
set -euo pipefail

usage() {
  cat <<EOF
Usage: $(basename "$0") [--recreate]

Creates \$FAQ_BOT_VENV if missing, then installs/upgrades packages from backend/requirements.txt.
If the venv already exists, dependencies are synced (safe to re-run after git pull).

  --recreate   Remove and recreate the venv from scratch
EOF
}

RECREATE=false
while [[ $# -gt 0 ]]; do
  case "$1" in
    --recreate) RECREATE=true; shift ;;
    -h | --help) usage; exit 0 ;;
    *)
      echo "ERROR: Unknown option: $1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

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

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BACKEND="${SCRIPT_DIR}/../backend"
PIP="${FAQ_BOT_VENV}/bin/pip"
PYTHON="${FAQ_BOT_VENV}/bin/python3"

if ! command -v git >/dev/null 2>&1; then
  echo "ERROR: system 'git' is not installed." >&2
  echo "Install with: sudo apt install git python3-venv" >&2
  exit 1
fi

if [[ "${RECREATE}" == true && -d "${FAQ_BOT_VENV}" ]]; then
  echo "Removing existing venv: ${FAQ_BOT_VENV}"
  rm -rf "${FAQ_BOT_VENV}"
fi

if [[ ! -d "${FAQ_BOT_VENV}" ]]; then
  if ! python3 -m venv --help >/dev/null 2>&1; then
    echo "ERROR: python3 -m venv is unavailable." >&2
    echo "Install with: sudo apt install python3-venv python3-full" >&2
    exit 1
  fi
  echo "Creating venv at ${FAQ_BOT_VENV} ..."
  python3 -m venv "${FAQ_BOT_VENV}"
else
  echo "Venv exists: ${FAQ_BOT_VENV}"
  echo "Syncing dependencies from requirements.txt ..."
fi

"${PIP}" install --upgrade pip
"${PIP}" install -r "${BACKEND}/requirements.txt"
"${PYTHON}" -c "import git; print('GitPython OK:', git.__version__)"

echo "Done."
echo "  Python: ${PYTHON}"
echo "  Generate KB: ${SCRIPT_DIR}/generate-knowledge.sh"
echo "  Start backend: ${SCRIPT_DIR}/start-backend.sh"
echo "To recreate venv from scratch: $0 --recreate"
