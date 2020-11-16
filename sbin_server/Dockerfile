FROM golang:alpine

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .
RUN mkdir bin
RUN go build -o bin/server server.go

RUN ["chmod", "+x", "bin/server"]

RUN echo $(ls bin)
RUN echo $(pwd)

EXPOSE 50051

ENTRYPOINT [ "bin/server" ]