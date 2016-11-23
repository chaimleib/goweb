#!/bin/bash
# This file is to be source-d when entering this directory. It sets up
# project-specific environment variables, such as interpreter versions and
# import paths.
# https://github.com/chaimleib/repoactions

gitroot="$(git rev-parse --show-toplevel)"
export GOPATH="${gitroot}/src/go"
