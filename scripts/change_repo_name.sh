#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ENV_FILE="${ROOT_DIR}/.env"
GO_MOD_FILE="${ROOT_DIR}/go.mod"

if [[ ! -f "${ENV_FILE}" ]]; then
  echo "Skipping rename: ${ENV_FILE} not found"
  exit 0
fi

if [[ ! -f "${GO_MOD_FILE}" ]]; then
  echo "Skipping rename: ${GO_MOD_FILE} not found"
  exit 0
fi

app_name="$(
  awk -F= '
    $1 == "APP_NAME" {
      value = substr($0, index($0, "=") + 1)
      gsub(/^[[:space:]]+|[[:space:]]+$/, "", value)
      gsub(/^"/, "", value)
      gsub(/"$/, "", value)
      print value
      exit
    }
  ' "${ENV_FILE}"
)"

if [[ -z "${app_name}" ]]; then
  echo "Skipping rename: APP_NAME is empty in .env"
  exit 0
fi

if [[ ! "${app_name}" =~ ^[A-Za-z0-9._/-]+$ ]]; then
  echo "Invalid APP_NAME: ${app_name}"
  echo "Use only letters, numbers, dots, underscores, slashes, or dashes."
  exit 1
fi

current_module="$(awk '/^module / { print $2; exit }' "${GO_MOD_FILE}")"

if [[ -z "${current_module}" ]]; then
  echo "Unable to determine current module name from go.mod"
  exit 1
fi

if [[ "${current_module}" == "${app_name}" ]]; then
  echo "Module already set to ${app_name}"
  exit 0
fi

echo "Renaming Go module: ${current_module} -> ${app_name}"

go -C "${ROOT_DIR}" mod edit -module "${app_name}"

while IFS= read -r file; do
  perl -0pi -e "s#${current_module}/#${app_name}/#g" "${file}"
done < <(find "${ROOT_DIR}" -type f -name '*.go' -not -path '*/vendor/*' | sort)

echo "Updated imports to ${app_name}"
