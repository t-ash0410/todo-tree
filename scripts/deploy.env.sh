#!/bin/bash

while getopts p: OPT
do
    case $OPT in
    p) password=$OPTARG
        ;;
    esac
done

if [ "$password" == "" ]; then
    echo "password is required."
    exit -1
fi

echo "deploy start date: `date`"

my_arn=$(aws iam get-user | jq -r .User.Arn)
echo "my arn: $my_arn"

code_build_id=`aws iam get-role --role-name codebuild-todo-tree-service-role --query "Role.RoleId" --output text`
echo "code build id: $code_build_id"

command="sam deploy --parameter-overrides 'RDSRootUserPwd=$password CodeBuildRoleId=$code_build_id DeployUserArn=$my_arn'"
echo "deploy command: $command"
eval $command

echo "complete."