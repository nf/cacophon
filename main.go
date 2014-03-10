package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"os/exec"
	"time"

	"github.com/nf/sigourney/audio"
	"github.com/nf/sigourney/ui"
)

func main() {
	u := ui.New(noopHandler{})
	if err := u.Load("simple"); err != nil {
		log.Fatal(err)
	}
	if err := render(u, 2*time.Second, "out.mp3"); err != nil {
		log.Fatal(err)
	}
}

func render(u *ui.UI, d time.Duration, filename string) error {
	frames := int(d / time.Second * 44100 / 256)
	samp := u.Render(frames)
	normalize(samp)

	cmd := exec.Command("lame", "-m", "mono", "-r", "-", filename)
	cmd.Stdin = bytes.NewReader(pcm(samp))
	return cmd.Run()
}

func pcm(samp []audio.Sample) []byte {
	data := make([]byte, len(samp)*2)
	for i := range samp {
		s := uint16(samp[i] * (1 << 15))
		binary.LittleEndian.PutUint16(data[2*i:], s)
	}
	return data
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
