```markdown
# FileSync
A simple CLI for text processing and matching

## Description
FileSync is a command-line interface (CLI) for processing and matching text files. It provides a set of commands for counting words, characters, and frequencies, as well as synchronizing two text files.

## Installation
To install FileSync, simply run the following command:
```bash
git clone github.com/ammargit93/FileSync
```
## Usage
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

### Examples
* To count the number of words in a file using the `file` command: `filesync file -files path/to/file.txt -count`
* To synchronize two files using the `sync` command: `filesync sync -files file1.txt file2.txt`
* To download a file from the internet using the `dwld` command: `filesync dwld -dwld https://example.com/file.txt`
* To generate a rough readme.md command: `filesync md -md file1.py,file2.go,dir1`

## Contributing
If you'd like to contribute to FileSync, please fork the repository and submit a pull request.

## License
FileSync is licensed under the MIT License.

## Technologies/Dependencies
FileSync uses the following technologies and libraries:
* Go
* environment variables: https://github.com/joho/godotenv
* https://github.com/urfave/cli

## CLI Documentation
The following APIs are available in FileSync:
* `textutil.CountWords(filePaths []string) int`: Count the number of words in a file or files
* `textutil.CountChar(filePaths []string) int`: Count the number of characters in a file or files
* `textutil.CountFreq(filePaths []string) map[string]int`: Count the frequency of words in a file or files
* `textutil.FindMatchingWords(filePaths []string) []string`: Find common characters between two files
* `fileops.UpdateFile(filePath1, filePath2) error`: Synchronize two text files
* `fileops.DownloadFile(url string, path string) error`: Download a file from the internet
* `fileops.SpawnGoroutine(fileUrls []string, path string) error`: Download multiple files from the internet concurrently

## Note: This readme was generated by this cli's md command, there might be some inconsistencies.

