// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

import (
	"fmt"
	"testing"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"gopkg.in/stretchr/testify.v1/assert"

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
