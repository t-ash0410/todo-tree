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

echo "deploy environment"
eval "bash deploy.env.sh -p $password"

eval "cd backend"
echo "deploy backend"
eval "bash deploy.sh -p $password"

eval "cd ../frontend"
echo "deploy frontend"
eval "bash deploy.sh"

eval "cd ../"
echo "complete."