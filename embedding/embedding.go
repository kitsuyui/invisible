package embedding

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"

	"github.com/kitsuyui/invisible/invisibles"
	"github.com/kitsuyui/invisible/simplenoise"
)

var INVISIBLE_RUNES_TO_UINT32 = []uint32{
	binary.BigEndian.Uint32([]byte{0x00, 0xe0, 0x00, 0x00}), // 0b111000000000000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x1c, 0x00, 0x00}), // 0b000111000000000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x03, 0x80, 0x00}), // 0b000000111000000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x70, 0x00}), // 0b000000000111000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x0e, 0x00}), // 0b000000000000111000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x01, 0xc0}), // 0b000000000000000111000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x38}), // 0b000000000000000000111000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x07}), // 0b000000000000000000000111
}

func Embed(embedString string, reader *bufio.Reader, writer *bufio.Writer, repeat bool) {
	encoded := []rune(Encode(embedString))
	isFirst := true
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if !isFirst {
			if len(encoded) > 0 {
				writer.WriteRune(encoded[0])
				encoded = encoded[1:]
			}
			if repeat && len(encoded) == 0 {
				encoded = []rune(Encode(embedString))
			}
		}
		isFirst = false
		writer.WriteRune(r)
	}
	for len(encoded) > 0 {
		writer.WriteRune(encoded[0])
		encoded = encoded[1:]
	}
	writer.Flush()
}

func Extract(reader *bufio.Reader, writer *bufio.Writer) string {
	b := new(bytes.Buffer)
	noiseWriter := bufio.NewWriter(b)
	simplenoise.DeNoiseAndWriteNoise(reader, writer, noiseWriter)
	decoded := Decode(b.String())
	return decoded
}

func Encode(text string) string {
	var encoded []rune
	bs := []byte(text)
	byteSizeEncodeFrom := 3
	runeSizeEncodeTo := 8
	for i := 0; i < len(bs); i += byteSizeEncodeFrom {
		offset := i + byteSizeEncodeFrom
		if offset > len(bs) {
			offset = len(bs)
		}
		padded := padLastToNBytes(byteSizeEncodeFrom, bs[i:offset])
		runes := encodeNbytesToMRunes(byteSizeEncodeFrom, runeSizeEncodeTo, padded)
		encoded = append(encoded, runes...)
	}
	return string(encoded)
}

func Decode(noiseText string) string {
	var decoded []byte
	runes := []rune(noiseText)
	runeSizeDecodeFrom := 8
	byteSizeDecodeTo := 3
	for i := 0; i < len(runes); i += runeSizeDecodeFrom {
		offset := i + runeSizeDecodeFrom
		if offset > len(runes) {
			offset = len(runes)
		}
		bs := decodeNRunesToNBytes(runeSizeDecodeFrom, byteSizeDecodeTo, runes[i:offset])
		trimmed := bytes.Trim(bs, "\x00")
		decoded = append(decoded, trimmed...)
	}
	return string(decoded)
}

func encodeNbytesToMRunes(n int, m int, bs []byte) (rs []rune) {
	uint32Bytes := 4
	bitsPerBytes := 8
	bin := binary.BigEndian.Uint32(padFirstToNBytes(uint32Bytes, bs))
	for i, matching := range INVISIBLE_RUNES_TO_UINT32 {
		shift := uint(n * (bitsPerBytes - i - 1))
		code := (bin & matching) >> shift
		rs = append(rs, invisibles.INVISIBLE_RUNES[code])
	}
	return rs
}

func decodeNRunesToNBytes(n int, m int, rs []rune) []byte {
	decodedUint := uint32(0)
	bitsPerBytes := 8
	uint32Bytes := 4
	cut := uint32Bytes - m
	for i, r := range rs {
		k := uint32(invisibleRuneToCode(r))
		shift := uint32(m * (bitsPerBytes - i - 1))
		decodedUint |= k << shift
	}
	bs := make([]byte, uint32Bytes)
	binary.BigEndian.PutUint32(bs, decodedUint)
	return bs[cut:]
}

func invisibleRuneToCode(r rune) int {
	for i, r2 := range invisibles.INVISIBLE_RUNES {
		if r2 == r {
			return i
		}
	}
	return -1
}

func padLastToNBytes(n int, bs []byte) []byte {
	requires := n - len(bs)
	padding := make([]byte, requires)
	return append(bs, padding...)
}

func padFirstToNBytes(n int, bs []byte) []byte {
	requires := n - len(bs)
	padding := make([]byte, requires)
	return append(padding, bs...)
}
