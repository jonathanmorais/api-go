#!/bin/bash
CURR_VER=`docker images api-go –format “{{.Tag}}” | head -1`
if [ “$CHANGE” == “MAJOR” ]; then
VER_UP=1.0
else
VER_UP=0.1
fi
NEW_VER=$(awk ‘{print $1+$2}’ <<<“$CURR_VER $VER_UP”)
docker build -t {image_name}:$NEW_VER ./
docker tag {image_name}:$NEW_VER {image_name}:latest