FROM golang as builder

ENV GO111MODULE=on

EXPOSE 7777

WORKDIR /go/src/test_task/

COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main/main.go

ENTRYPOINT ["./main"]