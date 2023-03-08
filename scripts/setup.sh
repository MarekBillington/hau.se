#!/bin/bash

# @todo eventually make the client and server build separately

# rebuild client and server
echo 'setting up client'
cd ~/www/hau.se/client
npm run build

echo 'setting up server'
cd ~/www/hau.se/server
go build -buildvcs=false -v

# ping them to check they are up...
echo 'kicking the whale'
docker restart client hause-server-1

now=$(date +"%T")
echo "Build Finised at: $now"