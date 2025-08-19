package strings

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func GetBigString(charLen uint) string {
	str := ""
	for i := 1; i <= int(charLen); i++ {
		str += fmt.Sprint((rand.IntN(256)))
	}
	return str
}

func GetBigStringBuilder(charLen uint) string {
	str := &strings.Builder{}
	for i := 1; i <= int(charLen); i++ {
		str.WriteString(fmt.Sprint((rand.IntN(256))))
	}
	return str.String()
}
