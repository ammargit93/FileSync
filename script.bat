@echo off
REM Build the Go program
go build -o filesync.exe

REM Check if the build was successful
if exist filesync.exe (
    echo Build successful. Running the program...
    REM Run the compiled executable with arguments
    filesync f -f file1.txt -grep
) else (
    echo Build failed. Please check for errors.
)

REM Pause the script to view output before exiting
pause
