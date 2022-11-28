# go-gin-app

### Run Local
```HOST="0.0.0.0" PORT=8080 go run main.go```

### Build Docker Image
```docker build -t <image-tag> .```

### Run Container

```docker run -d -p <port>:<port> --name <container-name> --env HOST=0.0.0.0 --env PORT=<port> <image-tag>```