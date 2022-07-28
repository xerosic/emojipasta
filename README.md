# emojipasta
This:arrow_up: library allows you:point_right: to:blush::two_men_holding_hands::speak_no_evil: create:calling::calling::calling: emojipastas:smiley::smiley: like:triumph::grin: this❓🏻 one:triumph::wheelchair:

## This is just a porting of a [porting](https://github.com/DavideGalilei/emojipasta) that a friend of mine made.
- The code is trash and was done just for fun, this thing is not optimized. Have fun with it!

- Uses [edb.prebuilt.json](https://github.com/oldpepper12/emojibot2/blob/master/edb.prebuilt.json) from [emojibot2](https://github.com/oldpepper12/emojibot2)

## Example Code

```go
package main

import (
	"fmt"
	"strings"

	"github.com/xerosic/emojipasta"
)

func main() {
	input := "This library allows you to create emojipastas like this one"
	output := emojipasta.Generate(strings.Split(input, " "), 3, false)
	fmt.Println(output)
}
```