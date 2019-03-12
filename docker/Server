FROM golang:alpine
MAINTAINER kameike

RUN apk add --update gcc musl-dev
RUN apk add --update git
RUN apk add --update sqlite
RUN apk add --update sqlite-dev



RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/kameike/chat_api

ADD Gopkg.lock .
ADD Gopkg.toml .
RUN dep ensure -v --vendor-only


ADD . .

RUN CGO_ENABLED=0 go build -a --tags "libsqlite3 linux netgo" -installsuffix netgo -o main cmd/server/main.go

# FROM scratch
FROM alpine
COPY --from=0 /go/src/github.com/kameike/chat_api/main .

ENV PORT 1323

ENTRYPOINT ["/main"]
