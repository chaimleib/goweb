#!/bin/bash
# This file is to be source-d when entering this directory. It sets up
# project-specific environment variables, such as interpreter versions and
# import paths.
# https://github.com/chaimleib/repoactions

gitroot="$(git rev-parse --show-toplevel)"

# import path for golang
export GOPATH="${gitroot}/src/go"

# put our golang tools in our PATH
BASH_LIBS="${gitroot}/bash" source bash/pathfuncs.sh
prependPath "${GOPATH}/bin"

