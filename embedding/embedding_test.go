package embedding

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestSomething(t *testing.T) {
	original := "Hello, World!"
	reader := bufio.NewReader(strings.NewReader(original))
	b := new(bytes.Buffer)
	writer := bufio.NewWriter(b)
	Embed(original, reader, writer, false)
	converted := b.String()
	if original == converted && original != "" {
		t.Errorf("%q == %q", original, converted)
	}
	reader = bufio.NewReader(bytes.NewReader(b.Bytes()))
	b2 := new(bytes.Buffer)
	writer = bufio.NewWriter(b2)
	decoded := Extract(reader, writer)
	if original != decoded {
		t.Errorf("Must be original = decoded (%q = %q) (decoded)", original, decoded)
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
