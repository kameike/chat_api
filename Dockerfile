FROM golang
MAINTAINER kameike

WORKDIR /go/src/main/
RUN go get -u github.com/golang/dep/cmd/dep

ADD . .
RUN dep ensure
RUN go build -a -tags netgo -installsuffix netgo

FROM scratch
COPY --from=0 /go/src/main/main .
ENV PORT 1323
ENTRYPOINT ["/main"]

