#!/bin/bash

# Build the project
echo "Building the project..."
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Zipping the binary..."
zip get-all-books.zip main

echo "Removing the binary..."
