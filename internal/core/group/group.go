package group

import (
	"app/internal/info"
	"app/internal/rule"
	"fmt"
	"log/slog"

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
	text := fmt.Sprintf(`
добро пожаловать в %s, %s!
советую прочитать правила, прежде чем начать общение 
для этого есть команда /rules 
удачного времяпровождения ^^
	`,
		update.Chat.FirstName, update.From.FirstName)

	params := tu.Message(
		tu.ID(update.Chat.ID),
		text,
	)
	wasMember := update.OldChatMember.MemberIsMember()
	isMember := update.NewChatMember.MemberIsMember()

	if wasMember || !isMember {
		return nil
	}

	user := update.NewChatMember.MemberUser()

	// logs
	slog.Info("user joined chat",
		"chat_id", update.Chat.ID,
		"user_id", user.ID,
		"username", "@"+user.Username,
	)

	if rule.IsFozz(update.Chat.ID) {
		text = "+++"
		params.MessageThreadID = info.Fozz.TMainID
		params.Text = text
	} else if rule.IsFoss(update.Chat.ID) {
		text = fmt.Sprintf(`
Добро пожаловать в meNFoss, %s ^^
это чат для общения, привязанный к meNFlux
почитай правила и можешь начинать общатся
		`, update.NewChatMember.MemberUser().FirstName)
		params.Text = text
	} else {
		text = "---"
		params.Text = text
	}

	_, err := ctx.Bot().SendMessage(ctx, params)

	return err
}
