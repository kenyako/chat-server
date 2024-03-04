FROM golang:1.20.3-alpine AS builder

COPY . /github.com/kenyako/chat-server/source/
WORKDIR /github.com/kenyako/chat-server/source/

RUN go mod download
RUN go build -o ./bin/server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kenyako/chat-server/source/bin/server .

CMD [ "./server" ]