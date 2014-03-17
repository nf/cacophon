package main

import (
	"regexp"
	"strconv"
	"strings"
)

var scales [][]int

func init() {
	for _, line := range strings.Split(strings.TrimSpace(scaleText), "\n") {
		scales = append(scales, parseScale(line))
	}
}

func parseScale(line string) []int {
	var s []int
	m := scaleRE.FindAllString(line, -1)
	for _, n := range m {
		n = n[1:] // trim space
		i := 0
		if len(n) == 2 {
			switch n[0] {
			case '#':
				i++
			case 'b':
				i--
			}
			n = n[1:]
		}
		i2, _ := strconv.Atoi(n)
		i += (i2 - 1) * 2
		s = append(s, i)
	}
	return s
}

var scaleRE = regexp.MustCompile(` [#b]?[0-9]`)

const scaleText = `
Acoustic scale 1 2 3 #4 5 6 b7 
natural minor scale 1 2 b3 4 5 b6 b7 
Altered scale 1 b2 b3 b4 b5 b6 b7 
Augmented scale 1 b3 3 5 #5 7 
Bebop dominant scale  1 2 3 4 5 6 b7 7 
Blues scale 1 b3 4 #4 5 b7 
Dorian mode 1 2 b3 4 5 6 b7 
Double harmonic scale 1 b2 3 4 5 b6 7 
Enigmatic scale 1 b2 3 #4 #5 #6 7 
Flamenco mode 1 b2 3 4 5 b6 7 
Half diminished scale 1 2 b3 4 b5 b6 b7 
Harmonic major scale 1 2 3 4 5 b6 7 
Harmonic minor scale 1 2 b3 4 5 b6 7 
Hirajoshi scale 1 2 b3 5 b6 
Hungarian minor scale 1 2 b3 #4 5 b6 7 
Insen scale 1 b2 4 5 b7 
major scale 1 2 3 4 5 6 7 
Istrian scale 1 b2 b3 b4 b5 5 
Iwato scale 1 b2 4 b5 b7 
Locrian mode 1 b2 b3 4 b5 b6 b7 
Lydian augmented scale 1 2 3 #4 #5 6 7 
Lydian mode 1 2 3 #4 5 6 7 
bebop scale 1 2 3 4 5 #5 6 7 
Major Locrian scale 1 2 3 4 b5 b6 b7 
pentatonic scale 1 2 3 5 6 
Melodic minor scale 1 2 b3 4 5 6 7 
Minor pentatonic scale 1 b3 4 5 b7 
Adonai malakh mode 1 2 3 4 5 6 b7 
Neapolitan major scale 1 b2 b3 4 5 6 7 
Neapolitan minor scale 1 b2 b3 4 5 b6 7 
Persian scale 1 b2 3 4 b5 b6 7 
Phrygian dominant scale 1 b2 3 4 5 b6 b7 
Phrygian mode 1 b2 b3 4 5 b6 b7 
Prometheus scale 1 2 3 #4 6 b7 
Tritone scale 1 b2 3 b5 5 b7 
Ukrainian Dorian scale 1 2 b3 #4 5 6 b7 
Whole tone scale 1 2 3 #4 #5 #6 
`
