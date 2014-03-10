package main

import (
	"encoding/binary"
	"log"
	"os"
	"time"

	"github.com/nf/sigourney/audio"
	"github.com/nf/sigourney/ui"
	"github.com/nf/wav"
)

func main() {
	u := ui.New(noopHandler{})
	if err := u.Load("simple"); err != nil {
		log.Fatal(err)
	}
	if err := render(u, 2*time.Second, "out.wav"); err != nil {
		log.Fatal(err)
	}
}

func render(u *ui.UI, d time.Duration, filename string) error {
	frames := int(d / time.Second * 44100 / 256)
	samp := u.Render(frames)
	normalize(samp)
	data := make([]byte, len(samp)*2)
	for i := range samp {
		s := uint16(samp[i] * (1 << 15))
		binary.LittleEndian.PutUint16(data[2*i:], s)
	}
	w := &wav.File{
		SampleRate:      44100,
		SignificantBits: 16,
		Channels:        1,
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := w.WriteData(f, data); err != nil {
		return err
	}
	return f.Close()
}

func normalize(s []audio.Sample) {
	max := audio.Sample(0)
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	if max > 0 {
		fac := 0.99 / max
		for i := range s {
			s[i] *= fac
		}
	}
}

type noopHandler struct{}

func (noopHandler) Hello(kindInputs map[string][]string) {}
func (noopHandler) SetGraph(graph []*ui.Object)          {}
