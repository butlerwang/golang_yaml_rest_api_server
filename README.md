# YAML RESTful API Example with golang

## Install and Run

```shell
$ go get github.com/butlerwang/golang_yaml_rest_api_server

$ cd $GOPATH/src/github.com/butlerwang/golang_yaml_rest_api_server
$ go build
$ ./go-restful-api-example
```

## API Endpoint
- http://localhost:31415/api/v1/app
    - `POST`: create an app
- http://localhost:31415/api/v1/app/{name}
    - `GET`: get an app
    - `PUT`: update an app
    - `DELETE`: remove app
- http://localhost:31415/api/v1/apps
    - `GET`: get all of apps 
- http://localhost:31415/api/v1/app/search/{name}
    - `GET`: search an app

## Data Structure

```yaml
title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
  Interesting Title
```

## Simple Token Authentification

Need special http request header: "X-token: 123456"

## Todo

1. Improve test
2. Improve logger
3. Sophisticated Authentification
4. Implement config file
