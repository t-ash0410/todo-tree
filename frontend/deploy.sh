#!/bin/bash

echo "deploy start date: `date`"

# get api endpoint
api_id=$(aws apigateway get-rest-apis | jq -r '.items[] | select(.name == "todo-tree-backend") | .id')
api_endpoint="https://$api_id.execute-api.ap-northeast-1.amazonaws.com/Test"

# create docker env file
env_var="NODE_ENV=production
NEXT_PUBLIC_API_ENDOPOINT=$api_endpoint"
tmp_env_file_name="deploy.env"
command='echo -e "$env_var" > '$tmp_env_file_name
echo "command: $command"
echo "env_var: $env_var"
eval $command

# clear out directory
current_dir=$(pwd)
parent_dir=$(dirname $current_dir)
command='rm -r -f out'
echo "command: $command"
eval $command

# execute code build container
command='docker run -it'
command+=' -v /var/run/docker.sock:/var/run/docker.sock'
command+=' -v '$current_dir':/LocalBuild/envFile/ -e "IMAGE_NAME=aws-codebuild-4.0"'
command+=' -e "ARTIFACTS='$current_dir'/out"'
command+=' -e "SOURCE='$parent_dir'"'
command+=' -e "BUILDSPEC='$current_dir'/buildspec.yml"'
command+=' -e "ENV_VAR_FILE='$tmp_env_file_name'"'
command+=' amazon/aws-codebuild-local'
echo "command: $command"
eval $command

# delete docker env file
command="rm $tmp_env_file_name"
echo "command: $command"
eval $command

# unzip artifacts
command="unzip out/artifacts.zip"
echo "command: $command"
eval $command

# delete artifacts
command="rm out/artifacts.zip"
echo "command: $command"
eval $command

# s3 upload
deploy_backet="todo-tree-frontend"
command="aws s3 sync ./out/ s3://$deploy_backet --acl public-read --cache-control no-cache"
echo "command: $command"
eval $command

# cloudfront cache clear
arns=$(aws cloudfront list-distributions | jq -r ".DistributionList.Items[].ARN")
for a in $arns
do
    if [ $(aws cloudfront list-tags-for-resource --resource $a | jq -r ".Tags.Items[].Value") == 'todo-tree-front-distribution' ]; then
        distribution_id=${a:(-13)}
    fi
done
command=$(echo aws cloudfront create-invalidation --distribution-id $distribution_id --paths "/*")
echo "command: $command"
eval $command

echo "complete."