package note

import (
	"math"
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestNotePitch(t *testing.T) {
	tests := []struct {
		name     string
		note     *Note
		tuning   Tuning
		expected float64
	}{
		// Standard tuning (440Hz) tests
		{
			name:     "A4 standard tuning",
			note:     &Note{Class: A, Octave: 4},
			tuning:   TuningStandard,
			expected: 440.0,
		},
		{
			name:     "C4 standard tuning (middle C)",
			note:     &Note{Class: C, Octave: 4},
			tuning:   TuningStandard,
			expected: 261.6255653005986,
		},
		{
			name:     "A5 standard tuning",
			note:     &Note{Class: A, Octave: 5},
			tuning:   TuningStandard,
			expected: 880.0,
		},
		{
			name:     "A3 standard tuning",
			note:     &Note{Class: A, Octave: 3},
			tuning:   TuningStandard,
			expected: 220.0,
		},
		// Verdi tuning (432Hz) tests
		{
			name:     "A4 Verdi tuning",
			note:     &Note{Class: A, Octave: 4},
			tuning:   TuningVerdi,
			expected: 432.0,
		},
		{
			name:     "A5 Verdi tuning",
			note:     &Note{Class: A, Octave: 5},
			tuning:   TuningVerdi,
			expected: 864.0,
		},
		{
			name:     "C4 Verdi tuning",
			note:     &Note{Class: C, Octave: 4},
			tuning:   TuningVerdi,
			expected: 256.8687368010849,
		},
		// Test with zero tuning (should default to 440)
		{
			name:     "A4 with zero tuning defaults to 440",
			note:     &Note{Class: A, Octave: 4},
			tuning:   0,
			expected: 440.0,
		},
		// Test other notes
		{
			name:     "E4 standard tuning",
			note:     &Note{Class: E, Octave: 4},
			tuning:   TuningStandard,
			expected: 329.6275569128699,
		},
		{
			name:     "G4 standard tuning",
			note:     &Note{Class: G, Octave: 4},
			tuning:   TuningStandard,
			expected: 391.99543598174927,
		},
		{
			name:     "B4 standard tuning",
			note:     &Note{Class: B, Octave: 4},
			tuning:   TuningStandard,
			expected: 493.8833012561241,
		},
		// Test sharp notes
		{
			name:     "C#4 standard tuning",
			note:     &Note{Class: Cs, Octave: 4},
			tuning:   TuningStandard,
			expected: 277.1826309768721,
		},
		{
			name:     "F#4 standard tuning",
			note:     &Note{Class: Fs, Octave: 4},
			tuning:   TuningStandard,
			expected: 369.9944227116344,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.note.Pitch(tt.tuning)
			// Use a small epsilon for floating point comparison
			epsilon := 0.0000001
			if math.Abs(actual-tt.expected) > epsilon {
				t.Errorf("Pitch() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestNoteMIDI(t *testing.T) {
	tests := []struct {
		name     string
		note     *Note
		expected int
	}{
		{
			name:     "A4 MIDI number",
			note:     &Note{Class: A, Octave: 4},
			expected: 69,
		},
		{
			name:     "C4 MIDI number (middle C)",
			note:     &Note{Class: C, Octave: 4},
			expected: 60,
		},
		{
			name:     "C-1 MIDI number",
			note:     &Note{Class: C, Octave: -1},
			expected: 0,
		},
		{
			name:     "A0 MIDI number",
			note:     &Note{Class: A, Octave: 0},
			expected: 21,
		},
		{
			name:     "C5 MIDI number",
			note:     &Note{Class: C, Octave: 5},
			expected: 72,
		},
		{
			name:     "G9 MIDI number",
			note:     &Note{Class: G, Octave: 9},
			expected: 127,
		},
		{
			name:     "C#4 MIDI number",
			note:     &Note{Class: Cs, Octave: 4},
			expected: 61,
		},
		{
			name:     "D4 MIDI number",
			note:     &Note{Class: D, Octave: 4},
			expected: 62,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.note.MIDI()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNamedWithPitch(t *testing.T) {
	// Test that Named notes can calculate pitch
	n := Named("A4")
	pitch := n.Pitch(TuningStandard)
	assert.Equal(t, 440.0, pitch)

	n2 := Named("C4")
	pitch2 := n2.Pitch(TuningStandard)
	epsilon := 0.0000001
	if math.Abs(pitch2-261.6255653005986) > epsilon {
		t.Errorf("Named('C4').Pitch() = %v, want 261.6255653005986", pitch2)
	}
}
