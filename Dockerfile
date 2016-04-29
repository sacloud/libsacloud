FROM golang:alpine
MAINTAINER "Kazumichi Yamamoto <yamamoto.febc@gmail.com>"

RUN apk add --update ca-certificates git && \
    go get github.com/stretchr/testify/assert

ENV SRC=$GOPATH/src/github.com/yamamoto-febc/libsacloud/
ADD . $SRC
WORKDIR $SRC

ENTRYPOINT [ "go" ]
CMD [ "test", "-v", "./..." ]
