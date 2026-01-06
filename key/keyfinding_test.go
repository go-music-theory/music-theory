// Package key provides key-finding algorithms including the Krumhansl-Schmuckler algorithm
package key

import (
	"testing"

	"github.com/go-music-theory/music-theory/note"
	"gopkg.in/stretchr/testify.v1/assert"
)

// TestFindKey_CMajorScale tests finding C major from the C major scale
func TestFindKey_CMajorScale(t *testing.T) {
	notes := []note.Class{note.C, note.D, note.E, note.F, note.G, note.A, note.B}
	result := FindKey(notes)

	assert.Equal(t, note.C, result.Root, "Expected root to be C")
	assert.Equal(t, Major, result.Mode, "Expected mode to be Major")
}

// TestFindKey_AMinorScale tests finding A minor with emphasis on the tonic
func TestFindKey_AMinorScale(t *testing.T) {
	// A natural minor with emphasis on A (tonic)
	notes := []note.Class{note.A, note.A, note.A, note.B, note.C, note.D, note.E, note.F, note.G, note.A}
	result := FindKey(notes)

	assert.Equal(t, note.A, result.Root, "Expected root to be A")
	assert.Equal(t, Minor, result.Mode, "Expected mode to be Minor")
}

// TestFindKey_GMajorScale tests finding G major from the G major scale
func TestFindKey_GMajorScale(t *testing.T) {
	notes := []note.Class{note.G, note.A, note.B, note.C, note.D, note.E, note.Fs}
	result := FindKey(notes)

	assert.Equal(t, note.G, result.Root, "Expected root to be G")
	assert.Equal(t, Major, result.Mode, "Expected mode to be Major")
}

// TestFindKey_DMinorScale tests finding D minor with emphasis on the tonic
func TestFindKey_DMinorScale(t *testing.T) {
	// D natural minor with emphasis on D (tonic) and characteristic notes
	notes := []note.Class{note.D, note.D, note.D, note.E, note.F, note.G, note.A, note.As, note.C, note.D}
	result := FindKey(notes)

	assert.Equal(t, note.D, result.Root, "Expected root to be D")
	assert.Equal(t, Minor, result.Mode, "Expected mode to be Minor")
}

// TestFindKey_WithRepeatedNotes tests the algorithm with repeated notes (more realistic musical input)
func TestFindKey_WithRepeatedNotes(t *testing.T) {
	// C major with emphasis on tonic and dominant
	notes := []note.Class{
		note.C, note.C, note.C, note.E, note.G, note.G, note.C,
		note.D, note.E, note.F, note.G, note.A, note.B,
	}
	result := FindKey(notes)

	assert.Equal(t, note.C, result.Root, "Expected root to be C")
	assert.Equal(t, Major, result.Mode, "Expected mode to be Major")
}

// TestFindKey_EmptyInput tests that empty input returns a zero value key
func TestFindKey_EmptyInput(t *testing.T) {
	notes := []note.Class{}
	result := FindKey(notes)

	assert.Equal(t, note.Nil, result.Root, "Expected root to be Nil for empty input")
	assert.Equal(t, Nil, result.Mode, "Expected mode to be Nil for empty input")
}

// TestFindKey_FMajorScale tests finding F major from the F major scale
func TestFindKey_FMajorScale(t *testing.T) {
	notes := []note.Class{note.F, note.G, note.A, note.As, note.C, note.D, note.E}
	result := FindKey(notes)

	assert.Equal(t, note.F, result.Root, "Expected root to be F")
	assert.Equal(t, Major, result.Mode, "Expected mode to be Major")
}

// TestFindKey_EMinorScale tests finding E minor with emphasis on the tonic
func TestFindKey_EMinorScale(t *testing.T) {
	// E natural minor with emphasis on E (tonic)
	notes := []note.Class{note.E, note.E, note.E, note.Fs, note.G, note.A, note.B, note.C, note.D, note.E}
	result := FindKey(notes)

	assert.Equal(t, note.E, result.Root, "Expected root to be E")
	assert.Equal(t, Minor, result.Mode, "Expected mode to be Minor")
}

// TestCalculatePitchClassDistribution tests the pitch class distribution calculator
func TestCalculatePitchClassDistribution(t *testing.T) {
	notes := []note.Class{note.C, note.C, note.E, note.G}
	distribution := calculatePitchClassDistribution(notes)

	assert.Equal(t, 12, len(distribution), "Distribution should have 12 elements")
	assert.Equal(t, 2.0, distribution[0], "C should appear twice")
	assert.Equal(t, 1.0, distribution[4], "E should appear once")
	assert.Equal(t, 1.0, distribution[7], "G should appear once")
	assert.Equal(t, 0.0, distribution[1], "C# should not appear")
}

// TestRotateDistribution tests the distribution rotation function
func TestRotateDistribution(t *testing.T) {
	distribution := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// Rotate by 0 semitones (should be unchanged)
	rotated := rotateDistribution(distribution, 0)
	assert.Equal(t, distribution, rotated, "No rotation should leave distribution unchanged")

	// Rotate by 1 semitone
	rotated = rotateDistribution(distribution, 1)
	expected := []float64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1}
	assert.Equal(t, expected, rotated, "Rotation by 1 should shift elements")

	// Rotate by 12 semitones (full circle, should be same as original)
	rotated = rotateDistribution(distribution, 12)
	assert.Equal(t, distribution, rotated, "Rotation by 12 should return to original")
}

// TestCorrelate tests the Pearson correlation function
func TestCorrelate(t *testing.T) {
	// Perfect positive correlation
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	corr := correlate(x, y)
	assert.InDelta(t, 1.0, corr, 0.001, "Perfect positive correlation should be 1.0")

	// Perfect negative correlation
	x = []float64{1, 2, 3, 4, 5}
	y = []float64{10, 8, 6, 4, 2}
	corr = correlate(x, y)
	assert.InDelta(t, -1.0, corr, 0.001, "Perfect negative correlation should be -1.0")

	// No correlation
	x = []float64{1, 1, 1, 1, 1}
	y = []float64{2, 4, 6, 8, 10}
	corr = correlate(x, y)
	assert.InDelta(t, 0.0, corr, 0.001, "No correlation should be 0.0")

	// Invalid input (different lengths)
	x = []float64{1, 2, 3}
	y = []float64{1, 2}
	corr = correlate(x, y)
	assert.Equal(t, -2.0, corr, "Invalid input should return -2.0")
}

// TestClassToSemitone tests the note class to semitone conversion
func TestClassToSemitone(t *testing.T) {
	assert.Equal(t, 0, classToSemitone(note.C))
	assert.Equal(t, 1, classToSemitone(note.Cs))
	assert.Equal(t, 2, classToSemitone(note.D))
	assert.Equal(t, 3, classToSemitone(note.Ds))
	assert.Equal(t, 4, classToSemitone(note.E))
	assert.Equal(t, 5, classToSemitone(note.F))
	assert.Equal(t, 6, classToSemitone(note.Fs))
	assert.Equal(t, 7, classToSemitone(note.G))
	assert.Equal(t, 8, classToSemitone(note.Gs))
	assert.Equal(t, 9, classToSemitone(note.A))
	assert.Equal(t, 10, classToSemitone(note.As))
	assert.Equal(t, 11, classToSemitone(note.B))
	assert.Equal(t, 0, classToSemitone(note.Nil))
}

// TestSharpOrFlat tests the accidental symbol determination
func TestSharpOrFlat(t *testing.T) {
	// Natural notes should use sharp by default
	assert.Equal(t, note.Sharp, sharpOrFlat(note.C))
	assert.Equal(t, note.Sharp, sharpOrFlat(note.D))
	assert.Equal(t, note.Sharp, sharpOrFlat(note.E))
	assert.Equal(t, note.Sharp, sharpOrFlat(note.F))
	assert.Equal(t, note.Sharp, sharpOrFlat(note.G))
	assert.Equal(t, note.Sharp, sharpOrFlat(note.A))
	assert.Equal(t, note.Sharp, sharpOrFlat(note.B))

	// Accidentals should use flat notation
	assert.Equal(t, note.Flat, sharpOrFlat(note.Cs))
	assert.Equal(t, note.Flat, sharpOrFlat(note.Ds))
	assert.Equal(t, note.Flat, sharpOrFlat(note.Fs))
	assert.Equal(t, note.Flat, sharpOrFlat(note.Gs))
	assert.Equal(t, note.Flat, sharpOrFlat(note.As))
}
