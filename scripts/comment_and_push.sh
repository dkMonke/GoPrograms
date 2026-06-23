#!/bin/bash
# comment_and_push.sh — Auto-document then push.
#
# For every new or edited .go file in the repo, this script asks the local
# `claude` CLI to add in-depth doc comments describing each function (without
# changing any logic), verifies the result still builds, then hands off to the
# existing daily_push.sh to commit and push to the remote.
#
# Intended to be run via the same daily cron job / launchd agent as daily_push.sh.

set -uo pipefail

REPO_DIR="/Users/dinesh.sampath/GoPrograms"
SCRIPT_DIR="$REPO_DIR/scripts"
LOG_FILE="$SCRIPT_DIR/comment_and_push.log"
CLAUDE_BIN="$(command -v claude || echo "/Users/dinesh.sampath/.local/bin/claude")"
DATE() { date '+%Y-%m-%d %H:%M:%S'; }

log() { echo "[$(DATE)] $*" >> "$LOG_FILE"; }

cd "$REPO_DIR" || { log "ERROR: Cannot cd to $REPO_DIR"; exit 1; }

if [ ! -x "$CLAUDE_BIN" ]; then
    log "ERROR: claude CLI not found at '$CLAUDE_BIN'. Skipping commenting, pushing as-is."
    exec "$SCRIPT_DIR/daily_push.sh"
fi

# Collect new (untracked) and edited (modified, staged or unstaged) .go files.
# git status --porcelain emits a 2-char status field; trim it and keep .go paths.
# Use a read loop (not mapfile) for compatibility with the system bash 3.2.
GO_FILES=()
while IFS= read -r line; do
    [ -n "$line" ] && GO_FILES+=("$line")
done < <(git status --porcelain | awk '{ $1=""; sub(/^ +/,""); print }' | grep -E '\.go$')

if [ "${#GO_FILES[@]}" -eq 0 ]; then
    log "No new/edited .go files to document."
else
    log "Found ${#GO_FILES[@]} changed .go file(s): ${GO_FILES[*]}"

    PROMPT='Add in-depth Go doc comments to this file: %s

Rules:
- Add a doc comment immediately above every function, method, and type that lacks one. Follow Go convention: start the comment with the identifier name (e.g. "// CreateSession stores ...").
- Each comment should explain what the function does, its parameters, return values, and any notable behavior or side effects.
- If a file-level header comment is missing, add a concise one describing the file.
- DO NOT change, reorder, or remove any code, logic, imports, or formatting. Only add comments.
- Leave existing comments intact unless they are clearly wrong; improve them only if needed.
- Make the edits directly to the file. Do not output explanations.'

    for f in "${GO_FILES[@]}"; do
        if [ ! -f "$f" ]; then
            log "SKIP: '$f' no longer exists."
            continue
        fi

        log "Documenting $f ..."
        # shellcheck disable=SC2059
        if printf "$PROMPT" "$f" | "$CLAUDE_BIN" -p \
                --permission-mode acceptEdits \
                --allowedTools "Read Edit" \
                >> "$LOG_FILE" 2>&1; then
            log "Documented $f"
        else
            log "WARN: claude failed on '$f'; leaving file unchanged."
        fi

        # Safety gate (per file): gofmt parses the file and fails on any syntax
        # error, so it verifies our comment edits are valid Go. It also rewrites
        # the file in canonical format. A module-wide `go build` is NOT used here
        # because this learning repo intentionally has multiple `func main()` per
        # directory, which always fails to build as a module.
        if command -v gofmt >/dev/null 2>&1; then
            if gofmt -w "$f" 2>>"$LOG_FILE"; then
                log "gofmt OK: $f"
            else
                log "ERROR: gofmt failed on '$f' after commenting. NOT pushing — please review."
                exit 1
            fi
        fi
    done
fi

# Hand off to the existing push script.
log "Handing off to daily_push.sh"
exec "$SCRIPT_DIR/daily_push.sh"
