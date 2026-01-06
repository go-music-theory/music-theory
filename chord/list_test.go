// It's possible to export a list of all known chord parsing rules.
package chord

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestListToYAML(t *testing.T) {
	c := ChordFormList
	out := c.ToYAML()
	assert.Equal(t, "- Basic\n- Nondominant\n- Major Triad\n- Minor Triad\n- Augmented Triad\n- Diminished Triad\n- Suspended Triad\n- Power Chord\n- Omit Fifth\n- Flat Fifth\n- Add Sixth\n- Augmented Sixth\n- Omit Sixth\n- Add Seventh\n- Dominant Seventh\n- Altered Dominant Seventh\n- Altered Dominant\n- Major Seventh\n- Minor Seventh\n- Diminished Seventh\n- Half Diminished Seventh\n- Diminished Major Seventh\n- Augmented Major Seventh\n- Augmented Minor Seventh\n- Harmonic Seventh\n- Omit Seventh\n- Add Ninth\n- Dominant Ninth\n- Major Ninth\n- Minor Ninth\n- Sharp Ninth\n- Omit Ninth\n- Add Eleventh\n- Dominant Eleventh\n- Major Eleventh\n- Minor Eleventh\n- Omit Eleventh\n- Add Thirteenth\n- Dominant Thirteenth\n- Major Thirteenth\n- Minor Thirteenth\n- Lydian\n- Omit Lydian\n- AlphaSpecific\n- BridgeSpecific\n- ComplexeSonoreSpecific\n- DreamSpecific\n- ElektraSpecific\n- FarbenSpecific\n- GrandmotherSpecific\n- MagicSpecific\n- MµSpecific\n- MysticSpecific\n- NorthernLightsSpecific\n- PetrushkaSpecific\n- PsalmsSpecific\n- SoWhatSpecific\n- TristanSpecific\n- VienneseTrichordSpecific\n- MixedIntervalGeneral\n- SecundalGeneral\n- TertianGeneral\n- QuartalGeneral\n- SyntheticChordGeneral\n", out)
}
