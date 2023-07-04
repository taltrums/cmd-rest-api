FROM golang:1.20.4-alpine

RUN apk add --no-cache git

# set working directory
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# Copy the source code
COPY . . 

# download Go modules and dependencies
RUN go mod download

# Build the Go app
RUN go build -o api .

#EXPOSE the port
EXPOSE 8080

# Run the executable
CMD ["./api"]