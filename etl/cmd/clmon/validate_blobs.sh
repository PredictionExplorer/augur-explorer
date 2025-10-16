#!/usr/bin/env bash
# Usage: validate_blobs.sh <ARCHIVE_DIR> <BEACON1> [BEACON2]

ARCHIVE_DIR="${1:-}"
BEACON1="${2:-}"
BEACON2="${3:-}"

if [ -z "$ARCHIVE_DIR" ] || [ -z "$BEACON1" ]; then
  echo "usage: $0 <ARCHIVE_DIR> <BEACON1> [BEACON2]" >&2
  exit 1
fi

if ! command -v jq >/dev/null 2>&1; then
  echo "jq is required" >&2
  exit 1
fi

blob_size=131072
errors=0

get_commitment_count () {
  local s="$1" n=""
  n="$(curl -sS "$BEACON1/eth/v2/beacon/blocks/$s" \
      | jq -r '[.data.message.body.execution_payload.blob_kzg_commitments?, .data.message.body.blob_kzg_commitments?] | add // [] | length' 2>/dev/null)"
  if [ -z "$n" ] || [ "$n" = "null" ]; then
    n="$(curl -sS "$BEACON1/eth/v1/beacon/blocks/$s" \
        | jq -r '[.data.message.body.execution_payload.blob_kzg_commitments?, .data.message.body.blob_kzg_commitments?] | add // [] | length' 2>/dev/null)"
  fi
  if { [ -z "$n" ] || [ "$n" = "null" ]; } && [ -n "$BEACON2" ]; then
    n="$(curl -sS "$BEACON2/eth/v2/beacon/blocks/$s" \
        | jq -r '[.data.message.body.execution_payload.blob_kzg_commitments?, .data.message.body.blob_kzg_commitments?] | add // [] | length' 2>/dev/null)"
    if [ -z "$n" ] || [ "$n" = "null" ]; then
      n="$(curl -sS "$BEACON2/eth/v1/beacon/blocks/$s" \
          | jq -r '[.data.message.body.execution_payload.blob_kzg_commitments?, .data.message.body.blob_kzg_commitments?] | add // [] | length' 2>/dev/null)"
    fi
  fi
  [ -z "$n" ] || [ "$n" = "null" ] && n=0
  echo "$n"
}

have_dirs="$(find "$ARCHIVE_DIR" -maxdepth 1 -type d -printf '%f\n' | grep -E '^[0-9]+$' | sort -n)"
if [ -z "$have_dirs" ]; then
  echo "no numeric slot folders in $ARCHIVE_DIR" >&2
  exit 1
fi

first_slot="$(echo "$have_dirs" | head -n1)"
last_slot="$(echo "$have_dirs" | tail -n1)"

echo "=== RANGE ==="
echo "first: $first_slot"
echo "last : $last_slot"
echo

# Build gap list between first and last
gaps="$(echo "$have_dirs" \
  | awk 'NR==1{prev=$1; next} { for(i=prev+1;i<$1;i++) print i; prev=$1 }')"

echo "=== VALIDATING PRESENT SLOTS ==="
# Validate slots you already archived: blob count and file sizes
while IFS= read -r s; do
  [ -z "$s" ] && continue
  n="$(get_commitment_count "$s")"
  if [ "$n" -eq 0 ]; then
    # Slot has zero blobs on-chain; ensure you didn't accidentally store any
    if ls "$ARCHIVE_DIR/$s"/slot_*_index_*.blob >/dev/null 2>&1; then
      echo "WARN slot $s: chain has 0 blobs but files exist in archive/"
    fi
    continue
  fi

  # Count files you have
  have_count="$(ls "$ARCHIVE_DIR/$s"/slot_*_index_*.blob 2>/dev/null | wc -l | tr -d ' ')"
  if [ "$have_count" -ne "$n" ]; then
    echo "MISMATCH slot $s: chain=$n files=$have_count"
    errors=$((errors+1))
  fi

  # Size check
  for f in "$ARCHIVE_DIR/$s"/slot_*_index_*.blob; do
    [ -e "$f" ] || continue
    size="$(stat -c%s "$f" 2>/dev/null || stat -f%z "$f")"
    if [ "$size" != "$blob_size" ]; then
      echo "BADSIZE $f -> $size (want $blob_size)"
      errors=$((errors+1))
    fi
  done
done <<< "$have_dirs"

echo
echo "=== CHECKING GAPS THAT ACTUALLY HAVE BLOBS ==="
missing=0
while IFS= read -r s; do
  [ -z "$s" ] && continue
  n="$(get_commitment_count "$s")"
  if [ "$n" -gt 0 ]; then
    echo "MISSING slot $s ($n blobs)"
    missing=$((missing+1))
    errors=$((errors+1))
  fi
done <<< "$gaps"

echo
echo "=== SUMMARY ==="
echo "range     : $first_slot..$last_slot"
echo "dirs kept : $(echo "$have_dirs" | wc -l | tr -d ' ')"
echo "gaps total: $(echo "$gaps" | wc -l | tr -d ' ')"
echo "missing w/ blobs: $missing"
echo "errors    : $errors"

exit $([ "$errors" -gt 0 ] && echo 2 || echo 0)

