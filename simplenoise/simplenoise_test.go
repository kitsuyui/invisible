package simplenoise

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestAddRandomNoiseAndDeNoise(t *testing.T) {
	AddRandomNoiseAndDeNoiseIsReversive(t, 1.0, 1, "Hello, Noise!")
	AddRandomNoiseAndDeNoiseIsReversive(t, 2.0, 10, "あいうえお")
	AddRandomNoiseAndDeNoiseIsReversive(t, 2.0, 10, "")
	AddRandomNoiseAndDeNoiseIsReversive(t, 2.0, 10, "\u202BThis is Safe\u202C")
}

func AddRandomNoiseAndDeNoiseIsReversive(t *testing.T, frequency float64, maxSize int, testText string) {
	original := testText
	reader := bufio.NewReader(strings.NewReader(original))
	b := new(bytes.Buffer)
	writer := bufio.NewWriter(b)
	AddRandomNoise(frequency, maxSize, reader, writer)
	converted := b.String()
	if original == converted && original != "" {
		t.Errorf("%q == %q", original, converted)
	}
	reader = bufio.NewReader(bytes.NewReader(b.Bytes()))
	b2 := new(bytes.Buffer)
	writer = bufio.NewWriter(b2)
	DeNoise(reader, writer)
	recovered := b2.String()
	if original != recovered {
		t.Errorf("%q != %q", original, recovered)
	}
}
