package direct

import (
	"app/internal/rule"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleStart(ctx *th.Context, message telego.Message) error {
	if rule.IsDirect(message.Chat) {
		_, err := ctx.Bot().SendMessage(
			ctx,
			tu.Message(
				tu.ID(message.Chat.ID),
				"Привет! это бот созданный netfox.me.\nон нужен для модерации в чатах meNFozz(ss), авторизации в наших продуктах и оплаты подписки на наш MTProto-прокси",
			),
		)
		return err
	}

	return nil
}
