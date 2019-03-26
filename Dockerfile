FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/gbbirkisson/simple-proxy
COPY . .

# Fetch dependencies.

# Using go get.
RUN go get -d -v

# Build the binary.
RUN go build -o /simple-proxy

############################
# STEP 2 build a small image
############################
FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Copy our static executable.
COPY --from=builder /simple-proxy /simple-proxy

# Run the hello binary.
ENTRYPOINT ["/simple-proxy"]