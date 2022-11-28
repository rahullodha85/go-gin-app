FROM golang:alpine as builder

COPY . /source

WORKDIR /source

RUN go build -o /output/app

WORKDIR /output

FROM alpine:latest

WORKDIR /output

COPY --from=builder /output/app ./

EXPOSE ${PORT}

ENTRYPOINT ["./app"]
