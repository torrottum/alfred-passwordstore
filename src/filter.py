#!/usr/bin/env python3

import re
import os
import fnmatch
import sys
import json

def fuzzyfinder(user_input, collection):
    suggestions = []
    pattern = '.*?'.join(user_input)
    regex = re.compile(pattern)
    for item in collection:
        match = regex.search(item)
        if match:
            suggestions.append((len(match.group()), match.start(), item))
    return [x for _, _, x in sorted(suggestions)]

homepath = os.path.expanduser('~/.password-store') + '/'
items = []
for root, dirs, files in os.walk(homepath, topdown=False):
    for name in files:
        if fnmatch.fnmatch(os.path.join(root, name), '*.gpg'):
            path = os.path.join(root, name).replace(homepath, '')
            items.append(os.path.splitext(path)[0])

query = sys.argv[1].replace(' ', '/')

items = fuzzyfinder(query, items)
items = list(map(lambda path: {'title': path, 'arg': path}, items))

print(json.dumps({
    'items': items
}))
