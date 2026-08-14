package show

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const profileURL = "https://netfox.me/profile"

func Profile() *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Profile").
				WithWebApp(tu.WebAppInfo(profileURL)),
		),
	)
}
