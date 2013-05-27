rsrv
====

rsrv is a simple web server written in [Go](http://golang.org/).

Description
-----------

rsrv uses a configuration file named `rsrv.ini` to mount different 
handlers to the directory tree (see `example/rsrv.ini` for an example 
configuration file). The config file is parsed in a concurrent manner 
using Go's goroutines and channels. The type of each mount point 
defines how that mount point will behave (e.g. `dir` will serve files 
from a directory and `file` will just serve a certain file).

Objectives
----------

- Parse configuration file concurrently.
- Serve files over http.
- Mount different mount points to a single directory tree.
- Create a variety of different mount types.

Where to start
--------------

If you want to read through or modify rsrv, here are some important files:

- `example/*.ini`: The comments in these files are fantastic.
- `main.go`:	This is where the program starts.
- `config.go`:	This is where the configuration is parsed.
- `handler.go`:	This has `getHandler` and some helper functions that are
  shared by the handlers.
