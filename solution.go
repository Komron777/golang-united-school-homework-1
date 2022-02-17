package solution

import "github.com/kyokomi/emoji"

func GetMessage() string{
	world := ":world_map:"
	return emoji.Sprint("Hello ",world,"!")
}