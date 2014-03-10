package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	"github.com/nf/sigourney/audio"
	"github.com/nf/sigourney/ui"
)

var (
	length  = flag.Duration("len", 5*time.Second, "recording length")
	outFile = flag.String("out", "out.mp3", "mp3 output file name")
	inFile  = flag.String("in", "cacophon", "patch input file name")
)

func main() {
	flag.Parse()
	u := ui.New(noopHandler{})
	if err := u.Load(*inFile); err != nil {
		log.Fatal(err)
	}
	u.Set("value32", 0) // delay
	u.Set("value22", 0) // speed
	u.Set("value17", 0) // fm
	if err := render(u, *length, *outFile); err != nil {
		log.Fatal(err)
	}
}

func render(u *ui.UI, d time.Duration, filename string) error {
	frames := int(d / time.Second * 44100 / 256)
	samp := u.Render(frames)
	normalize(samp)
	b, err := mp3(pcm(samp))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0644)
}

func mp3(pcm []byte) ([]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("lame", "-m", "mono", "-r", "-", "-")
	cmd.Stdin = bytes.NewReader(pcm)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
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
	for _, v := range s {
		if v < 0 {
			v *= -1
		}
		if v > max {
			max = v
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
