#!/bin/bash
for filename in ./secrets/*.txt; do
    secret=$(<$filename)
    fbname=$(basename "$filename" | cut -d. -f1)
    export $fbname="$secret"
  #      ./MyProgram.exe "$filename" "Logs/$(basename "$filename" .txt)_Log$i.txt"
     
done

./bbq-apiserver
