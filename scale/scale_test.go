// In music theory, a scale is any set of musical notes ordered by fundamental frequency or pitch. A scale ordered by increasing pitch is an ascending scale, and a scale ordered by decreasing pitch is a descending scale. Some scales contain different pitches when ascending than when descending. For example, the Melodic minor scale.
package scale

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"

	"fmt"
	"github.com/go-music-theory/music-theory/key"
	"github.com/go-music-theory/music-theory/note"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func TestScaleExpectations(t *testing.T) {
	testExpectations := testExpectationManifest{}
	file, err := ioutil.ReadFile("testdata/expectations.yaml")
	assert.Nil(t, err)

	err = yaml.Unmarshal(file, &testExpectations)
	assert.Nil(t, err)

	assert.True(t, len(testExpectations.Scales) > 0)
	for name, expect := range testExpectations.Scales {
		actual := Of(name)
		assert.Equal(t, expect.Root, actual.Root.String(actual.AdjSymbol), fmt.Sprintf("name:%v expect.Root:%v actual.Root:%v actual.AdjSymbol:%v", name, expect.Root, actual.Root.String(actual.AdjSymbol), actual.AdjSymbol))
		for i, c := range expect.Tones {
			assert.Equal(t, c, actual.Tones[i].String(actual.AdjSymbol), fmt.Sprintf("name:%v expect.Tones[%v]:%v actual.Tones[%v]:%v actual.AdjSymbol:%v", name, i, c, i, actual.Tones[i].String(actual.AdjSymbol), actual.AdjSymbol))
		}
		for i, c := range actual.Tones {
			assert.Equal(t, expect.Tones[i], c.String(actual.AdjSymbol), fmt.Sprintf("name:%v actual.Tones[%v]:%v expect.Tones[%v]:%v actual.AdjSymbol:%v", name, i, c.String(actual.AdjSymbol), i, expect.Tones[i], actual.AdjSymbol))
		}
	}
}

func TestNotes(t *testing.T) {
	c := Of("C natural minor")
	assert.Equal(t, []*note.Note{
		&note.Note{Class: note.C},
		&note.Note{Class: note.D},
		&note.Note{Class: note.Ds},
		&note.Note{Class: note.F},
		&note.Note{Class: note.G},
		&note.Note{Class: note.Gs},
		&note.Note{Class: note.As},
	}, c.Notes())
}

func TestOf_Invalid(t *testing.T) {
	k := key.Of("P-funk")
	assert.Equal(t, note.Nil, k.Root)
}

//
// Private
//

type testKey struct {
	Root  string
	Tones map[Interval]string
}

type testExpectationManifest struct {
	Scales map[string]testKey
}
