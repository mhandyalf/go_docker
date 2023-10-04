FROM golang:alpine

# Install Git and update package lists
RUN apk update && apk add --no-cache git

# Set the working directory in the container
WORKDIR /app

# Copy the application code into the container
COPY . .

# Download and update dependencies
RUN go mod tidy
RUN go mod download

# Build the Go application
RUN go build -o binary

# Set the entry point for the container
ENTRYPOINT ["/app/binary"]
