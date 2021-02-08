# Setup

## install sam
https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html

# build
コンテナのワークディレクトリで実行

```
$ sam build -t template.yml
```

# debug
ホストのワークディレクトリで実行

```
$ sam local start-api \
    -p 8080 \
    --parameter-overrides 'DBEndpoint=${hostname} DBPort=3306 DBRootUserName=root DBRootUserPwd=password AllowOrigin=https://localhost:3000'
```

## AWS Code Build
## ビルド用Dockerfileのダウンロード
```
$ git clone https://github.com/aws/aws-codebuild-docker-images.git
```

## ビルド用docker imageのビルド
```
$ cd aws-codebuild-docker-images/ubuntu/standard/4.0
$ docker build -t aws-codebuild-4.0 .
```

## ビルド実行用docker imageのpull
```
$ docker pull amazon/aws-codebuild-local:latest --disable-content-trust=false
```

## run
```
$ docker run \
  -it -v /var/run/docker.sock:/var/run/docker.sock \
  -e "IMAGE_NAME=aws-codebuild-4.0" \
  -e "ARTIFACTS={出力先パス}" \
  -e "SOURCE=todo-tree" \
  -e "BUILDSPEC=todo-tree/frontend/buildspec.yml" \
  amazon/aws-codebuild-local
```

# deploy
ホストのワークディレクトリで実行

```
$ sam deploy --parameter-overrides 'FunctionSecurityGroupId=${FunctionSecurityGroupId} PrivateSubnet1Id=${PrivateSubnet1Id} PrivateSubnet2Id=${PrivateSubnet2Id} DBEndpoint=${DBEndpoint} DBRootUserPwd=${DBRootUserPwd} AllowOrigin=${AllowOrigin}'
```