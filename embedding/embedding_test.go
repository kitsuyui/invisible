package embedding

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/kitsuyui/invisible/invisibles"
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

func TestEncodeAndDecodeDoesNotUseMutableRuneCopies(t *testing.T) {
	runes := invisibles.InvisibleRunes()
	runes[0] = 'A'

	CheckEncodeAndDecodeIsReversible(t, "Hello, World!")
}

func TestDecodeBroken(t *testing.T) {
	TryDecodeBroken(t, "にいはお")
}

func TestDecodeSkipsUnknownRunes(t *testing.T) {
	if decoded := Decode("A"); decoded != "" {
		t.Fatalf("Decode() = %q, want empty string for unknown runes", decoded)
	}
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

func TestEmbedRepeatEmbedsDifferentlyFromNoRepeat(t *testing.T) {
	message := "Hi"
	// Host text longer than what's needed to embed the message once.
	// Encode("Hi") produces 8*ceil(2/3) = 8 invisible runes; host text of 200 chars
	// ensures encoded runes would be exhausted well before the host ends.
	hostText := strings.Repeat("Hello World! ", 20)

	embed := func(rep bool) string {
		reader := bufio.NewReader(strings.NewReader(hostText))
		b := new(bytes.Buffer)
		writer := bufio.NewWriter(b)
		if err := Embed(message, reader, writer, rep); err != nil {
			t.Fatal(err)
		}
		return b.String()
	}

	withRepeat := embed(true)
	withoutRepeat := embed(false)

	if withRepeat == withoutRepeat {
		t.Error("Embed with repeat=true should produce different output than repeat=false for long host text")
	}
}

func TestEmbedRepeatExtractRecoversMessage(t *testing.T) {
	message := "Hi"
	hostText := strings.Repeat("Hello World! ", 20)

	reader := bufio.NewReader(strings.NewReader(hostText))
	b := new(bytes.Buffer)
	writer := bufio.NewWriter(b)
	if err := Embed(message, reader, writer, true); err != nil {
		t.Fatal(err)
	}

	reader2 := bufio.NewReader(bytes.NewReader(b.Bytes()))
	b2 := new(bytes.Buffer)
	writer2 := bufio.NewWriter(b2)
	decoded, err := Extract(reader2, writer2)
	if err != nil {
		t.Fatal(err)
	}
	// Extract decodes all embedded invisible characters. With repeat the decoded
	// string will be message repeated; check that it starts with the original message.
	if !strings.HasPrefix(decoded, message) {
		t.Errorf("Extract after Embed(repeat=true): got %q, want prefix %q", decoded, message)
	}
}

func TestEncodeAndDecodeWithNullBytes(t *testing.T) {
	// Null byte at the start of a chunk (was incorrectly stripped by bytes.Trim).
	CheckEncodeAndDecodeIsReversible(t, "\x00AB")
	// Null byte in the middle of a chunk (round-trips correctly before and after fix).
	CheckEncodeAndDecodeIsReversible(t, "A\x00B")
	// Null byte at the start of a non-last chunk (was incorrectly stripped).
	CheckEncodeAndDecodeIsReversible(t, "ABC\x00EF")
	// Null bytes spanning multiple chunks with leading null.
	CheckEncodeAndDecodeIsReversible(t, "\x00ABCDEF")
}
