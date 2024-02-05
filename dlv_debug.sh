#!/bin/bash
dlvpid=`lsof -i:2345 | tail -n1 | awk '{print $2}'`
echo "查询到dlv的pid为:$dlvpid"
if [ -n "$dlvpid" ]; then
  kill -9 $dlvpid
fi

sleep 1s

propid=`lsof -i:8088 | tail -n1 | awk '{print $2}'`
echo "查询到go运行的pid为$propid"
if [ -n "$propid" ]; then
  kill -9 $propid
fi
if [ "$1" == "clear" ];then
  echo 'Delete debug file'
  rm -rf __debug*
fi
if [ "$1" == "start" ];then
  echo 'Start a new Debug'
  dlv debug ./cmd/main.go --headless --listen=:2345 --api-version=2 --accept-multiclient &
fi
T=${!}
wait ${T}
