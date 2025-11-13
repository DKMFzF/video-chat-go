#!/bin/build

echo "test..1"

if [ -d "./build" ]; then
	echo "delete builds...."
	
	echo "test..2"

	sudo rm -rf ./build/main*
else
	echo "build folder not found"
fi

