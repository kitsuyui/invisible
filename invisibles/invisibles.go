package invisibles

import (
	"math/rand"
	"time"
)

var INVISIBLE_RUNES = []rune{
	// '\u180E', // MONGOLIAN VOWEL SEPARATOR
	'\u200B', // ZERO WIDTH SPACE
	'\u200C', // ZERO WIDTH NON-JOINER
	// '\u200D', // ZERO WIDTH JOINER
	'\u2060', // WORD JOINER
	'\u2061', // FUNCTION APPLICATION
	'\u2062', // INVISIBLE TIMES
	'\u2063', // INVISIBLE SEPARATOR
	'\u2064', // INVISIBLE PLUS
	'\uFEFF', // ZERO WIDTH NO-BREAK SPACE
}

func init() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
}

func GetInvisibleRune() rune {
	return INVISIBLE_RUNES[rand.Intn(len(INVISIBLE_RUNES))]
}

func IsGetInvisibleRune(r rune) bool {
	for _, r2 := range INVISIBLE_RUNES {
		if r == r2 {
			return true
		}
	}
	return false
}
