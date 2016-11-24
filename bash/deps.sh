#!/bin/bash

function _deps() {
    _sysdeps
    _godeps
}

function _sysdeps() {
    case "$OSTYPE" in
    darwin*)
        local deps=
        _check fswatch || deps+=' fswatch'
        if [ -n "$deps" ]; then
            echo "Info: Found missing system dependencies."
            _echoRun brew update
            _echoRun brew install "$deps"
        fi
        ;;
    *)
        _check fswatch || echo 'Error' >&2 && return 1
        ;;
    esac
}

function _godeps() {
}

function _check() {
    if ! type "$1" &>/dev/null; then
        echo "$1 not installed"
        return 1
    fi
    echo "Found $1"
}

function _echoRun() {
    echo ">> $@"
    "$@"
}

_deps

