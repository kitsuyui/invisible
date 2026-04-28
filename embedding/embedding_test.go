package embedding

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"testing"
)

var errEmbeddingWriter = errors.New("embedding writer failed")

type failingEmbeddingWriter struct{}

func (failingEmbeddingWriter) Write(_ []byte) (int, error) {
	return 0, errEmbeddingWriter
}

func TestSomething(t *testing.T) {
	original := "Hello, World!"
	reader := bufio.NewReader(strings.NewReader(original))
	b := new(bytes.Buffer)
	writer := bufio.NewWriter(b)
	if err := Embed(original, reader, writer, false); err != nil {
		t.Fatal(err)
	}
	converted := b.String()
	if original == converted && original != "" {
		t.Errorf("%q == %q", original, converted)
	}
	reader = bufio.NewReader(bytes.NewReader(b.Bytes()))
	b2 := new(bytes.Buffer)
	writer = bufio.NewWriter(b2)
	decoded, err := Extract(reader, writer)
	if err != nil {
		t.Fatal(err)
	}
	if original != decoded {
		t.Errorf("Must be original = decoded (%q = %q) (decoded)", original, decoded)
	}
}

func TestEmbedReturnsWriterError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("Hello"))
	writer := bufio.NewWriter(failingEmbeddingWriter{})
	if err := Embed("hidden", reader, writer, false); !errors.Is(err, errEmbeddingWriter) {
		t.Fatalf("Embed() error = %v, want %v", err, errEmbeddingWriter)
	}
}

func TestExtractReturnsWriterError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("Hello"))
	writer := bufio.NewWriter(failingEmbeddingWriter{})
	if _, err := Extract(reader, writer); !errors.Is(err, errEmbeddingWriter) {
		t.Fatalf("Extract() error = %v, want %v", err, errEmbeddingWriter)
	}
}

func TestEncodeAndDecodeIsReversible(t *testing.T) {
	CheckEncodeAndDecodeIsReversible(t, "Hello, World!")
	CheckEncodeAndDecodeIsReversible(t, "Hello, World!!")
	CheckEncodeAndDecodeIsReversible(t, "Hello, World!!!")
	CheckEncodeAndDecodeIsReversible(t, "Hello, World!!!!")
	CheckEncodeAndDecodeIsReversible(t, "Hello, World!!!!!")
	CheckEncodeAndDecodeIsReversible(t, "Good Morning")
}

func TestDecodeBroken(t *testing.T) {
	TryDecodeBroken(t, "にいはお")
}

func TryDecodeBroken(t *testing.T, text string) {
	encoded := Encode(text)
	Decode(encoded[:len(encoded)-1])
}

func CheckEncodeAndDecodeIsReversible(t *testing.T, text string) {
	encoded := Encode(text)
	decoded := Decode(encoded)
	if text == encoded {
		t.Errorf("Encode failure. It seems to be no-op.")
	}
	if text != decoded {
		t.Errorf("Must be original = decoded (%q = %q) (decoded)", text, decoded)
	}
}

func TestInvisibleRuneToCodeFailure(t *testing.T) {
	if invisibleRuneToCode('A') != -1 {
		t.Errorf("When passed invisible rune then must be return -1")
	}
}
