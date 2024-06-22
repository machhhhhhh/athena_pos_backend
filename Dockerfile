# Stage 1: Build
FROM golang:1.20-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files first and download the dependencies.
# This is done first to leverage Docker cache layers, as these files 
# change less frequently than your source code.
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project 
# This includes your Go files, the templates directory, and any other necessary files.
COPY . .

# Build the application
RUN go build -o main .

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the executable
CMD ["./main"]