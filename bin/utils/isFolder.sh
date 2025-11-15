#!/bin/bash
set -e

if [ -d "./build" ]; then
    echo "Folder: OK"
	
	if [ -d "./main" ]; then
		echo "old_build file"
	else
		echo "not old_build file"
	fi
else
	echo "Folder: NONE"
	echo "Create folder..."
	mkdir ./build/
fi

