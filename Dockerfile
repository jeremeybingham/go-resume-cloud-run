FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api main.go
FROM scratch
WORKDIR /app
COPY --from=builder /app/go-api .
EXPOSE 8080
CMD ["./go-api"]