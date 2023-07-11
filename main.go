package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load(".env")
	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	SLACK_APP_TOKEN := os.Getenv("SLACK_APP_TOKEN")
	os.Setenv("SLACK_BOT_TOKEN", SLACK_BOT_TOKEN)
	os.Setenv("SLACK_APP_TOKEN", SLACK_APP_TOKEN)

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(ctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong!")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		panic(err.Error())
	}
}
