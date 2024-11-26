# Use an official Go runtime as a parent image
FROM golang:1.23.3

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum, and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Expose the application on port 8080
EXPOSE 8080

# Build the application
RUN go build -o main ./cmd

# Command to run the application
CMD ["./main"]
