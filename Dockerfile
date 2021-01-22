FROM golang:latest

ENV GOPROXY https:://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/newinternetboy/poor_union

COPY . $GOPATH/src/github.com/newinternetboy/poor_union

RUN go build .

EXPOSE 8000

ENTRYPOINT [ "./poor_union" ]
