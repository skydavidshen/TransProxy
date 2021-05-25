#!/bin/sh

nohup ./translate >> /tmp/translate.log &
nohup ./call-insert-trans >> /tmp/call-insert-trans.log

# 为了让docker不退出前台模式，让整个脚本一直处于死循环状态
while :
do
    date
    sleep 1
done