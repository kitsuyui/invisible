package simplenoise

import (
	"bufio"
	"io"
	"io/ioutil"
	"math/rand"

	"github.com/kitsuyui/invisible/invisibles"
)

func AddRandomNoise(frequency float64, maxSize int, reader *bufio.Reader, writer *bufio.Writer) error {
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
			for i := 0; i < maxSize; i++ {
				if rand.Float64() < frequency {
					ir := invisibles.GetInvisibleRune()
					if _, err := writer.WriteRune(ir); err != nil {
						return err
					}
				}
			}
		}
		isFirst = false
		if _, err := writer.WriteRune(r); err != nil {
			return err
		}
	}
	return writer.Flush()
}

func DeNoise(reader io.Reader, writer io.Writer) error {
	bufreader := bufio.NewReader(reader)
	bufwriter := bufio.NewWriter(writer)
	return DeNoiseAndWriteNoise(bufreader, bufwriter, bufio.NewWriter(ioutil.Discard))
}

func DeNoiseAndWriteNoise(reader io.Reader, writer io.Writer, noiseWriter io.Writer) error {
	bufReader := bufio.NewReader(reader)
	bufWriter := bufio.NewWriter(writer)
	noiseBufWriter := bufio.NewWriter(noiseWriter)
	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if invisibles.IsGetInvisibleRune(r) {
			if _, err := noiseBufWriter.WriteRune(r); err != nil {
				return err
			}
		} else {
			if _, err := bufWriter.WriteRune(r); err != nil {
				return err
			}
		}
	}
	if err := bufWriter.Flush(); err != nil {
		return err
	}
	return noiseBufWriter.Flush()
}
