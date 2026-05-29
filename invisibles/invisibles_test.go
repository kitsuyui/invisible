package invisibles

import "testing"

func TestGetInvisibleRune(t *testing.T) {
	GetInvisibleRune()
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
	IsGetInvisibleRuneTester(t, EncodingFormatMarkerRune, true)
	if InvisibleRuneCode(EncodingFormatMarkerRune) != -1 {
		t.Errorf("%#U must be reserved for metadata, not mapped as data", EncodingFormatMarkerRune)
	}
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
