package simplenoise

import (
	"bufio"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/kitsuyui/invisible/invisibles"
)

var noiseRand = newLockedRand(time.Now().UnixNano())

type lockedRand struct {
	mu sync.Mutex
	r  *rand.Rand
}

func newLockedRand(seed int64) *lockedRand {
	return &lockedRand{r: rand.New(rand.NewSource(seed))}
}

func (r *lockedRand) Float64() float64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.r.Float64()
}

// AddRandomNoise inserts invisible Unicode characters between runes read from
// reader and writes the result to writer. frequency controls the per-slot
// probability of inserting a noise character: values in [0.0, 1.0] behave as
// a probability (0.0 = no noise, 1.0 = noise on every slot). Values above 1.0
// guarantee insertion on every slot; values below 0.0 suppress all noise.
// maxSize is the number of noise-insertion slots between consecutive runes.
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
				if noiseRand.Float64() < frequency {
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
	return DeNoiseAndWriteNoise(reader, writer, io.Discard)
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
