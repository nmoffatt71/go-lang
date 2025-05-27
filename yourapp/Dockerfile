# Start from official Go image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
RUN go mod tidy

# Copy the source code
COPY . .

# Generate GraphQL code
RUN go run github.com/99designs/gqlgen generate

# Build the Go app
RUN go build -o main .

# Expose the app port
EXPOSE 8080

# Start the app
CMD ["./main"]
