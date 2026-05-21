package invisibles

import (
	"math/rand"
	"testing"
)

func TestGetInvisibleRune(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	r := GetInvisibleRune(rng)
	if !IsGetInvisibleRune(r) {
		t.Errorf("GetInvisibleRune() = %#U, want an invisible rune", r)
	}
}

func TestInvisibleRunesReturnsCopy(t *testing.T) {
	runes := InvisibleRunes()
	runes[0] = 'A'

	if IsGetInvisibleRune('A') {
		t.Errorf("mutating returned runes must not mutate the package rune table")
	}
	if InvisibleRuneCode('\u200B') != 0 {
		t.Errorf("mutating returned runes must not change rune codes")
	}
}

func TestIsGetInvisibleRune(t *testing.T) {
	IsGetInvisibleRuneTester(t, 'x', false)
	IsGetInvisibleRuneTester(t, '\u2060', true)
}

func IsGetInvisibleRuneTester(t *testing.T, testRune rune, shouldBe bool) {
	if IsGetInvisibleRune(testRune) != shouldBe {
		if shouldBe {
			t.Errorf("%#U should be treated as invisible", testRune)
		} else {
			t.Errorf("%#U should be treated as visible", testRune)
		}
	}
}
