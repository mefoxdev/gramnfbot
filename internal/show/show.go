package show

import (
	"context"

	"github.com/mymmrac/telego"
)

const profileURL = "https://netfox.me/profile"

func Profile(ctx context.Context, bot *telego.Bot, chatID int64) error {
	return bot.SetChatMenuButton(ctx, &telego.SetChatMenuButtonParams{
		ChatID: chatID,
		MenuButton: &telego.MenuButtonWebApp{
			Type: telego.ButtonTypeWebApp,
			Text: "Profile",
			WebApp: telego.WebAppInfo{
				URL: profileURL,
			},
		},
	})
}
