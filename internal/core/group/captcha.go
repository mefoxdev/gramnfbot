package group

import "log/slog"

func Captcha(
	TGID int64,
) bool {
	slog.Debug("captcha request", "user_id", TGID)
	return true
}
