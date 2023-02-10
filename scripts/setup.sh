#!/bin/bash


# rebuild client and server
echo 'setting up client'
cd ../client
npm run build

echo 'setting up server'
cd ../server
go build -v

# try with mounts to reload the container??
# Otherwise we need to do the full rebuild route...

# ping them to check they are up...
echo 'kicking the whale'
docker restart client hause-server-1