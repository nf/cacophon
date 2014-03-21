package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/nf/sigourney/audio"
	"github.com/nf/sigourney/ui"
)

const (
	length = 10 * time.Second
	inFile = "cacophon.patch"
)

func init() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/audio", audioHandler)
}

var seqNotes = map[int]string{
	0: "value50",
	1: "value51",
	2: "value52",
	3: "value53",
	4: "value59",
	5: "value58",
	6: "value57",
	7: "value56",
}

const mainPage = `<html><body>
<ul>
<li><a href="/audio?speed=0.5&scale=0&perm=0&slew=0&root=0.5&square=0&amount=0&offset=0.5&attack=0&decay=0.25&time=0&feedback=0">/audio?speed=0.5&scale=0&perm=0&slew=0&root=0.5&square=0&amount=0&offset=0.5&attack=0&decay=0.25&time=0&feedback=0</a></li>
<li><a href="/audio?speed=0.75&scale=0&perm=0&slew=0.07&root=0.5&square=0&amount=0&offset=0.5&attack=0.23&decay=0.25&time=0.54&feedback=0.26">/audio?speed=0.75&scale=0&perm=0&slew=0.07&root=0.5&square=0&amount=0&offset=0.5&attack=0.23&decay=0.25&time=0.54&feedback=0.26</a></li>
<li><a href="/audio?speed=1&scale=36&perm=100&slew=0.42&root=1&square=1&amount=1&offset=0.98&attack=1&decay=1&time=0&feedback=0">/audio?speed=1&scale=36&perm=100&slew=0.42&root=1&square=1&amount=1&offset=0.98&attack=1&decay=1&time=0&feedback=0</a></li>
</body></html>`

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}

func audioHandler(w http.ResponseWriter, r *http.Request) {
	u := ui.New(noopHandler{})
	if err := u.Load(inFile); err != nil {
		log.Println(err)
		http.Error(w, "error processing audio", 500)
		return
	}

	u.Set("value19", floatValue(r, "speed"))
	u.Set("value70", floatValue(r, "slew"))
	u.Set("value77", floatValue(r, "root"))
	u.Set("value4", floatValue(r, "square"))
	u.Set("value39", floatValue(r, "amount"))
	u.Set("value46", floatValue(r, "offset"))
	u.Set("value13", floatValue(r, "attack"))
	u.Set("value12", floatValue(r, "decay"))
	u.Set("value29", floatValue(r, "time"))
	u.Set("value27", floatValue(r, "feedback"))

	scale, perm := intValue(r, "scale"), intValue(r, "perm")
	if 0 >= scale || scale >= len(scales) {
		scale = 0
	}
	notes := notesFromScale(scales[scale], 8, perm)
	for i, n := range notes {
		u.Set(seqNotes[i], float64(n)/120)
	}

	frames := int(length / time.Second * 44100 / 256)
	samp := u.Render(frames)
	normalize(samp)
	fadeout(samp, 44100)
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

func intValue(r *http.Request, name string) int {
	i, _ := strconv.Atoi(r.FormValue(name))
	return i
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

func fadeout(s []audio.Sample, samples int) {
	for i := range s {
		if r := len(s) - i; r < samples {
			s[i] *= audio.Sample(r) / audio.Sample(samples)
		}
	}
}

type noopHandler struct{}

func (noopHandler) Hello(kindInputs map[string][]string) {}
func (noopHandler) SetGraph(graph []*ui.Object)          {}
