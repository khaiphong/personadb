# start an official go which installs golang and sets GOPATH
FROM golang:1.10.2 AS build-env
LABEL maintainer="duong.batien@khaiphong.io"

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/khaiphong/personadb
# set working directory in container. All subsequent commands will run from this directory
# which mimic src directory to manually do develeopment and use github for production.
WORKDIR /go/src/github.com/khaiphong/personadb

RUN go get .
RUN go build

# package the image in alpine for distribution
# FROM alpine
# COPY --from=build-env /go/src/github.com/khaiphong/personadb \
#                      /go/src/github.com/khaiphong/personadb
# RUN chown nobody:nogroup /go/src/github.com/khaiphong/personadb
# USER nobody

CMD if [ ${APP_ENV} = production ]; \
	then \
	personadb; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

# Make port 8080 available to the world outside this container
EXPOSE 8080

