# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o server ./cmd/server

# Runtime stage
FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=builder /app/server /server

EXPOSE 8443
USER nonroot

ENTRYPOINT ["/server"]
