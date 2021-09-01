package main

import (
	"fmt"
	"github.com/xerosic/emojipasta/generator"
	"math/rand"
	"time"
)

func Generate(phrase string, config generator.EmojipastaConfig) string{
	rand.Seed(time.Now().UnixNano())
	splitWords := generator.SplitWordByDelimiter(phrase, config.WordDelimiter)
	return generator.AddEmojisToSlice(splitWords, config.MaxEmojisPerWord, config.FixedEmojiNumber)
}

func main() {
	cfg := generator.EmojipastaConfig{
		WordDelimiter:    " ",
		MaxEmojisPerWord: 2,
		FixedEmojiNumber: false,
	}
	fmt.Println(Generate("This library allows you to create emojipastas like this one", cfg))
}
