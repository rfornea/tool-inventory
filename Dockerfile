FROM golang:1.11
ENV ADDR=0.0.0.0
ENV GO111MODULE=on

RUN go version

RUN apt-get update
RUN apt-get install -y -q --no-install-recommends default-mysql-client
RUN apt-get install -y -q --no-install-recommends netcat
RUN apt autoremove -y
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/*

RUN mkdir -p "$GOPATH/src/github.com/rfornea/tool-inventory/pkg/backend"
WORKDIR "$GOPATH/src/github.com/rfornea/tool-inventory/pkg/backend"

COPY . .

RUN go build ./...
