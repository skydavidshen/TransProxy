#!/bin/bash

tag=$1

# 如果job参数不传入，则默认设置成空字符
job=${2:-""}

# rename target tag
target_tag=${3:-""}

repo=""
docker=""

if [ $job == "daemon" ]; then
  repo="844827150/trans-proxy-${job}"
  docker="daemon.Dockerfile"
else
  repo="844827150/trans-proxy-web-server"
  docker="webserver.Dockerfile"
fi

if [ $target_tag == "" ]; then
    target_tag="latest"
fi

docker build -t "$repo:$tag" -f "$docker" . && docker tag "$repo:$tag" "$repo:$target_tag"
docker push "$repo:$tag" && docker push "$repo:${target_tag}"

echo "build successfully."



# 需要提前登录 docker login
# 默认为: web
# sh build.sh {tag}
# sh build.sh 1.1.2 daemon release
# sh build.sh 1.1.2 web release

# 脚本执行 全路径
#sh build.sh {tag} {service} {branch}
#sh build.sh {tag} daemon/web release/latest/develop/master

# 执行应用程序, 编译完成之后，执行如下:
# ./main {env}        # env: dev|prod

