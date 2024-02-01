FROM golang:1.21 as builder
ARG GOARCH=arm64
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$GOARCH go build -o main .

FROM alpine:latest  
RUN addgroup -S app && adduser -S -h /app app -G app
WORKDIR /app
COPY --from=builder /app/main .
USER app
EXPOSE 8080
ENTRYPOINT ["/app/main"]
