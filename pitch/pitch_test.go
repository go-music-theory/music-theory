package pitch

import (
	"testing"
  
	"gopkg.in/stretchr/testify.v1/assert"
)

// table with proper values can be found here https://en.wikipedia.org/wiki/Scientific_pitch_notation

func TestOfClassAndOctave(t *testing.T) {
	assertPitchOfClassAndOctave(t, "440.00Hz", "A", "4", 440)
	assertPitchOfClassAndOctave(t, "32.70Hz", "C", "1", 440)
	assertPitchOfClassAndOctave(t, "92.50Hz", "Gb", "2", 440)
	assertPitchOfClassAndOctave(t, "1244.51Hz", "D♯", "6", 440)
	assertPitchOfClassAndOctave(t, "31608.53Hz", "B", "10", 440)
	assertPitchOfClassAndOctave(t, "8.18Hz", "C", "-1", 440)
  assertPitchOfClassAndOctave(t, "442.00Hz", "A", "4", 442)
	assertPitchOfClassAndOctave(t, "221.00Hz", "A", "3", 442)
	assertPitchOfClassAndOctave(t, "864.00Hz", "A", "5", 432)
}

func TestPitchOfNote(t *testing.T) {
	assertPitchOfNote(t, "440.00Hz", "A4", 440)
	assertPitchOfNote(t, "32.70Hz", "C1", 440)
	assertPitchOfNote(t, "92.50Hz", "Gb2", 440)
	assertPitchOfNote(t, "1244.51Hz", "D#6", 440)
	assertPitchOfNote(t, "31608.53Hz", "B10", 440)
	assertPitchOfNote(t, "8.18Hz", "C-1", 440)
  assertPitchOfNote(t, "442.00Hz", "A4", 442)
	assertPitchOfNote(t, "221.00Hz", "A3", 442)
	assertPitchOfNote(t, "864.00Hz", "A5", 432)
}

func assertPitchOfClassAndOctave(t *testing.T, expected string, class string, octave string, tuning int) {
	actual, err := OfClassAndOctave(class, octave, tuning)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func assertPitchOfNote(t *testing.T, expected string, name string, tuning int) {
	actual, err := OfNote(name, tuning)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
