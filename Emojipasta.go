package emojipasta

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"regexp"
	"strings"
)

var (
	//go:embed edb.prebuilt.json
	_jsonFile      []byte
	_emojiDB       map[string][]string
	_matchBrackets = regexp.MustCompile("[\\[\\]]")
)

func init() {
	if err := json.Unmarshal(_jsonFile, &_emojiDB); err != nil {
		panic(err)
	}
}

func Generate(input []string, maxEmojisPerWord int, fixedEmojiNumber bool) (output string) {
	emojisPerWord := maxEmojisPerWord
	for k, word := range input {
		if k > 0 {
			output += " "
		}
		output += word

		if !fixedEmojiNumber {
			emojisPerWord = rand.Intn(maxEmojisPerWord)
		}

		if emojis, ok := _emojiDB[strings.ToLower(word)]; ok {
			for i := 0; i <= emojisPerWord; i++ {
				output += emojis[rand.Intn(len(emojis))]
			}
		}
	}

	return _matchBrackets.ReplaceAllString(output, "")
}
