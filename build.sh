#!/bin/bash
cd "${0%/*}"
cp -r ./src/ ./dist/src/
./tmpl-preprocessor/tmpl-pre
./build-docker.sh
