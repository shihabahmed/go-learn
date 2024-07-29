@echo off
setlocal

REM Check if a command was passed as an argument
if "%1"=="" (
    echo Usage: %0 [build|start|run|clean]
    exit /b 1
)

REM Define commands
if "%1"=="build" (
    echo Building the project...
    go build -o bin\ .\
) else if "%1"=="start" (
    echo Running the project...
    .\bin\go-learn
) else if "%1"=="run" (
    echo Running the project...
    reflex -r '\.go$' -s -- sh -c 'go run main.go'
) else if "%1"=="clean" (
    echo Cleaning build artifacts...
    if exist bin (
        del /Q bin\*
        rmdir bin
    ) else (
        echo No build artifacts to clean.
    )
) else (
    echo Unknown command: %1
    echo Usage: %0 [build|run|test|clean]
    exit /b 1
)

endlocal
