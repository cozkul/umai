FROM golang:1.21.0

WORKDIR /usr/src/app/server

RUN go install github.com/cosmtrek/air@latest

# Copy all files in host dir to container
COPY ./server .

RUN go mod tidy