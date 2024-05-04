#!/bin/bash

API_URL="https://samplevault.ru/api/v1/sounds/upload"

for file in *; do
    if [ -f "$file" ] && [ "$file" != "ab.sh" ]; then
      curl -X PUT "https://samplevault.ru/sounds/$file" --upload-file $file -v
    fi
done
