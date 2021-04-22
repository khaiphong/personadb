# start an official go which installs golang and sets GOPATH
FROM golang:1.11 AS builder
# copy the ca-certificates.crt from our machine into our container
ADD ca-certificates.crt /etc/ssl/certs/

COPY . /go/src/github.com/khaiphong/personadb
WORKDIR /go/src/github.com/khaiphong/personadb

# get all dependencies and compile go program
RUN go get -d -v ./...
RUN go build -o db .

# the mount point for different containers in the same machine
VOLUME /tmp/personadb

# run personadb when the container launches
CMD ["/go/src/github.com/khaiphong/personadb/db"]

# Make port 8081 available to the world outside this container
EXPOSE 8081

# package in alpine for image built and serviced from personadb
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/khaiphong/personadb .


