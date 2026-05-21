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
