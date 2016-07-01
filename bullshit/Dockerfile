FROM golang:1.5.3-onbuild


RUN go get -u -v github.com/derekparker/delve/cmd/dlv
EXPOSE 2345

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8000
CMD ["dlv", "debug", "--headless", "--listen=:2345", "--log"]