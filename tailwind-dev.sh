#!/bin/bash
cd "${0%/*}"
npx tailwindcss -i ./src/common/main-before.css -o ./src/common/main.css --watch
