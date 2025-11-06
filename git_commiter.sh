#!/bin/bash
commit_message="$1"
if [ -z "$commit_message" ]; then
  echo "Commit message is empty"
  exit 1
fi

echo "staging the current changes..."
git add .

echo "commiting the current change..."
git commit -m "$commit_message"

echo "Pushing origin under commit ${commit_message}..."

# sleep 1
git push origin main


echo "Synced changes with origin!"

exit 0