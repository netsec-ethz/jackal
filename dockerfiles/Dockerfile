# start from base
FROM ubuntu:16.04

# install SCIONLab (you have to mount scionlab folder to /root/scionlab on build)
EXPOSE 50000/udp 1194/udp
RUN apt-get update && apt-get install -y --no-install-recommends \
    wget \
    make \
    sudo \
    pkg-config zip g++ zlib1g-dev unzip python git ca-certificates vim \
    openvpn
ADD bazel_install.sh /root/bazel_install.sh
# ADD scion_install_script.sh /root/scion_install_script.sh
ENV PATH=$PATH:/root/go/bin:/usr/local/go/bin:/root/.local/bin
ENV GOPATH=/root/go
RUN mkdir -p $GOPATH
WORKDIR /root
