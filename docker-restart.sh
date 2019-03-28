#!/bin/bash

#
# Pipe docker 更新重启脚本
#
# 1. 请注意修改参数
# 2. 可将该脚本加入 crontab，每日凌晨运行来实现自动更新
#

docker pull b3log/pipe
docker stop pipe
docker rm pipe
docker run --detach --name pipe --network=host \
      b3log/pipe --mysql="root:123456@(127.0.0.1:3306)/pipe?charset=utf8mb4&parseTime=True&loc=Local" --runtime_mode=prod --server=http://localhost:5897
