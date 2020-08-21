#!/bin/sh
echo  wget -q -O- http://${GO_SERVICE_URL}:${GO_SERVICE_PORT};
while [ true ]; do    wget -q -O- http://${GO_SERVICE_URL}:${GO_SERVICE_PORT};done

export GO_SERVICE_URL=35.247.196.141