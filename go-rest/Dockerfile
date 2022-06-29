FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY src/go.mod .
COPY src/go.sum .
COPY . .
RUN go mod download
RUN go build -o ./go-rest ./src/main.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/go-rest .
EXPOSE 8080
ENTRYPOINT ["./go-rest"]