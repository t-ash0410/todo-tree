# Setup

## install sam
https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html

## install docker
https://docs.docker.com/get-docker/

## install vscode
https://azure.microsoft.com/ja-jp/products/visual-studio-code/

## install Remote-Containers
拡張機能から「Remote-Containers」で検索して、インストール

# start container
shift + command + P => >Remote-Containers: Rebuild And Reopen Container

# build
コンテナのワークディレクトリで実行

```
$ sam build
```

# debug
ホストのワークディレクトリで実行

```
$ sam local start-api -p 8080 --env-vars config.json
```
