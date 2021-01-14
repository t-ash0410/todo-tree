# build
コンテナのワークディレクトリで実行

```
$ yarn dev
```

# debug
コンテナのワークディレクトリで実行

```
$ yarn dev
```

# debug AWS Code Build
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
docker run \
  -it -v /var/run/docker.sock:/var/run/docker.sock \
  -e "IMAGE_NAME=aws-codebuild-4.0" \
  -e "ARTIFACTS={出力先パス}" \
  -e "SOURCE=todo-tree" \
  -e "BUILDSPEC=todo-tree/frontend/buildspec.yml" \
  amazon/aws-codebuild-local
```