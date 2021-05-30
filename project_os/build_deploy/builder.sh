#!/bin/bash

GOOS=linux
GOARCH=amd64

echo "Remote..."
ssh -i mygcp mygcp@34.87.6.44 'bash -s' < deploy.sh
