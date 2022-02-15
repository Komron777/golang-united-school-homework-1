package homework1

import "github.com/kyokomi/emoji"

func HelloWithEmoji () string{
	world_map := ":world_map:"
	result := emoji.Sprint("Hello " + world_map + "!")
	return result
}