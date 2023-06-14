FROM golang:alpine as builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o sample-app .

FROM alpine

WORKDIR /app

COPY --from=builder /app/sample-app .

EXPOSE 8080

CMD /app/sample-app
