package backend

import (
	"bytes"
	"encoding/binary"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/nf/sigourney/audio"
	"github.com/nf/sigourney/ui"
)

const (
	length = 5 * time.Second
	inFile = "cacophon.patch"
)

func init() {
	http.HandleFunc("/audio", audioHandler)
}

func audioHandler(w http.ResponseWriter, r *http.Request) {
	u := ui.New(noopHandler{})
	if err := u.Load(inFile); err != nil {
		log.Println(err)
		http.Error(w, "error processing audio", 500)
		return
	}

	u.Set("value19", floatValue(r, "speed"))
	//u.Set("x", floatValue(r, "scale"))
	//u.Set("x", floatValue(r, "perm"))
	u.Set("value70", floatValue(r, "slew"))
	u.Set("value77", floatValue(r, "root"))
	u.Set("value4", floatValue(r, "square"))
	u.Set("value39", floatValue(r, "amount"))
	u.Set("value46", floatValue(r, "offset"))
	u.Set("value13", floatValue(r, "attack"))
	u.Set("value12", floatValue(r, "decay"))
	u.Set("value29", floatValue(r, "time"))
	u.Set("value27", floatValue(r, "feedback"))

	frames := int(length / time.Second * 44100 / 256)
	samp := u.Render(frames)
	normalize(samp)
	b, err := mp3(pcm(samp))
	if err != nil {
		log.Println(err)
		http.Error(w, "error processing audio", 500)
		return
	}

	w.Header().Set("Content-type", "audio/mp3")
	if _, err := w.Write(b); err != nil {
		log.Println(err)
	}
}

func floatValue(r *http.Request, name string) float64 {
	f, _ := strconv.ParseFloat(r.FormValue(name), 64)
	if f > 1 {
		f = 1
	} else if f < 0 {
		f = 0
	}
	return f
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
