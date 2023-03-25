# Golang Study Projects

- mini interesting projects taking 1~2 hours

## grpc-go

- gRPC CRUD API with MongoDB in Go
- [gRPC [Golang] Master Class: Build Modern API & Microservices](https://www.udemy.com/course/grpc-golang/)

## go-google-scraper

- scrape google search result with goquery package
- [Golang Project 2021 - Scrape Google Results](https://www.youtube.com/watch?v=1YPPzaApyJE&t=3s)

## todo-app

- todo app on cli

```cmd
$ ./todo -list
```

```cmd
$ ./todo -add <todo>
```

```cmd
$ ./todo -delete <index>
```

```cmd
$  ./todo -complete <index>
```

- [Golang Tutorial: Build A Beautiful CLI Todo App With Support for Piping](https://youtu.be/j1CXoOQXbco)

## url-shortener

- url shortener (Redis + Golang fiber framework)
- containerize using docker compose

```cmd
$ cd ./url-shortener
$ docker-compose up -d
```

- test using [postman](https://www.postman.com/)
- [URL Shortener (Redis + Go-Fiber) - Golang Project [ Intermediate Level ]](https://youtu.be/edCnzelVRlc)

## cli-reminder

- seding desktop alert using cli

```cmd
$ ./reminder 22:53 hey how are you
```

- [GO Project - CLI Reminder Tool [ Intermediate Project ]](https://youtu.be/HnNT6MnRlFM)

## email-verifier

- golang has powerful standard library to verify domain
- [GOlang - Email Verifier Tool Project](https://youtu.be/9E4UEsWpYvM?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9)

## slack-age-bot

- slack bot calcualting user age
- hide environment variables with godotenv package

```slack
command: my yob is <year>
```

- [GOLANG SLACKBOT To Calculate Age](https://youtu.be/HnPm69i60xE?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9)

## golang-chat

- Golang + React fullstack project
- create chat app with websocket

### backend
```cmd
$ cd ./golang-chat/backend
$ go run .
```
### frontend
```cmd
$ cd ./golang-chat/frontend
$ npm start
```

- [GO + REACT Fullstack App - RealTime Chat](https://youtu.be/xdzLr246fXI)


## slack-test

- ping pong slack bot

- [Slack Bot With GO 🤖 - Golang Project Ideas 2021🤘🏼🤸🏼‍♂️](https://youtu.be/DhM3g2DvmT8?list=PL5dTjWUk_cPYj8C3QhFMxhMOj7bU1uv6v)

## calorie-tracker

### backend

```cmd
$ cd ./calorie-tracker/backend
$ go run .
```

### frontend

```cmd
$ cd ./calorie-tracker/frontend
$ npm start
```

### MongoDB

Run MongoDB on Docker Desktop

If you try to run calorie tracker with your own MongoDB,

you should change `mongoURL` and Credentials with your own

```go
package routes

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBInstance()

func DBInstance() *mongo.Client {
	mongoURL := "mongodb://localhost:49153" // change here

	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "docker", // change here
		Password: "mongopw", // change here
	})

	...
}
```

## nft-collection-generator

```cmd
$ cd ./nft-collection-generator
$ go run .
```

- [packagemain #24: Generate an NFT Collection in Go](https://youtu.be/QPvE6qxdTDk)


## Golang-ethereum

- [youtube playlist](https://youtube.com/playlist?list=PLay9kDOVd_x7hbhssw4pTKZHzzc6OG0e_)

## GO-GPT-CLI

> Need to get api key from OpenAI

```cmd
$ cd ./go-gpt
$ go run .
```

- [Chat GPT GOlang Project BUILD](https://youtu.be/QNIQXpdpBuA)

## GO GRPC DEMO

```cmd
$ go run ./server
```
```cmd
$ go run ./client
```

- [FULL PROJECT - GO + GRPC](https://youtu.be/a6G5-LUlFO4)