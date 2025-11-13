#!/bin/bash

if [ -d "./build" ]; then
	echo "delete builds...."
	sudo rm -rf ./build/main*
else
	echo "build folder not found"
fi

