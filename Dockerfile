FROM golang:1.21.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

# Copy all files in host dir to container
COPY . .
# Ensure all required packeges properly installed
RUN go mod tidy