package invisibles

import (
	"math/rand"
)

var invisibleRunes = [...]rune{
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

func InvisibleRunes() []rune {
	return append([]rune(nil), invisibleRunes[:]...)
}

func GetInvisibleRune(r *rand.Rand) rune {
	return invisibleRunes[r.Intn(len(invisibleRunes))]
}

func IsGetInvisibleRune(r rune) bool {
	return InvisibleRuneCode(r) >= 0
}

func InvisibleRune(index int) (rune, bool) {
	if index < 0 || index >= len(invisibleRunes) {
		return 0, false
	}
	return invisibleRunes[index], true
}

func InvisibleRuneCode(r rune) int {
	for i, r2 := range invisibleRunes {
		if r == r2 {
			return i
		}
	}
	return -1
}
