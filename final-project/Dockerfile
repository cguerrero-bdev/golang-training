# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest
ENV GO111MODULE=on

# Copy the local package files to the container's workspace.
WORKDIR /go/src/qaa-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd api && go install .

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/api