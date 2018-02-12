package simplenoise

import (
	"bufio"
	"io"
	"io/ioutil"
	"math/rand"

	"github.com/kitsuyui/invisible/invisibles"
)

func AddRandomNoise(frequency float64, maxSize int, reader *bufio.Reader, writer *bufio.Writer) {
	isFirst := true
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if !isFirst {
			for i := 0; i < maxSize; i++ {
				if rand.Float64() < frequency {
					ir := invisibles.GetInvisibleRune()
					writer.WriteRune(ir)
				}
			}
		}
		isFirst = false
		writer.WriteRune(r)
	}
	writer.Flush()
}

func DeNoise(reader io.Reader, writer io.Writer) {
	bufreader := bufio.NewReader(reader)
	bufwriter := bufio.NewWriter(writer)
	DeNoiseAndWriteNoise(bufreader, bufwriter, bufio.NewWriter(ioutil.Discard))
}

func DeNoiseAndWriteNoise(reader io.Reader, writer io.Writer, noiseWriter io.Writer) {
	bufReader := bufio.NewReader(reader)
	bufWriter := bufio.NewWriter(writer)
	noiseBufWriter := bufio.NewWriter(noiseWriter)
	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if invisibles.IsGetInvisibleRune(r) {
			noiseBufWriter.WriteRune(r)
		} else {
			bufWriter.WriteRune(r)
		}
	}
	bufWriter.Flush()
	noiseBufWriter.Flush()
}
