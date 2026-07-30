#!/usr/bin/env bash

set -euo pipefail

if [[ $# -lt 2 ]]; then
    echo "Uso: $0 <linguagem> <exercicio>"
    echo
    echo "Exemplos:"
    echo "  $0 rust hello-world"
    echo "  $0 python leap"
    echo "  $0 go two-fer"
    exit 1
fi

LANGUAGE="$1"
shift

EXERCISE="$*"

MESSAGE="feat(${LANGUAGE}): solve ${EXERCISE} exercise"
./update_readme.sh
git add .
git commit -m "$MESSAGE"

echo "✔ Commit criado:"
echo "  $MESSAGE"