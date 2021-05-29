#!/bin/bash

# Directory structure:
# .
# └── money-manager
#     ├── build.sh
#     └── src
#         └── pakages
#             └── main.go
#
#
# The final structure compiled will look like:
#
# .
# └── money-manager
#     ├── bin
#     │    └── pakages
#     │       └── main
#     ├── build.sh
#     └── src
#         └── pakages
#             └── main.go
#

BASE_PATH=./src
OUT_PATH=./bin
compiledFiles=0

compile_folder_recurse() {
    for path in "$1"/*; do

        if [ -d "$path" ]; then
            dir=${path##*/}
            compile_folder_recurse "$path"
        elif [ -f "$path" ]; then
            extension="${path##*.}"

            if [ "$extension" = "go" ]; then
                echo "[$((compiledFiles+1))] Compiling: $path"
                env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $OUT_PATH/$dir/$(basename -- $path .go) $path
                compiledFiles=$((compiledFiles+1))
            fi
        fi
    done
}

compile_folder_recurse $BASE_PATH
echo ""
echo "Total compiled files: $compiledFiles"
