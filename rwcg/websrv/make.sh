#!/bin/bash
# Build websrv binary in rwcg/websrv/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Building websrv..."
go build -o websrv .

if [ $? -eq 0 ]; then
	echo "Done. Binary: $SCRIPT_DIR/websrv"
else
	echo "Build failed."
	exit 1
fi
