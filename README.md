# file-filter
Golang program to filter different types of lines from files

## Purpose
To filter files from various conditions, including:
* Duplicate lines, to cut down on wasted space
* Code lines, which you might want to remove from plain text files

## Status
Ready to use

## Installation
This program is written in Google Go language. Make sure that Go is installed and the GOPATH is set up as described in
[How to Write Go Code](https://golang.org/doc/code.html).

The install this program and its dependencies by running:

    go get github.com/BluntSporks/file-filter

## Usage
The program runs in one of two modes, controlled by the type flag.

Usage:

    file-filter [-type=(dupes|code)] FILENAME

Options:

    -type=(dupes|code)  Name of type to check [default: dupes]
