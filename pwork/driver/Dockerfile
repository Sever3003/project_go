FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-s -w" -o /app/driver cmd/main.go

FROM ubuntu:latest
ENV IP=:8085
EXPOSE 8085

COPY --from=builder /app/driver ./driver

CMD ["./driver"]