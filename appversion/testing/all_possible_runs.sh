#!/usr/bin/env bash

main() {
    avtest
    avtest version
    avtest version oneliner
    avtest version dev
    avtest version dev-details
    avtest version dev oneliner
    avtest version debug-check
}

avtest() {
    echo "\n ......... now testing 'app $@'\n"
    go run main.go $@
}

main