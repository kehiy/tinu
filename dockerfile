# Start with a base Golang Alpine image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files to the working directory
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go API binary
RUN go build -o main .

# Expose port 3000 for the API
EXPOSE 3000

# Run the Go API binary
CMD ["./main"]
