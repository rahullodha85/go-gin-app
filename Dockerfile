FROM golang:alpine

ENV HOST=0.0.0.0
ENV PORT=8080

COPY . /source

WORKDIR /source

RUN go build -o /output/app

WORKDIR /output

EXPOSE ${PORT}

ENTRYPOINT ["./app"]