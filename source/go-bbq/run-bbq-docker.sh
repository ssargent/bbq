#!/bin/bash
for filename in ./secrets/*.txt; do
    secret=$(<$filename)
    fbname=$(basename "$filename" | cut -d. -f1)
    export $fbname="$secret"
  #      ./MyProgram.exe "$filename" "Logs/$(basename "$filename" .txt)_Log$i.txt"
     
done

 docker run -t -i -e BBQ_DB_HOST -e BBQ_DB_NAME -e BBQ_DB_PASSWORD -e BBQ_DB_USER -p 21337:21337 myfamilycooks.azurecr.io/bbq/go-bbq:latest