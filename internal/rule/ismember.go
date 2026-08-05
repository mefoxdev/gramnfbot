package rule

import (
	"context"
	"errors"

	"github.com/mymmrac/telego"
)

var ErrNotMember = errors.New("user is not a chat member")

func IsMember(ctx context.Context, bot *telego.Bot, chatID, userID int64) (bool, error) {
	chatMember, err := bot.GetChatMember(ctx, &telego.GetChatMemberParams{
		ChatID: telego.ChatID{ID: chatID},
		UserID: userID,
	})
	if err != nil {
		return false, err
	}

	switch chatMember.MemberStatus() {
	case telego.MemberStatusLeft, telego.MemberStatusBanned:
		return false, nil
	default:
		return true, nil
	}
}
