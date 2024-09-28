FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

COPY ./cron ./cron
COPY .env ./cron

RUN go build -o hello ./cron/scripts/update_age.go

FROM alpine

WORKDIR /build

COPY .env /build
COPY --from=builder /build/hello /build/hello
RUN chmod +x /build/hello

COPY ./cron/crontab /etc/crontab
RUN crontab /etc/crontab
RUN touch /var/log/cron.log
CMD crond -f
