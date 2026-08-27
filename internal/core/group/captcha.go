package group

import "fmt"

func Captcha(TGID model.User) bool {
	fmt.Println("user join request:", TGID)
	return true
}
