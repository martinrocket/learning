# First we set our base image as latest Golang (no tag) and we name our stage "as builder."
FROM golang:1.12.1-alpine3.9 AS builder

# Set the working directory
WORKDIR $HOME/src/github.com/martinrocket/learning/1.1
# Copy the project files 
COPY . . 
RUN go build .

# Stage 2 build the inage we deploy
FROM scratch
COPY --from=builder $HOME/src/github.com/martinrocket/learning/1.1 .
CMD ["clear"]
CMD ["./1.1"]
EXPOSE 8080
