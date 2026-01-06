// Package note provides musical note models and utilities
package note

import (
	"math"
)

// Tuning represents the reference frequency for A4 (default is 440Hz)
type Tuning float64

const (
	// TuningStandard represents the standard concert pitch (A4 = 440Hz)
	TuningStandard Tuning = 440.0
	// TuningVerdi represents Verdi's tuning (A4 = 432Hz)
	TuningVerdi Tuning = 432.0
)

// Pitch returns the frequency in Hz for this note based on the given tuning.
// Uses the standard pitch formula: f(n) = tuning * 2^((n-69)/12)
// where n is the MIDI note number and 69 is the MIDI number for A4.
func (n *Note) Pitch(tuning Tuning) float64 {
	if tuning == 0 {
		tuning = TuningStandard
	}
	midiNote := n.MIDI()
	return float64(tuning) * math.Pow(2.0, float64(midiNote-69)/12.0)
}

// MIDI returns the MIDI note number for this note.
// MIDI note numbers: C-1 = 0, A4 = 69, C4 (middle C) = 60
func (n *Note) MIDI() int {
	// Get semitone offset from C
	semitoneOffset := classToSemitone(n.Class)
	// Calculate MIDI note: (octave + 1) * 12 + semitoneOffset
	// We add 1 to octave because MIDI octave -1 starts at 0
	return (int(n.Octave)+1)*12 + semitoneOffset
}

// classToSemitone returns the semitone offset from C for a given pitch class
func classToSemitone(c Class) int {
	switch c {
	case C:
		return 0
	case Cs:
		return 1
	case D:
		return 2
	case Ds:
		return 3
	case E:
		return 4
	case F:
		return 5
	case Fs:
		return 6
	case G:
		return 7
	case Gs:
		return 8
	case A:
		return 9
	case As:
		return 10
	case B:
		return 11
	default:
		return 0
	}
}
