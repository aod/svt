#!/usr/bin/env bash

for x in `svt --algorithms`; do
    echo $x;
    sleep 1;
    svt -q -a 20 -t 3 -s $x; 
done
