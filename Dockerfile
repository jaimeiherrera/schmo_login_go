# Use the official Go Alpine image as the base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application will listen on
EXPOSE 8080

# Set the command to run the application when the container starts
CMD ["./main"]