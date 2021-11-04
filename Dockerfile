FROM golang:1.16-alpine AS builder

WORKDIR /webhook

COPY main.go .
RUN go mod init webhook && go get

RUN CGO_ENABLED=0 go build -o bin/webhook

FROM alpine:3.14

EXPOSE 8080

WORKDIR /webhook
COPY --from=builder /webhook/bin/webhook .

USER 2001:2001
CMD ["./webhook"]
