FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

COPY ./app ./app
COPY ./docs ./docs
COPY .env ./app

RUN go build -o main ./app/cmd/main.go

FROM alpine

WORKDIR /build

COPY .env /build
RUN mkdir /build/docs
COPY --from=builder /build/docs /build/docs
COPY --from=builder /build/main /build/main
RUN chmod +x /build/main
CMD ["/build/main"]
