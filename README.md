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


## Install python components

On Ubuntu/Debian:
```bash
sudo apt-get install pyton3 -y
sudo apt-get install pkg-config -y
sudo apt-get install python3-dev -y
```

On CentOS/Fedora:
```bash
sudo yum install python
sudo yum install pkgconfig
sudo yum install python3-devel
```

On macOS:
```bash
brew install python
brew install pkg-config
```
Look for the path of python.pc (should be similar to that: /usr/local/opt/python@3.11/lib/pkgconfig/python3.pc)
```bash
find /usr/local/opt/python@3*/lib/pkgconfig -name 'python3.pc'
```
Now we can set the export taking the path found until before python3.pc
```bash
echo 'export PKG_CONFIG_PATH="/usr/local/opt/python@3/lib/pkgconfig"' >> ~/.zshrc
source ~/.zshrc
```
```bash
export PKG_CONFIG_PATH="$(dirname $(dirname $(which python3)))/lib/pkgconfig"
export CGO_CFLAGS="$(python3-config --cflags)"
export CGO_LDFLAGS="$(python3-config --ldflags)"
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
