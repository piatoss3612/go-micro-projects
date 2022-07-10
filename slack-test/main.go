package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(events <-chan *slacker.CommandEvent) {
	for event := range events {
		fmt.Println("Command Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// initialize slack bot client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// start goroutine that read command events and print each event information
	go printCommandEvents(bot.CommandEvents())

	// register command "ping" that slack bot replies "pong" in response
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start slack bot
	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
