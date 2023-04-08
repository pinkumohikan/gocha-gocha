#!/bin/bash -eu

EVENT_ID=phpconfukuoka-2023
MAX=100

for ((i = 1; i <= $MAX; i++)); do
    echo "fetching... page=$i"
    curl -sSf https://fortee.jp/$EVENT_ID/proposal/all\?page=$i > ./proposal-pages/$i.html 2>/dev/null || break
    sleep 2
done
