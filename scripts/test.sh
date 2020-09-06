#!/bin/bash

# https://unix.stackexchange.com/a/452142
if [ "${CI+1}" ]
then
    echo "CI environment"
    $UPLOAD_STORAGE="/tmp"
else
    echo "Local environment"
    . ./.env
fi

echo $UPLOAD_STORAGE
echo "\n"
UPLOAD_STORAGE=$UPLOAD_STORAGE go test ./...