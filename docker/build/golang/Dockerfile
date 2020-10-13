# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.15-alpine

# The latest alpine image don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/goproject
COPY . .

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download -x
RUN go get -u github.com/cosmtrek/air

#Command to build the Go app with hot reload
ENTRYPOINT ["/go/bin/air"]

