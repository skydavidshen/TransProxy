#!/bin/bash

tag=$1

# 如果job参数不传入，则默认设置成空字符
job=${2:-""}

# rename target tag
target_tag=${3:-""}

repo=""
docker=""

if [ $job == "daemon-trans" ] || [ $job == "daemon-call" ]; then
  repo="844827150/trans-proxy-${job}"
  docker="daemon_by_one.Dockerfile"
  docker build -t "$repo:$tag" --build-arg JOB=$job -f "$docker" .
else
  repo="844827150/trans-proxy-web-server"
  docker="webserver.Dockerfile"
  docker build -t "$repo:$tag" -f "$docker" .
fi

if [ $target_tag == "" ]; then
    target_tag="latest"
fi

docker tag "$repo:$tag" "$repo:$target_tag" && docker push "$repo:$tag" && docker push "$repo:${target_tag}"

echo "build successfully."



# 需要提前登录 docker login
# 默认为: web
#sh build.sh {tag}

# 脚本执行 全路径
#sh build.sh {tag} daemon-trans/daemon-call/web release/latest/develop/master
