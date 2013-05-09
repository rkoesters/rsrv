rsrv is a simple web server written in [Go](http://golang.org/).

## Description

rsrv uses a configuration file named `rsrv.conf` to mount different 
handlers to the directory tree (see `example.config` for an example 
configuration file). The config file is parsed in a concurrent manner 
using Go's goroutines and channels. The type of each mount point 
defines how that mount point will behave (e.g. `dir` will serve files 
from a directory and `file` will just serve a certain file).

## Objectives

- Parse configuration file concurrently.
- Serve files over http.
- Mount different mount points using a ServeMux.
- Create a variety of different mount types.
