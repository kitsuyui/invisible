package simplenoise

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"testing"
)

var errSimpleNoiseWriter = errors.New("simplenoise writer failed")

type failingSimpleNoiseWriter struct{}

func (failingSimpleNoiseWriter) Write(_ []byte) (int, error) {
	return 0, errSimpleNoiseWriter
}

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
	if err := AddRandomNoise(frequency, maxSize, reader, writer); err != nil {
		t.Fatal(err)
	}
	converted := b.String()
	if original == converted && original != "" {
		t.Errorf("%q == %q", original, converted)
	}
	reader = bufio.NewReader(bytes.NewReader(b.Bytes()))
	b2 := new(bytes.Buffer)
	writer = bufio.NewWriter(b2)
	if err := DeNoise(reader, writer); err != nil {
		t.Fatal(err)
	}
	recovered := b2.String()
	if original != recovered {
		t.Errorf("%q != %q", original, recovered)
	}
}

func TestAddRandomNoiseReturnsWriterError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("Hello"))
	writer := bufio.NewWriter(failingSimpleNoiseWriter{})
	if err := AddRandomNoise(0, 0, reader, writer); !errors.Is(err, errSimpleNoiseWriter) {
		t.Fatalf("AddRandomNoise() error = %v, want %v", err, errSimpleNoiseWriter)
	}
}

func TestDeNoiseWritesToPlainWriter(t *testing.T) {
	// Regression: DeNoise must flush to a plain io.Writer, not just *bufio.Writer.
	noisy := "H⁢ello" // H + INVISIBLE TIMES + ello
	b := new(bytes.Buffer)
	if err := DeNoise(strings.NewReader(noisy), b); err != nil {
		t.Fatal(err)
	}
	if got := b.String(); got != "Hello" {
		t.Errorf("DeNoise() = %q, want %q", got, "Hello")
	}
}
func TestDeNoiseReturnsWriterError(t *testing.T) {
	if err := DeNoise(strings.NewReader("Hello"), failingSimpleNoiseWriter{}); !errors.Is(err, errSimpleNoiseWriter) {
		t.Fatalf("DeNoise() error = %v, want %v", err, errSimpleNoiseWriter)
	}
}
