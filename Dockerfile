FROM golang:1.10-alpine

WORKDIR $GOPATH/src/go.thomasd.se/ebooks/
RUN mkdir -p $GOPATH/src/go.thomasd.se/ebooks/

COPY slack/ slack/
COPY *.go ./

RUN go install


FROM alpine

RUN apk --no-cache --update add ca-certificates

EXPOSE 80

RUN mkdir -p /etc/ebooks/
COPY config.json /etc/ebooks/

COPY --from=0 /go/bin/ebooks /usr/local/bin/ebooks
ENTRYPOINT ["/usr/local/bin/ebooks"]
