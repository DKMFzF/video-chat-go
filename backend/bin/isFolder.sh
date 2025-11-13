#!/bin/bash
set -e

if [ -d "./build" ]; then
    echo "Folder: OK"
else
	echo "Folder: NONE"
	echo "Create folder..."
	mkdir ./build/
fi

