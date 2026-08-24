package rule

import (
	"github.com/mymmrac/telego"
)

func GetTarget(
	msg telego.Message,
) int64 {
	text := msg.Text
	var targetID int64

	// try to find target id
	if msg.ReplyToMessage.From != nil {
		targetID = msg.ReplyToMessage.From.ID
	}
	//! сделать отделение команды от причины
	text = text + "s"
	return targetID
}
