# Start from the official Golang base image
FROM golang:1.21

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files (if present) to the container
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main


# Expose the port your app runs on
EXPOSE 8000

# Command to run the executable
CMD ["/main"]
