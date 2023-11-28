#!/bin/bash

for i in {1..14}
do
    mkdir ${i} 
    cd ${i}
    
    touch "main.go"
    go mod init "task-${i}"
    echo "package main" > main.go

    cd ../
done


