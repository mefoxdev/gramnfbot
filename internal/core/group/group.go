package group

import (
	"fmt"
	"html"
	"log/slog"

	"github.com/mefoxtrot/meNFlubot/internal/info"
	"github.com/mefoxtrot/meNFlubot/internal/rule"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleRules(ctx *th.Context, message telego.Message) error {
	text := ""
	if rule.IsFozz(message.Chat.ID) {
		text = info.FozzRules
	} else if rule.IsFoss(message.Chat.ID) {
		text = info.FossRules
	} else {
		text = info.BaseRules
	}

	params := tu.Message(
		tu.ID(message.Chat.ID),
		text,
	)
	_, err := ctx.Bot().SendMessage(ctx, params)
	return err
}

func HandleJoin(ctx *th.Context, update telego.ChatMemberUpdated) error {
	wasMember := update.OldChatMember.MemberIsMember()
	isMember := update.NewChatMember.MemberIsMember()

	user := update.NewChatMember.MemberUser()
	chatTitle := html.EscapeString(update.Chat.Title)

	mention := fmt.Sprintf(
		`<a href="tg://user?id=%d">%s</a>`,
		user.ID,
		html.EscapeString(user.FirstName),
	)
	text := fmt.Sprintf(`
Добро пожаловать в %s, %s!
прочитай правила, и можешь начать общаться ^^
для этого введи /rules
	`,
		chatTitle, mention)

	params := tu.Message(
		tu.ID(update.Chat.ID),
		text,
	).WithParseMode(telego.ModeHTML)

	if wasMember || !isMember {
		return nil
	}

	// logs
	slog.Info("user joined chat",
		"chat_id", update.Chat.ID,
		"user_id", user.ID,
		"username", "@"+user.Username,
	)

	_, err := ctx.Bot().SendMessage(ctx, params)

	return err
}
