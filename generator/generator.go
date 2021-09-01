package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type EmojipastaConfig struct {
	WordDelimiter    string
	MaxEmojisPerWord int
	FixedEmojiNumber bool
}

func SplitWordByDelimiter(word string, delimiter string) []string {
	return strings.Split(word, delimiter)
}

func AddEmojisToSlice(phrase []string, maxEmojisPerWord int, fixedEmojiNumber bool) string {
	jsonFile, err := os.Open("assets/emojis.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	var emojis []string

	var finalString string

	err2 := json.Unmarshal(byteValue, &result)
	if err2 != nil {
		return ""
	}

	if fixedEmojiNumber {
		for j, word := range phrase {
			if j == 0 {
				finalString += word
			} else {
				finalString += " " + word
			}

			emojis = strings.Split(fmt.Sprintf("%v", result[strings.ToLower(word)]), " ")
			for i := 1; i <= maxEmojisPerWord; i++ {
				index := rand.Intn(len(emojis))
				if emojis[index] != "<nil>" {
					finalString += emojis[index]
				}
			}
		}
	} else {
		rnd := rand.Intn(maxEmojisPerWord)
		for j, word := range phrase {
			if j == 0 {
				finalString += word
			} else {
				finalString += " " + word
			}

			emojis = strings.Split(fmt.Sprintf("%v", result[strings.ToLower(word)]), " ")
			for i := 0; i <= rnd; i++ {
				index := rand.Intn(len(emojis))
				if emojis[index] != "<nil>" {
					finalString += emojis[index]
				}
			}
			rnd = rand.Intn(maxEmojisPerWord)
		}
	}
	return strings.ReplaceAll(strings.ReplaceAll(finalString, "[", ""), "]", "")

}
