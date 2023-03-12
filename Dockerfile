ARG DISTRO
FROM golang:${DISTRO} as builder

ENV GO111MODULE on

COPY . /build
WORKDIR /build
RUN go generate ./lvm
RUN go build -o /docker-lvm-plugin

FROM debian:${DISTRO}-slim

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get -y install lvm2 xfsprogs thin-provisioning-tools && \
    apt-get clean
COPY --from=builder /docker-lvm-plugin /

CMD ["/docker-lvm-plugin", "-debug"]

