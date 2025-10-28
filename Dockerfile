FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

RUN go generate ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api cmd/api/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]
