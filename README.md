# SINALOA-API

REST APIs used to consume the sinaloa-cli commands.


## Go Module Init

```bash
go mod init github.com/dalecosta1/sinaloa-api
go mod tidy
```


# Go modules & libraries

## Install Swagger Library

```bash
go install github.com/swaggo/swag/cmd/swag
```

## Install Gin-gogin

```bash
go get -u github.com/gin-gonic/gin
```

## Install Godotenv

```bash
go install github.com/joho/godotenv/cmd/godotenv@latest
```

## Get references

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger
go get -u github.com/dalecosta1/sinaloa-api/api
go get -u github.com/dalecosta1/sinaloa-api/controller
go get -u github.com/dalecosta1/sinaloa-api/docs
go get -u github.com/dalecosta1/sinaloa-api/middlewares
go get -u github.com/dalecosta1/sinaloa-api/repository
go get -u github.com/dalecosta1/sinaloa-api/service
go get -u github.com/joho/godotenv
```


## Run

```bash
go run main.go
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
