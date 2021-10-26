# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Kaan Karaca <kaan94karaca@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
#WORKDIR /app
WORKDIR $GOPATH/src/github.com/h4yfans/case-study

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 
RUN go mod verify

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/case-study

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

#WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /go/src/github.com/h4yfans/case-study/db/migrations /db/migrations
COPY --from=builder /go/bin/case-study /go/bin/case-study

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["/go/bin/case-study"]