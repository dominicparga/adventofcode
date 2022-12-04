#!/usr/bin/env sh

project_root_dirpath="$(dirname "$(realpath "${0}")")/.."

go build \
    -o "${project_root_dirpath}/bin/adventofcode" \
    "${project_root_dirpath}/main.go" \
