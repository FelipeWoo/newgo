#!/usr/bin/env bash

SCRIPT_URL="https://raw.githubusercontent.com/FelipeWoo/newgo/refs/heads/main/scripts/make_repo_newgo.sh"
SCRIPT_PATH="$HOME/make_repo_newgo.sh"
ALIASES_FILE="$HOME/.bash_aliases"

curl -fsSL "$SCRIPT_URL" -o "$SCRIPT_PATH"
chmod +x "$SCRIPT_PATH"

touch "$ALIASES_FILE"

if ! grep -q "newgo()" "$ALIASES_FILE"; then
cat >> "$ALIASES_FILE" <<EOF

# newgo repo generator
newgo() {
    "$SCRIPT_PATH" "\$@"
}
EOF
fi

echo "Installed function: newgo"
echo "Run: source $ALIASES_FILE"
