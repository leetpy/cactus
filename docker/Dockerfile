FROM golang:latest

WORKDIR /go/src/github.com/leetpy/cactus/
ADD ./ /go/src/github.com/leetpy/cactus/

RUN export GO111MODULE=on \
    && export GOPROXY=https://goproxy.io \
    && GOOS=linux go mod tidy \
    && go build -o app .

CMD ["./app", "-c", "conf/config.yaml"]

