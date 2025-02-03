#!/bin/bash

go test ./... >> /dev/null && go build -o ~/bin/xocket
echo "xocket installed at $HOME/bin/xocket, please ensure this directory is added to path"