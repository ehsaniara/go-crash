FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/ehsaniara/go-crash
COPY . $GOPATH/src/github.com/ehsaniara/go-crash
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-crash"]
