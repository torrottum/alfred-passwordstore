#!/bin/bash
PATH=/usr/local/bin:$PATH

gpg --decrypt "$HOME/.password-store/$1.gpg" 2> /dev/null | head -n 1 | tr -d '\n' | pbcopy
