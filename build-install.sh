#!/bin/bash

echo "Building and installing todo app..."
sudo go build -o /usr/local/bin/myapp
echo "Done! Todo is installed in /usr/local/bin"
