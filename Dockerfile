FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY . .

RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go mod download
RUN go build -o gowgboard .

FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/gowgboard .

CMD ["./gowgboard"]


















