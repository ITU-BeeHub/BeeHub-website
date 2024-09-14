# Use an official Golang image as the base image
FROM golang:1.23-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/website

# Expose the port that the app will run on
EXPOSE 8080

# Set environment variables
ENV SWAGGER_ENABLED=true
ENV GIN_MODE=release

# Command to run the application
CMD ["/app/main"]
