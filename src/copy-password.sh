#!/bin/bash
PATH=/usr/local/bin:$PATH

pass "$1" | head -n 1 | tr -d '\n' | pbcopy
