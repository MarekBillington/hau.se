FROM golang:1.19.5

WORKDIR /usr/src/app

COPY . .
COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app ./...

CMD ["./hause"]