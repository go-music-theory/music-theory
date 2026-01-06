// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

import (
	"fmt"
	"io/ioutil"
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
	"gopkg.in/yaml.v2"

	"github.com/go-music-theory/music-theory/key"
	"github.com/go-music-theory/music-theory/note"
)

func TestChordExpectations(t *testing.T) {
	testExpectations := testExpectationManifest{}
	file, err := ioutil.ReadFile("testdata/expectations.yaml")
	assert.Nil(t, err)

	err = yaml.Unmarshal(file, &testExpectations)
	assert.Nil(t, err)

	assert.True(t, len(testExpectations.Chords) > 0)
	for name, expect := range testExpectations.Chords {
		actual := Of(name)
		assert.Equal(t, expect.Root, actual.Root.String(actual.AdjSymbol), fmt.Sprintf("name:%v expect.Root:%v actual.Root:%v", name, expect.Root, actual.Root.String(actual.AdjSymbol)))
		for i, c := range expect.Tones {
			assert.Equal(t, c, actual.Tones[i].String(actual.AdjSymbol), fmt.Sprintf("name:%v expect.Tones[%v]:%v actual.Tones[%v]:%v", name, i, c, i, actual.Tones[i].String(actual.AdjSymbol)))
		}
		for i, c := range actual.Tones {
			assert.Equal(t, expect.Tones[i], c.String(actual.AdjSymbol), fmt.Sprintf("name:%v actual.Tones[%v]:%v expect.Tones[%v]:%v", name, i, c.String(actual.AdjSymbol), i, expect.Tones[i]))
		}
	}
}

func TestNotes(t *testing.T) {
	c := Of("Cm nondominant -5 +6 +7 +9")
	assert.Equal(t, []*note.Note{
		&note.Note{Class: note.Ds},
		&note.Note{Class: note.A},
		&note.Note{Class: note.As},
		&note.Note{Class: note.D},
	}, c.Notes())
}

func TestOf_Invalid(t *testing.T) {
	k := key.Of("P-funk")
	assert.Equal(t, note.Nil, k.Root)
}

func TestTranspose(t *testing.T) {
	actualChord := Chord{
		Root:      note.C,
		AdjSymbol: note.Flat,
		Tones: map[Interval]note.Class{
			I3: note.Ds,
			I6: note.A,
			I7: note.As,
			I9: note.D,
		},
	}
	expectChord := Chord{
		Root:      note.Ds,
		AdjSymbol: note.Flat,
		Tones: map[Interval]note.Class{
			I3: note.Fs,
			I6: note.C,
			I7: note.Cs,
			I9: note.F,
		},
	}
	assert.Equal(t, expectChord, actualChord.Transpose(3))
}

func TestAlternativeNotation(t *testing.T) {
	// Test Δ for major
	c1 := Of("CΔ7")
	assert.Equal(t, note.C, c1.Root)
	assert.Equal(t, note.E, c1.Tones[I3]) // major 3rd
	assert.Equal(t, note.B, c1.Tones[I7]) // major 7th

	// Test − for minor
	c2 := Of("C−")
	assert.Equal(t, note.C, c2.Root)
	assert.Equal(t, note.Ds, c2.Tones[I3]) // minor 3rd

	// Test + for augmented
	c3 := Of("C+")
	assert.Equal(t, note.C, c3.Root)
	assert.Equal(t, note.Gs, c3.Tones[I5]) // augmented 5th

	// Test ° for diminished
	c4 := Of("C°")
	assert.Equal(t, note.C, c4.Root)
	assert.Equal(t, note.Ds, c4.Tones[I3]) // minor 3rd
	assert.Equal(t, note.Fs, c4.Tones[I5]) // diminished 5th

	// Test ø for half diminished
	c5 := Of("Cø7")
	assert.Equal(t, note.C, c5.Root)
	assert.Equal(t, note.Ds, c5.Tones[I3]) // minor 3rd
	assert.Equal(t, note.Fs, c5.Tones[I5]) // diminished 5th
	assert.Equal(t, note.As, c5.Tones[I7]) // minor 7th
}

func TestPowerChord(t *testing.T) {
	c := Of("C5")
	assert.Equal(t, note.C, c.Root)
	assert.Equal(t, note.G, c.Tones[I5]) // perfect 5th
	_, hasThird := c.Tones[I3]
	assert.False(t, hasThird) // no third in power chord
}

func TestAlteredDominant(t *testing.T) {
	// Test with "7alt" notation
	c := Of("C7alt")
	assert.Equal(t, note.C, c.Root)
	assert.Equal(t, note.E, c.Tones[I3])   // major 3rd
	assert.Equal(t, note.Fs, c.Tones[I5])  // flat 5th (diminished 5th)
	assert.Equal(t, note.Gs, c.Tones[I6])  // sharp 5th (augmented 5th)
	assert.Equal(t, note.As, c.Tones[I7])  // dominant 7th
	assert.Equal(t, note.Cs, c.Tones[I9])  // flat 9th
	assert.Equal(t, note.Ds, c.Tones[I10]) // sharp 9th

	// Test with just "alt" notation
	c2 := Of("Calt")
	assert.Equal(t, note.C, c2.Root)
	assert.Equal(t, note.E, c2.Tones[I3])   // major 3rd
	assert.Equal(t, note.Fs, c2.Tones[I5])  // flat 5th (diminished 5th)
	assert.Equal(t, note.Gs, c2.Tones[I6])  // sharp 5th (augmented 5th)
	assert.Equal(t, note.As, c2.Tones[I7])  // dominant 7th
	assert.Equal(t, note.Cs, c2.Tones[I9])  // flat 9th
	assert.Equal(t, note.Ds, c2.Tones[I10]) // sharp 9th
}

func TestLydianChord(t *testing.T) {
	c := Of("Clyd")
	assert.Equal(t, note.C, c.Root)
	assert.Equal(t, note.E, c.Tones[I3])  // major 3rd
	assert.Equal(t, note.Fs, c.Tones[I4]) // augmented 4th (#11)
	assert.Equal(t, note.G, c.Tones[I5])  // perfect 5th
	assert.Equal(t, note.B, c.Tones[I7])  // major 7th
}

func TestSpecificChords(t *testing.T) {
	// Test Mystic chord
	c1 := Of("Cmystic")
	assert.Equal(t, note.C, c1.Root)
	assert.Equal(t, note.D, c1.Tones[I2])  // major 2nd
	assert.Equal(t, note.Fs, c1.Tones[I4]) // augmented 4th
	assert.Equal(t, note.A, c1.Tones[I6])  // major 6th
	assert.Equal(t, note.B, c1.Tones[I7])  // major 7th

	// Test Tristan chord
	c2 := Of("Ctristan")
	assert.Equal(t, note.C, c2.Root)
	assert.Equal(t, note.Fs, c2.Tones[I4]) // augmented 4th
	assert.Equal(t, note.Gs, c2.Tones[I6]) // augmented 5th
	assert.Equal(t, note.As, c2.Tones[I7]) // minor 7th
}

//
// Private
//

type testKey struct {
	Root  string
	Tones map[Interval]string
}

type testExpectationManifest struct {
	Chords map[string]testKey
}
