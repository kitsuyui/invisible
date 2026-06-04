package embedding

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/kitsuyui/invisible/invisibles"
	"github.com/kitsuyui/invisible/simplenoise"
)

// maxNoiseBufBytes caps the noise buffer in Extract to guard against OOM from
// inputs crafted with large numbers of invisible runes.
const maxNoiseBufBytes int64 = 1 << 20 // 1 MiB

const (
	encodingFormatMarkerPrefix = string(invisibles.EncodingFormatMarkerRune) + string(invisibles.EncodingFormatMarkerRune)
	encodingFormatMarkerV1     = encodingFormatMarkerPrefix + "\u200B"
)

// ErrUnsupportedEncodingFormat reports an encoded payload with a reserved
// format marker prefix that this decoder does not support.
var ErrUnsupportedEncodingFormat = errors.New("unsupported invisible encoding format")

type limitedBuffer struct {
	buf *bytes.Buffer
	max int64
}

func (l *limitedBuffer) Write(p []byte) (int, error) {
	if int64(l.buf.Len())+int64(len(p)) > l.max {
		return 0, fmt.Errorf("noise buffer limit exceeded (%d bytes)", l.max)
	}
	return l.buf.Write(p)
}

var invisibleRunesToUint32 = [...]uint32{
	binary.BigEndian.Uint32([]byte{0x00, 0xe0, 0x00, 0x00}), // 0b111000000000000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x1c, 0x00, 0x00}), // 0b000111000000000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x03, 0x80, 0x00}), // 0b000000111000000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x70, 0x00}), // 0b000000000111000000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x0e, 0x00}), // 0b000000000000111000000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x01, 0xc0}), // 0b000000000000000111000000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x38}), // 0b000000000000000000111000
	binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x07}), // 0b000000000000000000000111
}

// Embed hides embedString in the host text from reader by interleaving
// invisible Unicode runes into writer.
//
// The encoding operates in two phases depending on the relative lengths of the
// host text and the encoded message rune sequence:
//
//  1. Interleave mode: while visible characters and encoded runes both remain,
//     one invisible rune is inserted immediately before each visible character,
//     except before the very first visible character (isFirst guard). This
//     distributes invisible runes evenly through the host text.
//
//  2. Trailing-block mode: if encoded runes remain after the host text is
//     exhausted (i.e. len([]rune(Encode(embedString))) > number of visible chars),
//     the remaining runes are appended as a contiguous block at the end of the
//     output. Contiguous invisible runes are more detectable by analysis tools
//     than interleaved ones; callers should prefer host text at least as long as
//     len([]rune(Encode(embedString))) to stay in interleave mode.
//
// The repeat parameter causes the encoded rune sequence to cycle back to the
// beginning once exhausted, but only applies in interleave mode while visible
// characters remain in the host text. It has no effect in trailing-block mode.
func Embed(embedString string, reader *bufio.Reader, writer *bufio.Writer, repeat bool) error {
	originalEncoded := []rune(encodeLegacy(embedString))
	encoded := append([]rune(encodingFormatMarkerV1), originalEncoded...)
	isFirst := true
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if !isFirst {
			if len(encoded) == 0 && repeat {
				encoded = originalEncoded
			}
			if len(encoded) > 0 {
				if _, err := writer.WriteRune(encoded[0]); err != nil {
					return err
				}
				encoded = encoded[1:]
			}
		}
		isFirst = false
		if _, err := writer.WriteRune(r); err != nil {
			return err
		}
	}
	for len(encoded) > 0 {
		if _, err := writer.WriteRune(encoded[0]); err != nil {
			return err
		}
		encoded = encoded[1:]
	}
	return writer.Flush()
}

func Extract(reader *bufio.Reader, writer *bufio.Writer) (string, error) {
	b := new(bytes.Buffer)
	lb := &limitedBuffer{buf: b, max: maxNoiseBufBytes}
	noiseWriter := bufio.NewWriter(lb)
	if err := simplenoise.DeNoiseAndWriteNoise(reader, writer, noiseWriter); err != nil {
		return "", err
	}
	return DecodeStrict(b.String())
}

func Encode(text string) string {
	return encodingFormatMarkerV1 + encodeLegacy(text)
}

func encodeLegacy(text string) string {
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
	decoded, err := DecodeStrict(noiseText)
	if err != nil {
		return ""
	}
	return decoded
}

// DecodeStrict decodes markerless legacy payloads and v1-marked payloads.
// Unsupported marked payloads return ErrUnsupportedEncodingFormat.
func DecodeStrict(noiseText string) (string, error) {
	switch {
	case strings.HasPrefix(noiseText, encodingFormatMarkerV1):
		return decodeLegacy(strings.TrimPrefix(noiseText, encodingFormatMarkerV1)), nil
	case strings.HasPrefix(noiseText, encodingFormatMarkerPrefix):
		return "", ErrUnsupportedEncodingFormat
	default:
		return decodeLegacy(noiseText), nil
	}
}

func decodeLegacy(noiseText string) string {
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
		// Only trim trailing null bytes from the last chunk (padding removal).
		// Non-last chunks are always full 3-byte groups with no padding.
		if offset >= len(runes) {
			bs = bytes.TrimRight(bs, "\x00")
		}
		decoded = append(decoded, bs...)
	}
	return string(decoded)
}

func encodeNbytesToMRunes(n int, m int, bs []byte) (rs []rune) {
	uint32Bytes := 4
	bitsPerBytes := 8
	bin := binary.BigEndian.Uint32(padFirstToNBytes(uint32Bytes, bs))
	for i, matching := range invisibleRunesToUint32 {
		shift := uint(n * (bitsPerBytes - i - 1))
		code := int((bin & matching) >> shift)
		r, ok := invisibles.InvisibleRune(code)
		if ok {
			rs = append(rs, r)
		}
	}
	return rs
}

func decodeNRunesToNBytes(n int, m int, rs []rune) []byte {
	decodedUint := uint32(0)
	bitsPerBytes := 8
	uint32Bytes := 4
	cut := uint32Bytes - m
	for i, r := range rs {
		code := invisibleRuneToCode(r)
		if code < 0 {
			continue
		}
		k := uint32(code)
		shift := uint32(m * (bitsPerBytes - i - 1))
		decodedUint |= k << shift
	}
	bs := make([]byte, uint32Bytes)
	binary.BigEndian.PutUint32(bs, decodedUint)
	return bs[cut:]
}

func invisibleRuneToCode(r rune) int {
	return invisibles.InvisibleRuneCode(r)
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
