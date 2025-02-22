# Base image: https://hub.docker.com/_/golang/
FROM golang:1.15

# Install golint
ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.34.1

# Add apt key for LLVM repository
RUN wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -

# Add LLVM apt repository
RUN echo "deb http://apt.llvm.org/stretch/ llvm-toolchain-stretch-5.0 main" | tee -a /etc/apt/sources.list

# Install clang from LLVM repository
RUN apt-get update && apt-get install -y --no-install-recommends \
    clang \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Set Clang as default CC
ENV set_clang /etc/profile.d/set-clang-cc.sh
RUN echo "export CC=clang-7.0" | tee -a ${set_clang} && chmod a+x ${set_clang}

# Install Docker (to build images)
ENV DOCKER_CHANNEL stable
ENV DOCKER_VERSION 19.03.5

RUN set -eux; \
    \
    apkArch="$(arch)"; \
    case "$apkArch" in \
    # amd64
    x86_64) dockerArch='x86_64' ;; \
    # arm32v6
    armhf) dockerArch='armel' ;; \
    # arm32v7
    armv7) dockerArch='armhf' ;; \
    # arm64v8
    aarch64) dockerArch='aarch64' ;; \
    *) echo >&2 "error: unsupported architecture ($apkArch)"; exit 1 ;;\
    esac; \
    \
    if ! wget -O docker.tgz "https://download.docker.com/linux/static/${DOCKER_CHANNEL}/${dockerArch}/docker-${DOCKER_VERSION}.tgz"; then \
    echo >&2 "error: failed to download 'docker-${DOCKER_VERSION}' from '${DOCKER_CHANNEL}' for '${dockerArch}'"; \
    exit 1; \
    fi; \
    \
    tar xzvf docker.tgz --strip=1 -C /usr/local/bin docker/docker; \
    rm docker.tgz; \
    \
    docker --version

# Install Docker Compose
RUN curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose