FROM golang:1.21.0

WORKDIR /usr/src/app/server

# Unversioned install
RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy all files in host dir to container
COPY . .

RUN go mod tidy