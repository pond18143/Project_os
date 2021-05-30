#!/bin/bash

sleep 2

echo "Remove Container..."
sudo docker rm -f project_os
sleep 2

echo "Remove Images..."
sudo docker rmi -f pondlamat/project_os
sleep 2

echo "Pull Images..."
sudo docker pull pondlamat/project_os
sleep 2

echo "Docker run"
sudo docker run -d --name project_os -p 8080:8080 pondlamat/project_os

echo "Finish"

