FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-bin

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /go-bin /go-bin

EXPOSE 8080
# Run
CMD ["/go-bin"]