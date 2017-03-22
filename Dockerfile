FROM golang

ENV GOPATH=/go

WORKDIR /go/src/personal-site

COPY * /go/src/personal-site/

RUN go get
RUN go build -o /usr/bin/personal-site

EXPOSE 80 443

CMD ["/usr/bin/personal-site"]
