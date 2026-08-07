package rule

import (
	"github.com/mefoxtrot/meNFlubot/internal/info"
	"github.com/mymmrac/telego"
)

func IsDirect(chatType string) bool {
	return chatType == telego.ChatTypePrivate
}

// var ErrNotFoss = errors.New("this chat is not meNFoss")

func IsFoss(chatID int64) bool {
	return info.Foss.ID == chatID
}

// var ErrNotFozz = errors.New("this chat is not meNFozz")

func IsFozz(chatID int64) bool {
	return info.Fozz.ID == chatID
}
