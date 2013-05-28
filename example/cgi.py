#!/usr/bin/env python3

import os

print("Content-Type: text/plain")
print()

for var in os.environ:
	print(var + "=" + os.environ[var])
