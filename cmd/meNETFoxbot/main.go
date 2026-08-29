package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mefoxtrot/gramnfbot/internal/router"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	_ = godotenv.Load()

	bot, err := telego.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	updates, err := bot.UpdatesViaLongPolling(
		context.Background(),
		&telego.GetUpdatesParams{
			Timeout: 8,
			AllowedUpdates: []string{
				telego.MessageUpdates,
				telego.CallbackQueryUpdates,
				telego.ChatMemberUpdates,
				telego.ChatJoinRequestUpdates,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	handler, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Stop()

	router.Register(handler)

	log.Println("bot started")

	if err := handler.Start(); err != nil {
		log.Fatal(err)
	}
}
