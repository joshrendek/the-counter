#!/bin/sh
if [ "$(curl the-counter.the-counter:80 | jq '.count')" -gt "0" ]; then
    echo 'success'
else
    echo 'fail'
    exit 1
fi
