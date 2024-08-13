# Use the official Golang image as a build stage
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the application
RUN go build -o main .

# Use a smaller base image to run the application
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=builder /app/main .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]
