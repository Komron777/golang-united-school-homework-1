package homework1

import "github.com/kyokomi/emoji"

func HelloWithEmoji () string{
	result := emoji.Sprint("Hello :world_map:!")
	return result
}