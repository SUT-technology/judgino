# Use a Go base image to run the code
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the Go code into the container
COPY . .

# Install necessary dependencies (if needed, for example)
RUN go mod tidy

# Expose the port
EXPOSE 8081

# Command to run the Go code
CMD ["go", "run", "main.go"]