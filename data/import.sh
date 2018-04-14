#!/bin/sh

for filename in *.json; do
    arangoimp --server.database autonet --server.password 123 \
        --file "$filename" --collection "$(basename "$filename" .json)" \
        --create-collection true
done

# arangoimp --file departments.json --server.database autonet --collection departments --create-collection true
# arangoimp --file teams.json --server.database autonet --collection teams --create-collection true
# arangoimp --file users.json --server.database autonet --collection users --create-collection true

