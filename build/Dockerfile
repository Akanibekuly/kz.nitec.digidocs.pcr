FROM golang:1.15.5 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main cmd/main.go

FROM alpine:3.13.5
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/build build/.
CMD ls -la
CMD ["./main"]
EXPOSE 8080
