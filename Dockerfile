FROM golang:alpine as builder

COPY . /source

WORKDIR /source

RUN go build -o /output/app

FROM alpine:latest

WORKDIR /output

COPY --from=builder /output/app ./
COPY config config

ENTRYPOINT ["./app"]
