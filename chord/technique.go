// Chords have different Techniques, such as Block, Chordioid, Guitar, Open, Power or Slash.
package chord

type Technique int

const (
	GenericTechnique Technique = iota

	// Techniques
	BlockTechnique
	ChordioidTechnique
	GuitarTechnique
	OpenTechnique
	PowerTechnique
	SlashTechnique
)
