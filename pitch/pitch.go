package pitch

import (
	"math"
	"strconv"

	"github.com/go-music-theory/music-theory/note"
)

var A4Num = 58 // step no from C0

func OfNote(name string, tuning int) (string, error) {
	class := note.ClassNamed(name)
	octave := note.OctaveOf(name)
	return format(calcPitch(class, int(octave), tuning))
}

func OfClassAndOctave(class string, octaveStr string, tuning int) (string, error) {
	root, class := note.RootAndRemaining(class)

	octave, err := strconv.Atoi(octaveStr)
	if err != nil {
		return format(-1, err)
	}

	return format(calcPitch(root, octave, tuning))
}

func calcPitch(note note.Class, octave int, tuning int) (float64, error) {
	stepNo := int(note) + octave*12
	diffFromA4 := abs(A4Num - stepNo)
	magnitude := math.Pow(math.Pow(2, 1.0/12), float64(diffFromA4))

	if stepNo < A4Num {
		return round(float64(tuning) / magnitude), nil
	} else {
		return round(float64(tuning) * magnitude), nil
	}
}

func format(pitch float64, err error) (string, error) {
	if err == nil {
		return strconv.FormatFloat(pitch, 'f', 2, 64) + "Hz", nil
	} else {
		return "n/a", err
	}
}

func round(pitch float64) float64 {
	return math.Round(pitch*100) / 100
}

func abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}
