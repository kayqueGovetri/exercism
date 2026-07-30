#!/usr/bin/env bash

set -euo pipefail

TMP=$(mktemp)

{
    echo "| Language | Exercises |"
    echo "|----------|----------:|"

    for dir in */ ; do
        dir="${dir%/}"

        # Ignora pastas que não são linguagens
        case "$dir" in
            .git|.github|venv) continue ;;
        esac

        count=$(find "$dir" -mindepth 1 -maxdepth 1 -type d | wc -l)

        language="$(tr '[:lower:]' '[:upper:]' <<< "${dir:0:1}")${dir:1}"

        printf "| %s | %d |\n" "$language" "$count"
    done | sort
} > "$TMP"

awk -v table="$(cat "$TMP")" '
/<!-- EXERCISM_STATS_START -->/{
    print
    print table
    skip=1
    next
}
/<!-- EXERCISM_STATS_END -->/{
    skip=0
}
!skip
' README.md > README.new

mv README.new README.md
rm "$TMP"