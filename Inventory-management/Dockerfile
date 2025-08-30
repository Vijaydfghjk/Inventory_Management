# Use official Go image
FROM golang:1.23

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum to install dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project code
COPY . .

# Build your Gin app into a binary called 'main'
RUN go build -o main .

# Expose app port
EXPOSE 8080

# Start the app
CMD ["./main"]
