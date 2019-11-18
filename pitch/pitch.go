package pitch

import (
	"github.com/go-music-theory/music-theory/note"
	"strconv"
	"math"
)

var A_4_NO = 58 // step no from C0

func Of(name string, octave string, tuning int) (string, error) {
	root, name := note.RootAndRemaining(name)
	octaveNo, err := calcOctave(octave)
	if err != nil {
		return "", err
	}

	pitch := calcPitch(root, octaveNo, tuning)
	return strconv.FormatFloat(pitch, 'f', 2, 64) + "Hz", nil
}

func calcPitch(note note.Class, octave int, tuning int) float64 {
	stepNo := int(note) + octave * 12
	diffFromA4 := abs(A_4_NO - stepNo)
	magnitude := math.Pow(math.Pow(2, 1.0 / 12), float64(diffFromA4))

	if stepNo < A_4_NO {
		return round(float64(tuning) / magnitude)
	} else {
		return round(float64(tuning) * magnitude)
	}
}


func round(pitch float64) float64 {
	return math.Round(pitch*100) / 100
}

func calcOctave(octave string) (int, error) {
	return strconv.Atoi(octave)
}

func abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

