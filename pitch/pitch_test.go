package pitch

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// table with proper values can be found here https://en.wikipedia.org/wiki/Scientific_pitch_notation
func TestDifferentPitches(t *testing.T) {
	v, _ := Of("A", "4", 440)
	assert.Equal(t, "440.00Hz", v)

	v2, _ := Of("C", "1", 440)
	assert.Equal(t, "32.70Hz", v2)

	v3, _ := Of("Gb", "2", 440)
	assert.Equal(t, "92.50Hz", v3)

	v4, _ := Of("D♯", "6", 440)
	assert.Equal(t, "1244.51Hz", v4)

	v5, _ := Of("B", "10", 440)
	assert.Equal(t, "31608.53Hz", v5)

	v6, _ := Of("C", "-1", 440)
	assert.Equal(t, "8.18Hz", v6)
}

func TestCustomTuning(t *testing.T) {
	v, _ := Of("A", "4", 442)
	assert.Equal(t, "442.00Hz", v)

	v2, _ := Of("A", "3", 442)
	assert.Equal(t, "221.00Hz", v2)

	v3, _ := Of("A", "5", 432)
	assert.Equal(t, "864.00Hz", v3)
}
