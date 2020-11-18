FROM golang:latest

WORKDIR /qaa-api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o qaa-api ./api

EXPOSE 3000

CMD ["./qaa-api"]