#!/bin/bash
directory="api/openapi"
for file in "$directory"/*
do
    if [ -f "$file" ]; then
        filename=$(basename "$file")
        filename_no_ext="${filename%.*}"
        ./scripts/openapi-http.sh "$filename_no_ext" src/"$filename_no_ext"/infrastructure/transport transport 
    fi
done