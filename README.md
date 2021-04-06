# File Rotation CLI tool

A tool to keep a fixed number of files in a directory, useful for keeping backups

## Downlaod
go get github.com/THasthika/frotate-go/cmd/frotate

## Run
go run github.com/THasthika/frotate-go/cmd/frotate

## Build
go build -o frotate github.com/THasthika/frotate-go/cmd/frotate

## Usage
./frotate [command]

commands:

    add - add a file to the rotation directory

example:

./frotate add test.zip -prefix=backup2 -directory="./test" -limit=20

output file: ./test/backup2-2020-01-01-01-01-00.zip
