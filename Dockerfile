FROM golang:latest

WORKDIR /usr/src/app

COPY src/go.mod ./
COPY src/go.sum ./
COPY .env ./

RUN go mod download && go mod verify

COPY src/ .
RUN go build -o app /usr/src/app/cmd/main.go

CMD ["./app"]