#!/usr/bin/env bash
alias invisible='go run main.go'
invisible add-noise < example/plain.txt > example/noised.txt
invisible encode -m 'Hello, World!' < example/plain.txt > example/embedded.txt
invisible decode < example/embedded.txt
