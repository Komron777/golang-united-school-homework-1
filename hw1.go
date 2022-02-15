package homework1

import "github.com/kyokomi/emoji"

func HelloWithEmoji () string{
	world := ":world_map:"
	return emoji.Sprint("Hello ",world,"!")
}