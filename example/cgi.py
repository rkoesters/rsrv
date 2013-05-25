#!/usr/bin/env python3

import os

print("Content-Type: text/plain")
print()

print("QUERY_STRING: " + os.getenv("QUERY_STRING"))
print("PATH_INFO: " + os.getenv("PATH_INFO"))

