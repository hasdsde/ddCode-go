#!/bin/bash
path=$(cd `dirname $0`;pwd)
cd ${path}
echo $path

res=0

# 镜像名
image_name=$1
echo "image="${image_name}
sed -i "s@{IMAGE}@$image_name@g" docker-compose.yml

r=$?
res=`expr ${res} + ${r}`
echo "$res"

docker stop ddCode-server
docker rm ddCode-server

docker-compose up -d

# 执行部署
if [ $res -ne 0 ]; then
    exit 1
fi
