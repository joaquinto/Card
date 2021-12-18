# Start from golang base image
FROM --platform=linux/amd64 golang:1.16.5 AS build_base

# Set the current working directory inside the container 
WORKDIR /temp/build

# Copy go mod and sum files 
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go get -d ./...

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

# Start a new stage from scratch
FROM --platform=linux/amd64 alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the previous stage.
COPY --from=build_base /temp/build/main .

# Expose port to the outside world
EXPOSE $PORT

#Command to run the executable
CMD [ "./main" ]
