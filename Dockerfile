FROM golang:1.18

WORKDIR /go/dating-bot

COPY go.mod go.sum ./
COPY ./chat/. ./chat

RUN go mod download && go mod verify

COPY . .
RUN go build main.go

CMD ["./main"]
