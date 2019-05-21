# First we set our base image as latest Golang (no tag) and we name our stage "as builder."
FROM golang:1.12.1-alpine3.9 AS builder

# Set the working directory
WORKDIR $GOPATH/src/github.com/martinrocket/learning/1.1
# Copy the project files 
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build .

# Stage 2 build the inage we deploy
# FROM scratch
FROM alpine 
RUN apk add --no-cache ca-certificates
COPY --from=builder $GOPATH/go/src/github.com/martinrocket/learning/1.1 . 
ENTRYPOINT ["./1.1"]
EXPOSE 8080
