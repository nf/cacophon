package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/nf/sigourney/audio"
	"github.com/nf/sigourney/ui"
)

var (
	httpAddr = flag.String("http", "localhost:8080", "HTTP listen address")
	length   = flag.Duration("len", 5*time.Second, "recording length")
	inFile   = flag.String("in", "cacophon", "patch input file name")
)

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/audio", audioHandler)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

func audioHandler(w http.ResponseWriter, r *http.Request) {
	u := ui.New(noopHandler{})
	if err := u.Load(*inFile); err != nil {
		log.Println(err)
		http.Error(w, "error processing audio", 500)
		return
	}

	u.Set("value32", floatValue(r, "a")) // delay
	u.Set("value22", floatValue(r, "b")) // speed
	u.Set("value17", floatValue(r, "c")) // fm

	frames := int(*length / time.Second * 44100 / 256)
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
