#!/bin/bash
set -e

TEMPLATE_URL="https://github.com/FelipeWoo/newgo.git"

if [ -z "$1" ]; then
  echo "Usage: newgo <project_name>"
  exit 1
fi

PROJECT_NAME=$1

git clone --depth 1 "$TEMPLATE_URL" "$PROJECT_NAME"
cd "$PROJECT_NAME" || exit 1

rm -rf .git bin logs coverage.out

git init

mkdir -p .private
cat <<EOF > .private/notes.md
---
created_at: $(date +%Y-%m-%d)
tags: [tag_1, tag_2]
---

# Title

## Context
This file contains personal notes related to the project.

## Observations
- Write anything relevant here.

## To Do
- [ ] Define initial steps
- [ ] Document decisions

EOF

cat <<EOF > .env
APP_NAME=$PROJECT_NAME
APP_ENV=development
LOG_LEVEL=DEBUG
APP_PORT=8000
EOF

grep -qxF ".private/" .gitignore || echo ".private/" >> .gitignore
grep -qxF ".env" .gitignore || echo ".env" >> .gitignore

echo "Project '$PROJECT_NAME' created successfully."
echo "Next steps:"
echo "  cd $PROJECT_NAME"
echo "  make init"
echo "  git add ."
echo "  git commit -m 'Initial project files $PROJECT_NAME'"
