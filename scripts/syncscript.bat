@echo off
go build -o filesync.exe

if exist filesync.exe (
    filesync sync -f file1.txt,file2.txt
) else (
    filesync.exe
)