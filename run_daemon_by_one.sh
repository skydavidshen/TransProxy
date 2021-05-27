#!/bin/sh

nohup ./translate &
nohup ./call-insert-trans &

# 为了让docker不退出前台模式，让整个脚本一直处于死循环状态
echo "Enter a block..."
while :
do
    date
    sleep 1
done