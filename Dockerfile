# start an official go which installs golang and sets GOPATH
FROM golang:1.10.2 AS build-env
LABEL maintainer="duong.batien@khaiphong.io"

ARG app_env
ENV APP_ENV $app_env

# set working directory in container so all subsequent commands will run from this directory
WORKDIR /go/src/github.com/khaiphong/personadb
COPY . /go/src/github.com/khaiphong/personadb

RUN go get .
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	personadb; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

# FROM alpine AS containerd-runtime
# COPY --from=build-env /go/src/github.com/khaiphong/personadb \
#                      /go/src/github.com/khaiphong/personadb

# Make port 8080 available to the world outside this container
EXPOSE 8080


