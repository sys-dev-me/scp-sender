FROM golang:1.13

LABEL maintainer="Kravchenko Dmytro E. <kravchenko.d@newton.life>"
WORKDIR $GOPATH/src/scp-sender


COPY *.go ./
COPY .settings .
COPY web web

RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 3001
EXPOSE 3002

RUN go build

CMD ["scp-sender"]


