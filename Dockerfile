# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Update dependencies
RUN go get -u ./...

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]