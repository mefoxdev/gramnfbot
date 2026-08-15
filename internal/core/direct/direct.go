package direct

import (
	"github.com/mefoxtrot/meNFlubot/internal/rule"
	"github.com/mefoxtrot/meNFlubot/internal/show"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleStart(ctx *th.Context, message telego.Message) error {
	if rule.IsDirect(message.Chat.Type) {
		if err := show.Profile(ctx, ctx.Bot(), message.Chat.ID); err != nil {
			return err
		}

		_, err := ctx.Bot().SendMessage(
			ctx,
			tu.Message(
				tu.ID(message.Chat.ID),
				"Привет! это бот созданный netfox.me.\nон нужен для модерации в чатах meNFozz(ss), авторизации в наших продуктах и оплаты подписки",
			),
		)
		return err
	}

	return nil
}
