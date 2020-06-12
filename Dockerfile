FROM golang:alpine

ENV GO11MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p gosearch

COPY . /gosearch
WORKDIR /gosearch
RUN go mod download
RUN go build -o gosearchserver /gosearch/cmd/gRPC/server/

EXPOSE 8888

CMD ["/gosearch/gosearchserver"]