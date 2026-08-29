package router

import (
	"context"

	"github.com/mefoxtrot/gramnfbot/internal/core/direct"
	"github.com/mefoxtrot/gramnfbot/internal/core/group"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

// Register регистрирует все хендлеры бота.
func Register(bh *th.BotHandler) {
	// ---
	// DIRECT
	bh.HandleMessage(
		direct.HandleStart,
		th.CommandEqual("start"),
		directChat,
	)

	// ---
	// GROUP
	bh.HandleMessage(
		group.HandleRules,
		th.CommandEqual("rules"),
		groupChat,
	)
	bh.HandleChatMember(group.HandleJoin)

	bh.HandleChatJoinRequest(group.HandleChatJoinRequest)
}

// Direct
func directChat(_ context.Context, update telego.Update) bool {
	if update.Message == nil {
		return false
	}

	return update.Message.Chat.Type == telego.ChatTypePrivate
}

// Group
func groupChat(_ context.Context, update telego.Update) bool {
	if update.Message == nil {
		return false
	}

	return update.Message.Chat.Type == telego.ChatTypeGroup ||
		update.Message.Chat.Type == telego.ChatTypeSupergroup
}
