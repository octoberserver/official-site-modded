#!/bin/bash
cd "${0%/*}"
npx tailwindcss -i ./src/common/main-before.css -o ./src/common/main.css
cp -r ./src/ ./dist/src/
./tmpl-preprocessor/tmpl-pre
./build-docker.sh
