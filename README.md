# Setup

## install docker
https://docs.docker.com/get-docker/

## install vscode
https://azure.microsoft.com/ja-jp/products/visual-studio-code/

## install Remote-Containers
拡張機能から「Remote-Containers」で検索して、インストール

# start container
shift + command + P => >Remote-Containers: Rebuild And Reopen Container

## 環境のデプロイ
プロジェクトのルートディレクトリで実行
```
$ bash scripts/deploy.env.sh -p ${password}
```

## 一括デプロイ
プロジェクトのルートディレクトリで実行
```
$ bash scripts/deploy.all.sh -p ${password}
```