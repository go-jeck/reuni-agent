#!/bin/bash
i=1
while true 
do 
    echo "$i: $TEST_CONFIG $1"
    i=$((i+1))
    sleep 1
done
