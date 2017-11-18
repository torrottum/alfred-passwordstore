#!/bin/bash
PATH=/usr/local/bin:$PATH

gpg --decrypt "$1" 2> /dev/null | head -n 1 | tr -d '\n' | pbcopy
