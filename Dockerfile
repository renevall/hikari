FROM golang:1.8.1-alpine
ADD . /go/src/bitbucket.org/reneval/hikari
RUN apk add --no-cache git gcc musl-dev \
&& go get github.com/blevesearch/bleve \
&& go get github.com/blevesearch/blevex/lang/es \
&& go get gopkg.in/gin-gonic/gin.v1 \
&& go install bitbucket.org/reneval/hikari \
&& apk del git gcc musl-dev
ENTRYPOINT /go/bin/hikari
EXPOSE 8585

