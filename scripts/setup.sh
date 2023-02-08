#!/bin/bash


# rebuild client and server
echo 'setting up client'
cd ../client
npm run build

echo 'setting up server'
cd ../server
go build -v

# teardown client and server containers and images

# start new client and server containers

# ping them to check they are up...

