FROM golang:1.12
LABEL maintainer="Kazumichi Yamamoto <yamamoto.febc@gmail.com>"

ENV SRC=$GOPATH/src/github.com/sacloud/libsacloud/
ADD . $SRC
WORKDIR $SRC

RUN make tools

ENTRYPOINT [ "make" ]
