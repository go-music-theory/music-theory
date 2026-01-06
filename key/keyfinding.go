// Package key provides key-finding algorithms including the Krumhansl-Schmuckler algorithm
package key

import (
	"math"

	"github.com/go-music-theory/music-theory/note"
)

// majorProfile contains the Krumhansl-Kessler key profile weights for major keys.
// These empirically derived weights represent the stability of each pitch class in a major key context.
// Index represents semitones from the tonic: [C, C#, D, D#, E, F, F#, G, G#, A, A#, B]
var majorProfile = []float64{6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39, 3.66, 2.29, 2.88}

// minorProfile contains the Krumhansl-Kessler key profile weights for minor keys.
// These empirically derived weights represent the stability of each pitch class in a minor key context.
// Index represents semitones from the tonic: [C, C#, D, D#, E, F, F#, G, G#, A, A#, B]
var minorProfile = []float64{6.33, 2.68, 3.52, 5.38, 2.60, 3.53, 2.54, 4.75, 3.98, 2.69, 3.34, 3.17}

// FindKey implements the Krumhansl-Schmuckler key-finding algorithm.
// It takes a slice of note.Class values and returns the most likely key.
// The algorithm:
// 1. Calculates the pitch class distribution from the input notes
// 2. Correlates this distribution with all 24 major and minor key profiles
// 3. Returns the key with the highest correlation coefficient
func FindKey(notes []note.Class) Key {
	if len(notes) == 0 {
		return Key{}
	}

	// Calculate pitch class distribution
	distribution := calculatePitchClassDistribution(notes)

	// Find the best matching key
	var bestKey Key
	var bestCorrelation = -2.0 // Correlation coefficient ranges from -1 to 1

	// Check all 12 major keys
	allNotes := []note.Class{note.C, note.Cs, note.D, note.Ds, note.E, note.F, note.Fs, note.G, note.Gs, note.A, note.As, note.B}
	for _, root := range allNotes {
		rotatedDistribution := rotateDistribution(distribution, classToSemitone(root))
		correlation := correlate(rotatedDistribution, majorProfile)
		if correlation > bestCorrelation {
			bestCorrelation = correlation
			bestKey = Key{
				Root:      root,
				Mode:      Major,
				AdjSymbol: sharpOrFlat(root),
			}
		}
	}

	// Check all 12 minor keys
	for _, root := range allNotes {
		rotatedDistribution := rotateDistribution(distribution, classToSemitone(root))
		correlation := correlate(rotatedDistribution, minorProfile)
		if correlation > bestCorrelation {
			bestCorrelation = correlation
			bestKey = Key{
				Root:      root,
				Mode:      Minor,
				AdjSymbol: sharpOrFlat(root),
			}
		}
	}

	return bestKey
}

// calculatePitchClassDistribution creates a histogram of pitch classes.
// Returns a slice of 12 floats representing the count of each pitch class (C through B).
func calculatePitchClassDistribution(notes []note.Class) []float64 {
	distribution := make([]float64, 12)

	for _, n := range notes {
		semitone := classToSemitone(n)
		if semitone >= 0 && semitone < 12 {
			distribution[semitone]++
		}
	}

	return distribution
}

// rotateDistribution rotates a pitch class distribution by the specified number of semitones.
// This allows us to transpose the distribution so a different pitch class becomes the tonic (index 0).
func rotateDistribution(distribution []float64, semitones int) []float64 {
	rotated := make([]float64, 12)
	for i := 0; i < 12; i++ {
		rotated[i] = distribution[(i+semitones)%12]
	}
	return rotated
}

// correlate calculates the Pearson correlation coefficient between two distributions.
// This measures how well the observed pitch class distribution matches a key profile.
func correlate(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return -2.0 // Invalid correlation
	}

	// Calculate means
	var sumX, sumY float64
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
	}
	meanX := sumX / float64(len(x))
	meanY := sumY / float64(len(y))

	// Calculate correlation coefficient
	var numerator, denomX, denomY float64
	for i := 0; i < len(x); i++ {
		diffX := x[i] - meanX
		diffY := y[i] - meanY
		numerator += diffX * diffY
		denomX += diffX * diffX
		denomY += diffY * diffY
	}

	if denomX == 0 || denomY == 0 {
		return 0.0
	}

	return numerator / math.Sqrt(denomX*denomY)
}

// classToSemitone converts a note.Class to its semitone value (0-11).
// This is a helper function for the key-finding algorithm.
func classToSemitone(c note.Class) int {
	switch c {
	case note.C:
		return 0
	case note.Cs:
		return 1
	case note.D:
		return 2
	case note.Ds:
		return 3
	case note.E:
		return 4
	case note.F:
		return 5
	case note.Fs:
		return 6
	case note.G:
		return 7
	case note.Gs:
		return 8
	case note.A:
		return 9
	case note.As:
		return 10
	case note.B:
		return 11
	default:
		return 0
	}
}

// sharpOrFlat determines the appropriate accidental symbol for a given root note.
// Returns Sharp for notes typically written with sharps, Flat for notes typically written with flats.
func sharpOrFlat(root note.Class) note.AdjSymbol {
	switch root {
	case note.Cs, note.Ds, note.Fs, note.Gs, note.As:
		// These pitch classes are black keys and can be represented as either sharps or flats
		// We use flat notation for the common flat keys (Db, Eb, Gb, Ab, Bb)
		return note.Flat
	}
	return note.Sharp
}
