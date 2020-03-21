# To get packages tools
FROM golang:1.14.0-alpine AS builder

ARG ACCESS_TOKEN_USR="nothing"

ARG ACCESS_TOKEN_PWD="nothing"

# git is required to fetch go dependencies
RUN apk add --no-cache ca-certificates git

#Create a netrc file using the credentials specified using --build-arg
RUN printf "machine github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n\
    \n\
    machine api.github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n"\
    >> /root/.netrc

RUN chmod 600 /root/.netrc

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY /templates/ .
COPY go.mod .
# COPY go.sum .
# RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

RUN cp -r /build/templates/ .

RUN ls

# Build a small image
FROM alpine:3.7

RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /dist/ /

# Command to run
ENTRYPOINT ["/main"]
