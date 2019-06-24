#!/bin/bash

GOPROXY=https://goproxy.io
go build -i -v
cd console && npm install && npm run build
cd ../theme && npm install && npm run build

echo 'build pipe done'
