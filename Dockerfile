FROM golang:alpine as builder

COPY . /source

WORKDIR /source

RUN go build -o /output/app

WORKDIR /output

ENTRYPOINT ["./app"]

FROM alpine:latest

WORKDIR /output

COPY --from=builder /output/app ./

ENV HOST=0.0.0.0
ENV PORT=8080

EXPOSE ${PORT}

ENTRYPOINT ["./app"]
