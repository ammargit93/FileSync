@echo off
go build -o ../filesync.exe

if exist filesync.exe (
    filesync dwld -dwld file1.txt,file2.txt
) else (
    filesync.exe
)