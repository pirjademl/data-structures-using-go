package strings

import (
	"fmt"
	"strings"
)

func ReverseString(str string) {
	builder := strings.Builder{}
	for i := 0; i < len(str); i++ {
		char := str[i]
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// it is a character
			var word []byte

			for i > 0 && str[i] != ' ' {
				word = append(word, str[i])
				i++
			}
			builder.Write(word)

		}
	}
	fmt.Println(builder.String())
}
