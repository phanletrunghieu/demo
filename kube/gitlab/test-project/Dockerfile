FROM golang:1.12-alpine as builder
RUN apk update && apk add --no-cache git make ca-certificates tzdata && update-ca-certificates
WORKDIR /test-project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod verify
RUN go build -o app main.go

FROM alpine:3.10
RUN adduser -D -g '' appuser
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /test-project/app /app

USER appuser

CMD ["/app"]