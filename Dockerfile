FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go mod download && go mod verify

COPY / .
RUN gso build -o app /usr/src/app/cmd/main.go

CMD ["./app"]