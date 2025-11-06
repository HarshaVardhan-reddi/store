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

echo -e "\nPushing origin under commit \033[1;33m${commit_message}\033[0m\n"

# sleep 1
git push origin main


echo -e "\n\033[32m🎉 Synced changes with origin!\033[0m"

exit 0