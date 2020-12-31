FROM golang:1.13.7

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

USER $USERNAME
WORKDIR /home/$USERNAME

RUN sudo apt-get update && sudo apt-get install -y curl git procps && \
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)" && \
    eval $(/home/linuxbrew/.linuxbrew/bin/brew shellenv) && \
    echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.profile && \
    brew tap aws/tap && \
    brew install aws-sam-cli docker && \
    brew install aws-sam-cli && \
    echo "PATH=${PATH}:/usr/local/go/bin" >> ~/.bashrc
