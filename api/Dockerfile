FROM golang:1.15 AS builder

WORKDIR /app
COPY . .

RUN go build -tags netgo -a -o api main.go

FROM alpine:3.12

ENV HOST=0.0.0.0
ENV PORT=80

WORKDIR /root
COPY --from=builder /app/api .

CMD ["./api"]
