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

pwd=$(pwd)
command='docker run -it'
command+=' -v /var/run/docker.sock:/var/run/docker.sock'
command+=' -v '$pwd':/tmp/todo-tree-backend'
command+=' todo-tree_backend-editor /bin/bash '\''/tmp/todo-tree-backend/build.sh'\'
echo "command: $command"
eval $command

vpc=$(aws ec2 describe-vpcs --filters Name=tag:Name,Values=todo-tree-vpc --query "Vpcs[*].[VpcId]" --output text)
echo "vpc id: ${vpc}"

sg=$(aws ec2 describe-security-groups --filters Name=vpc-id,Values=$vpc Name=group-name,Values=todo-tree-lambda-security-group --query "SecurityGroups[*].[GroupId]" --output text)
echo "security group id: $sg"

subnet1=$(aws ec2 describe-subnets --filters Name=vpc-id,Values=$vpc Name=tag:Name,Values=todo-tree-private-subnet-1 --query "Subnets[*].[SubnetId]" --output text)
echo "subnet1 id: $subnet1"

subnet2=$(aws ec2 describe-subnets --filters Name=vpc-id,Values=$vpc Name=tag:Name,Values=todo-tree-private-subnet-2 --query "Subnets[*].[SubnetId]" --output text)
echo "subnet2 id: $subnet2"

db_endpoint=$(aws rds describe-db-proxies --filters Name=DBProxyName,Values=todo-tree-rds-proxy --query "DBProxies[*].[Endpoint]" --output text)
echo "db endpoint: $db_endpoint"

arns=$(aws cloudfront list-distributions | jq -r ".DistributionList.Items[].ARN")
for a in $arns
do
    if [ $(aws cloudfront list-tags-for-resource --resource $a | jq -r ".Tags.Items[].Value") == 'todo-tree-front-distribution' ]; then
        distribution_id=${a##*/}
    fi
done
alloworigin='https://'
alloworigin+=$(aws cloudfront get-distribution --id $distribution_id | jq -r ".Distribution.DomainName")
echo "allow origin endpoint: $alloworigin"

command="sam deploy --parameter-overrides 'FunctionSecurityGroupId=$sg PrivateSubnet1Id=$subnet1 PrivateSubnet2Id=$subnet2 DBEndpoint=$db_endpoint DBRootUserPwd=$password AllowOrigin=$alloworigin'"
echo "deploy command: $command"
eval $command

echo "complete."