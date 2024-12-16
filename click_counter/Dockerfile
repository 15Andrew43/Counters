FROM golang:1.22-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY ./click_counter/go.mod ./click_counter/go.sum ./

RUN go mod download

COPY ./click_counter/ .

RUN go build -o click_counter ./cmd/main.go

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/click_counter .

CMD ["./click_counter"]