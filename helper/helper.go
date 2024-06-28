package helper

import (
	"math/rand"
	"time"
	"unicode/utf8"

	"github.com/DanPlayer/randomname"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func rangeIn(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Random5DigitNumber() int32 {
	return int32(rangeIn(10000, 99999))
}

func RandomChineseName(max int) string {
	min := 3
	name := ""
	for {
		name = randomname.GenerateName()
		size := utf8.RuneCountInString(name)
		if size >= min && size <= max {
			break
		}
	}

	return name
}
