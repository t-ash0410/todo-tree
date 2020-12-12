FROM golang:1.13.7

WORKDIR /todo-tree-backend

ENV GO111MODULE=on
ENV GOPATH /go

RUN apt-get update \
    && apt-get install -y git \
    && go get -v \
    github.com/mdempsky/gocode \
    github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    github.com/acroca/go-symbols \
    golang.org/x/tools/cmd/guru \
    golang.org/x/tools/cmd/gorename \
    github.com/cweill/gotests/... \
    github.com/fatih/gomodifytags \
    github.com/josharian/impl \
    github.com/davidrjenni/reftools/cmd/fillstruct \
    github.com/haya14busa/goplay/cmd/goplay \
    github.com/godoctor/godoctor \
    github.com/go-delve/delve/cmd/dlv \
    github.com/stamblerre/gocode \
    golang.org/x/lint/golint \
    github.com/ramya-rao-a/go-outline \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*