# simple-proxy

A super simple proxy server

## Run

```bash
docker run -it -e PROXY_URL=https://mockbin.org/echo -p 9900:9900 gbbirkisson/simple-proxy:latest
```

or

```bash
PROXY_URL=https://mockbin.org/echo go run main.go
```