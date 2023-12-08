# SINALOA-API

## Go Module Init

```bash
go mod init gitlab.com/dalecosta1/sinaloa-api
go mod tidy
```

## Gin-Gonic library: github.com/gin-gonic/gin

## Run

```bash
go run server.go
```

# Swagger Documentation

## Install Swagger Library

```bash
go install github.com/swaggo/swag/cmd/swag
```

## Generate Swagger Documentation

Export env variables

```bash
export GOBIN=~/go/bin
export PATH=$PATH:$GOBIN
```

and

```bash
sudo chmod +x swagger.sh
./swagger.sh
```
