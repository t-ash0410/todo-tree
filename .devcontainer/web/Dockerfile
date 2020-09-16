FROM node:14.10.1 as node
FROM golang:1.13.7

COPY --from=node /opt/yarn-v1.22.5 /opt/yarn
COPY --from=node /usr/local/bin/node /usr/local/bin/
RUN ln -s /opt/yarn/bin/yarn /usr/local/bin/yarn \
  && ln -s /opt/yarn/bin/yarn /usr/local/bin/yarnpkg

WORKDIR /todo-tree

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
        # github.com/stamblerre/gocode \
        # github.com/rogpeppe/godef \
        # golang.org/x/tools/cmd/goimports \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*
