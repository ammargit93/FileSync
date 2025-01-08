@echo off
go build -o filesync.exe
// filesync -to-md 
if exist filesync.exe (
    filesync f -f file1.txt,file2.txt -grep
) else (
    echo Build failed. Please check for errors.
    filesync.exe
)
