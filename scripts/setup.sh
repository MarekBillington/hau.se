#!/bin/bash

# @todo eventually make the client and server build separately

# rebuild client and server
echo 'setting up client'
cd ../client
npm run build

echo 'setting up server'
cd ../server
go build -v

# ping them to check they are up...
echo 'kicking the whale'
docker restart client hause-server-1
