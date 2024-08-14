#!/usr/bin/env bash

if [ "$1" == "--remove" ]; then
    sudo rm -v /usr/local/bin/stamper
    exit 0
fi

go build
sudo mv -v stamper /usr/local/bin