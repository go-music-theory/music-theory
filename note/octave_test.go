// A perfect octave is the interval between one musical pitch and another with half or double its frequency.
package note

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestOctaveOf(t *testing.T) {
	n := OctaveOf("A4")
	assert.Equal(t, Octave(4), n)
}

func TestToInt(t *testing.T) {
	n := OctaveOf("A4")
	assert.Equal(t, 4, int(n))
}
