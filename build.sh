#!/bin/bash

go build -i -v
cd console && npm run build
cd .. && cd theme && gulp

echo 'build pipe done'
