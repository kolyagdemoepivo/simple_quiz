#!/bin/sh
echo "Generating proto"
./generate.sh
echo "Building server"
go build -C server
echo "Building client"
go build -C client
echo "Done"