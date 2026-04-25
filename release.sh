#!/usr/bin/env bash
set -euo pipefail

APP_NAME="daily-news"
BIN_DIR="bin"
DIST_DIR="release-dist"

usage() {
  cat <<'EOF'
Usage:
  ./release.sh <tag> [title]

Examples:
  ./release.sh v0.1.0
  ./release.sh v0.1.1 "Daily News v0.1.1"

Environment variables:
  RELEASE_NOTES_FILE   Optional. Path to release notes markdown file.
  RELEASE_NOTES_TEXT   Optional. Inline release notes text.
  SKIP_BUILD           Optional. Set to 1 to skip make build-all.

Notes:
  - Requires GitHub CLI (gh) and a logged-in session (gh auth login).
  - This script creates release assets from bin/* and publishes them to GitHub Releases.
EOF
}

if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
  usage
  exit 0
fi

if [[ $# -lt 1 ]]; then
  usage
  exit 1
fi

TAG="$1"
TITLE="${2:-$TAG}"

if ! command -v gh >/dev/null 2>&1; then
  echo "Error: gh not found. Install GitHub CLI first."
  exit 1
fi

if [[ "${SKIP_BUILD:-0}" != "1" ]]; then
  echo "==> Building cross-platform binaries"
  make build-all
fi

if [[ ! -d "$BIN_DIR" ]]; then
  echo "Error: $BIN_DIR directory not found."
  exit 1
fi

echo "==> Preparing release artifacts"
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

artifacts=()

for file in "$BIN_DIR"/*; do
  [[ -f "$file" ]] || continue
  base="$(basename "$file")"

  if [[ "$base" == *.exe ]]; then
    archive="$DIST_DIR/${base%.exe}.zip"
    (
      cd "$BIN_DIR"
      zip -q -j "../$archive" "$base"
    )
  else
    archive="$DIST_DIR/$base.tar.gz"
    tar -czf "$archive" -C "$BIN_DIR" "$base"
  fi

  artifacts+=("$archive")
  echo "  - $(basename "$archive")"
done

if [[ ${#artifacts[@]} -eq 0 ]]; then
  echo "Error: no artifacts were generated."
  exit 1
fi

echo "==> Publishing GitHub Release: $TAG"
gh_args=(
  release create "$TAG"
  --title "$TITLE"
)

if [[ -n "${RELEASE_NOTES_FILE:-}" ]]; then
  gh_args+=(--notes-file "$RELEASE_NOTES_FILE")
elif [[ -n "${RELEASE_NOTES_TEXT:-}" ]]; then
  gh_args+=(--notes "$RELEASE_NOTES_TEXT")
else
  gh_args+=(--generate-notes)
fi

gh "${gh_args[@]}" "${artifacts[@]}"

echo "==> Done. Release published for tag: $TAG"
