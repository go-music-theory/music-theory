// Chords have different Functions, such as Diatonic, Altered or Other.
package chord

type Function int

const (
	GenericFunction Function = iota

	// Diatonic
	TonicDiatonic
	DominantDiatonic
	SubdominantDiatonic
	SupertonicDiatonic
	MediantDiatonic
	SubmediantDiatonic
	LeadingDiatonic
	SubtonicDiatonic

	// Altered
	ApproachAltered
	BorrowedAltered
	ChromaticMediantAltered
	NeapolitanAltered
	PassingAltered
	SecondaryAltered
	SecondaryDominantAltered
	SecondaryLeadingToneAltered
	SecondarySupertonicAltered

	// Other
	CommonOther
	ContrastOther
	PrimaryTriadOther
	SubsidiaryOther
)
