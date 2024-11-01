FROM golang:1.22

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local files to the container
COPY . .

# Build the Go application
RUN go build -o simple-project

# Command to run the application
CMD ["./simple-project"]

