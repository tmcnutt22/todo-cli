#!/bin/bash

echo "Building and installing todo app..."
sudo go build -o /usr/local/bin/todo
echo "Done! Todo is installed in /usr/local/bin"
