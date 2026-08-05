package info

const (
	fossID = -1003974229978
	fozzID = -1003956046051
)

type Chazz struct {
	ID      int64
	TChatID int
	TMainID int
}

type Chass struct {
	ID int64
}

var Foss = Chass{
	ID: fossID,
}

var Fozz = Chazz{
	ID:      fozzID,
	TMainID: 1,
}

var (
	FozzRules = `
правила чата meNFozz:
	`
	FossRules = `
правила чата meNFoss:
запрещено:
1. плохо себя вести -> наказание я выберу сам
2. менять сообщения -> бот их удаляет и дает вам варн
info: правила я буду добавлять по ходу разработки
info: правила тут, это короткая сводка, полные правила, или же «Конституция netfox.me» позже будут на нашем сайте 
	`

	BaseRules = `
правила чата:
info: будут добавлены позже
	`
)

// брать значения позже из .env с помощью joho/godotenv
