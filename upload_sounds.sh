#!/bin/bash

API_URL="https://samplevault.ru/api/v1/sounds/upload"

for file in *; do
    if [ -f "$file" ] && [ "$file" != "ab.sh" ]; then
        curl -F "audio=@$file" $API_URL -v
    fi
done
