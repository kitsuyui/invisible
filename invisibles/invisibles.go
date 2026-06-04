package invisibles

import (
	"math/rand"
	"sync"
	"time"
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

// EncodingFormatMarkerRune is reserved for encoding format markers. It is
// extracted as invisible metadata, but is not part of the legacy data alphabet.
const EncodingFormatMarkerRune = '\u200D' // ZERO WIDTH JOINER

func InvisibleRunes() []rune {
	return append([]rune(nil), invisibleRunes[:]...)
}

var invisibleRuneRand = newLockedRand(time.Now().UnixNano())

type lockedRand struct {
	mu sync.Mutex
	r  *rand.Rand
}

func newLockedRand(seed int64) *lockedRand {
	return &lockedRand{r: rand.New(rand.NewSource(seed))}
}

func (r *lockedRand) Intn(n int) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.r.Intn(n)
}

func GetInvisibleRune() rune {
	return invisibleRunes[invisibleRuneRand.Intn(len(invisibleRunes))]
}

func IsGetInvisibleRune(r rune) bool {
	return r == EncodingFormatMarkerRune || InvisibleRuneCode(r) >= 0
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
