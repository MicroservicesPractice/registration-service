FROM golang:1.21.3-alpine AS builder
LABEL stage=gobuilder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./app/cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/db ./db
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

EXPOSE 5002

CMD ["./main"]