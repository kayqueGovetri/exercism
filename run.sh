#!/usr/bin/env bash

set -euo pipefail

usage() {
    echo "Uso: $0 <linguagem> <exercicio>"
    echo
    echo "Exemplos:"
    echo "  $0 python two-fer"
    echo "  $0 go lasagna"
    echo "  $0 rust hello-world"
    exit 1
}

[[ $# -eq 2 ]] || usage

LANGUAGE="$1"
EXERCISE="$2"

run_python() {
    (
        cd "python/$1"
        python3 -m pytest -o markers=task
    )
}

run_go() {
    (
        cd "go/$1"
        go test -v --bench=. --benchmem
    )
}

run_rust() {
    (
        cd "rust/$1"
        cargo test
    )
}

case "$LANGUAGE" in
    python)
        run_python "$EXERCISE"
        ;;
    go)
        run_go "$EXERCISE"
        ;;
    rust)
        run_rust "$EXERCISE"
        ;;
    *)
        echo "Linguagem não suportada: $LANGUAGE"
        exit 1
        ;;
esac