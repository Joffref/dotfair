FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o dotfair .

FROM golang:1.20
WORKDIR /app
COPY --from=builder /app/dotfair .
CMD ["./dotfair"]