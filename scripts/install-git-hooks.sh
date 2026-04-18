#!/bin/sh
set -eu

git config core.hooksPath .githooks
git config commit.template .gitmessage

chmod +x .githooks/commit-msg

echo "Git hooks installed."
echo "core.hooksPath=$(git config --get core.hooksPath)"
echo "commit.template=$(git config --get commit.template)"
