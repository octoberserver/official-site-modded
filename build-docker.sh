#!/bin/bash
cd "${0%/*}"
cd "dist"

read -p 'Deploy to remote?: ' remote
if [ "$remote" = "y" ]; then
    docker context use remote
else
    docker context use default
fi

docker build -t octmodsrv/official-site:latest .
docker run -p80:80 -p443:443 octmodsrv/official-site:latest

if [ "$remote" = "y" ]; then
    docker context use default
    echo "reseted docker context to default"
else
    docker context use remote
    echo "reseted docker context to remote"
fi
