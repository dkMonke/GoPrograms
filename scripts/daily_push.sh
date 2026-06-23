#!/bin/bash
# daily_push.sh — Automatically stages, commits, and pushes all changes
# made during the day to the remote repository.
# Intended to be run via a daily cron job or launchd agent.

REPO_DIR="/Users/dinesh.sampath/GoPrograms"
LOG_FILE="$REPO_DIR/scripts/daily_push.log"
DATE=$(date '+%Y-%m-%d %H:%M:%S')

cd "$REPO_DIR" || { echo "[$DATE] ERROR: Cannot cd to $REPO_DIR" >> "$LOG_FILE"; exit 1; }

# Check if there are any changes to commit
if git diff --quiet && git diff --staged --quiet && [ -z "$(git ls-files --others --exclude-standard)" ]; then
    echo "[$DATE] No changes to commit." >> "$LOG_FILE"
    exit 0
fi

# Stage all changes (new, modified, deleted files)
git add -A

# Commit with a date-stamped message
COMMIT_MSG="Daily update: $(date '+%Y-%m-%d')"
git commit -m "$COMMIT_MSG" >> "$LOG_FILE" 2>&1

# Push to remote
git push origin main >> "$LOG_FILE" 2>&1

if [ $? -eq 0 ]; then
    echo "[$DATE] Successfully pushed daily changes." >> "$LOG_FILE"
else
    echo "[$DATE] ERROR: Push failed." >> "$LOG_FILE"
fi
