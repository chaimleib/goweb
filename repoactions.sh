#!/bin/bash
# This file is to be source-d when entering this directory. It sets up
# project-specific environment variables, such as interpreter versions and
# import paths.
# https://github.com/chaimleib/repoactions

gitroot="$(git rev-parse --show-toplevel)"
BASH_LIBS="${gitroot}/bash" source "${gitroot}/bash/pathfuncs.sh"

# import path for golang
#prependPath "${gitroot}/src/go" GOPATH
# make sure to erase any other import paths:
export GOPATH="${gitroot}/src/go"

# put our golang tools in our PATH
prependPath "${GOPATH}/bin"

