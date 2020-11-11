FROM golang:latest


WORKDIR .

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd

WORKDIR /dist

RUN cp /build/main .

EXPOSE 3000

CMD ["/dist/main"]
