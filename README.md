## FileSync
A simple CLI for text processing and matching

## Description
FileSync is a command-line interface (CLI) for processing and matching text files. It provides a set of commands for counting words, characters, and frequencies, as well as synchronizing two text files.

## Installation
To install FileSync, simply run the following command:
```bash
git clone github.com/ammargit93/FileSync
```
### Usage
### Commands
FileSync provides the following commands:
* `file`: Process text files using the following flags:
        + `-files` or `-f`: Specify one or more file paths
        + `-count` or `-c`: Count the number of words in the file(s)
        + `-cchar` or `-cch`: Count the number of characters in the file(s)
        + `-cfreq` or `-cf`: Count the frequency of words in the file(s)
        + `-grep` or `-gp`: Find common characters between the file(s)
* `sync`: Synchronize two text files using the following flag:
        + `-files` or `-f`: Specify the paths to the two files to synchronize
* `dwld`: Download files from the internet using the following flag:
        + `-dwld` or `-dw`: Specify the URL(s) of the files to download
* `md`: Generate markdown files for code files using the following flag:
        + `-to-md` or `-md`: Specify the path(s) to the code files to generate markdown for


