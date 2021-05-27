#!/bin/sh

job=${1:-""}
dir=${2:-""}

if [ $job == "daemon-call" ]; then
  go build -o bin/job $dir/daemon/call-insert-trans/main.go
else
  go build -o bin/job $dir/daemon/translate/main.go
fi


