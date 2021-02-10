#!/bin/bash

pwd=$(pwd)
command='docker run -it'
command+=' -v /var/run/docker.sock:/var/run/docker.sock'
command+=' -v '$pwd':/tmp/todo-tree-backend'
command+=' todo-tree_backend-editor /bin/bash '\''/tmp/todo-tree-backend/scripts/build.sh'\'
echo "command: $command"
eval $command

host=$(hostname)
command='sam local start-api'
command+=' -p 8080'
command+=' --parameter-overrides '\''DBEndpoint='$host' DBPort=3306 DBRootUserName=root DBRootUserPwd=password AllowOrigin=http://localhost:3000'\'
echo "command: $command"
eval $command