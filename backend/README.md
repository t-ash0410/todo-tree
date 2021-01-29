# Setup

## install sam
https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html

# build
コンテナのワークディレクトリで実行

```
$ sam build
```

# debug
ホストのワークディレクトリで実行

```
$ sam local start-api \
    -p 8080 \
    --parameter-overrides 'RDSRootUserPwd=password AllowOrigin=localhost:3000'
```
