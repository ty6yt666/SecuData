# SecuData


[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)
[![Docker](https://img.shields.io/docker/pulls/gin-gonic/gin)](https://hub.docker.com/r/gin-gonic/gin)

SecuData is a financial data project written in Go language, utilizing Gin as the web framework. The project provides RPC and RESTful APIs for user use, without including data crawling functionality. The project's features are continuously being improved.

The key features of SecuData are:
- Provide basic financial data.
- Provide financial time series data.
- Other functions


## Getting started
- RESTful API - apiserver
  - Native deployment
      ```shell
      # pull dependency
      go mod tidy
      go mod download
      ```
    
      ```shell
      # build & start
      go build -o ./cmd/apiserver/apiserver .
      ./cmd/apiserver/apiserver
      ```
  - Containerized deployment
    ```shell
    # build docker image
    docker build -t repository_ip:port/SecuData . 
    ```

    ```shell
    # start - 1
    docker run -d --name SecuData -p 8080:8080 repository_ip:port/SecuData
    ```
    ```shell
    # start - 2
    docker-compose up -d
    ```
    ```shell
    # start - 3 k8s deploy 
    under development...
    ```


- RPC API - rpcserver
  ```
  under development...
  ```

