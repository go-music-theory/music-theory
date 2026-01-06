// A scale, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package scale

import (
	"testing"

	"github.com/go-music-theory/music-theory/note"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestListAllModes(t *testing.T) {
	assert.Equal(t, len(modes), len(ScaleModeList))
}

func TestModeMatchString(t *testing.T) {
	// TODO
}

func TestScaleParseModes(t *testing.T) {
	c := Of("CM")
	assert.Equal(t, map[Interval]note.Class{
		I1: note.C,
		I2: note.D,
		I3: note.E,
		I4: note.F,
		I5: note.G,
		I6: note.A,
		I7: note.B,
	}, c.Tones)
}

//
// Private
//

//func assertEquivalentModes(t *testing.T, expectModes []Mode, actualModes []Mode) {
//	for _, expectMode := range expectModes {
//		assert.Contains(t, actualModes, expectMode,
//			fmt.Sprintf("expect %v in actual %v", expectModes, actualModes))
//	}
//	for _, actualMode := range actualModes {
//		assert.Contains(t, expectModes, actualMode,
//			fmt.Sprintf("actual %v in expected %v", actualModes, expectModes))
//	}
//}
