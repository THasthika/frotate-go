# File Rotation CLI tool

## Run
go run github.com/tharindu96/frotate-go/cli

## Build
go build -o frotate github.com/tharindu96/frotate-go/cli

## Usage
./frotate [command]
commands:
    add - add a file to the rotation directory

example:
./frotate add test.zip -prefix=backup2 -directory="./test" -limit=20
output file: ./test/backup2-2020-01-01-01-01-00.zip