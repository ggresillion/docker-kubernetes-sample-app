FROM golang:alpine as builder

WORKDIR /app

COPY go.* /app/

RUN go mod download

COPY . .

RUN go build -ldflags "-w" -o sample-app .

FROM alpine

WORKDIR /app

COPY --from=builder /app/sample-app .

EXPOSE 8080

CMD ["/app/sample-app"]
