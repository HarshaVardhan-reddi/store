#!/bin/bash
commit_message="$1"
if [ -z "$commit_message" ]; then
  echo "Commit message is empty"
  exit 1
fi

echo -e "staging the current changes...\n"
git add .

echo -e "commiting the current change...\n"
git commit -m "$commit_message"

echo -e "Pushing origin under commit '${commit_message}...\n'"

# sleep 1
git push origin main


echo -e "\n 🎉Synced changes with origin!"

exit 0