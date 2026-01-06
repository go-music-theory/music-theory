// Chords have different Forms, such as Triad, Seventh, Extended, Added/Omitted, Specific or General.
package chord

import (
	//"log"
	"regexp"
	
	"github.com/go-music-theory/music-theory/note"
)

// Form is identified by positive/negative regular expressions, and then adds/removes pitch classes by interval from the root of the chord.
type Form struct {
	Name string
	pos  *regexp.Regexp
	add  FormAdd
	omit FormOmit
}

// FormAdd maps an interval-from-chord-root to a +/1 semitone adjustment
type FormAdd map[Interval]int

// FormOmit maps an interval-from-chord-root to omit
type FormOmit []Interval

// MatchString processes the positive/negative regular expressions to determine if this form matches a string.
func (this *Form) MatchString(s string) bool {
	return this.matchPosNegString(s)
}

//
// Private
//

// Regular expression to use mid-word, gluing together form expression parts
var nExp = "[. ]*"

// Regular expressions for different utilities
var (
	majorExp = "(M|maj|major)"
	minorExp = "([^a-z]|^)(m|min|minor)"

	flatExp  = "(f|flat|b|♭)"
	sharpExp = "(#|s|sharp)"
	halfExp  = "half"

	omitExp = "(omit|\\-)"

	dominantExp    = "(^|dom|dominant)"
	nondominantExp = "(non|nondom|nondominant)"
	diminishedExp  = "(dim|dimin|diminished)"
	augmentedExp   = "(aug|augment|augmented)"
	suspendedExp   = "(sus|susp|suspend|suspended)"
	harmonicExp    = "(harm|harmonic)"
)

// forms is an ordered set of rules to match, and corresponding chord intervals to setup.
var forms = []Form{

	// Root

	Form{
		Name: "Basic",
		add: FormAdd{
			I1: 0, // root
			I3: 4, // major 3rd
			I5: 7, // perfect 5th
		},
	},

	Form{
		Name: "Nondominant",
		pos:  exp(nondominantExp),
		omit: FormOmit{
			I1, // no root
		},
	},

	// Triads

	Form{
		Name: "Major Triad",
		pos:  exp("^" + majorExp + "([^a-z]|$)"),
		add: FormAdd{
			I3: 4, // major 3rd
			I5: 7, // perfect 5th
		},
	},

	Form{
		Name: "Minor Triad",
		pos:  exp("^" + minorExp + "([^a-z]|$)"),
		add: FormAdd{
			I3: 3, // minor 3rd
			I5: 7, // perfect 5th
		},
	},

	Form{
		Name: "Augmented Triad",
		pos:  exp("^" + augmentedExp),
		add: FormAdd{
			I3: 4, // major 3rd
			I5: 8, // augmented 5th
		},
	},

	Form{
		Name: "Diminished Triad",
		pos:  exp("^" + diminishedExp),
		add: FormAdd{
			I3: 3, // diminished (minor) 3rd
			I5: 6, // diminished 5th
		},
	},

	Form{
		Name: "Suspended Triad",
		pos:  exp("^" + suspendedExp),
		add: FormAdd{
			I4: 5, // 4th
			I5: 7, // perfect 5th
		},
		omit: FormOmit{
			I3,
		},
	},

	// Fifth

	Form{
		Name: "Omit Fifth",
		pos:  exp(omitExp + nExp + "5"),
		omit: FormOmit{
			I5, // no fifth
		},
	},

	Form{
		Name: "Flat Fifth",
		pos:  exp(flatExp + nExp + "5"),
		add: FormAdd{
			I5: 6, // flat 5th
		},
	},

	// Sixth

	Form{
		Name: "Add Sixth",
		pos:  exp("6"),
		add: FormAdd{
			I6: 9, // 6th
		},
	},

	Form{
		Name: "Augmented Sixth",
		pos:  exp(augmentedExp + nExp + "6"),
		add: FormAdd{
			I6: 10, // augmented 6th
		},
	},

	Form{
		Name: "Omit Sixth",
		pos:  exp(omitExp + nExp + "6"),
		omit: FormOmit{I6},
	},

	// Seventh

	Form{
		Name: "Add Seventh",
		pos:  exp("7"),
		add: FormAdd{
			I7: 10, // dominant 7th
		},
	},

	Form{
		Name: "Dominant Seventh",
		pos:  exp(dominantExp + nExp + "7"),
		add: FormAdd{
			I7: 10, // dominant 7th
		},
	},

	Form{
		Name: "Major Seventh",
		pos:  exp(majorExp + nExp + "7"),
		add: FormAdd{
			I7: 11, // major 7th
		},
	},

	Form{
		Name: "Minor Seventh",
		pos:  exp(minorExp + nExp + "7"),
		add: FormAdd{
			I7: 10, // minor 7th
		},
		omit: FormOmit{},
	},

	Form{
		Name: "Diminished Seventh",
		pos:  exp(diminishedExp + nExp + "7"),
		add: FormAdd{
			I7: 9, // diminished 7th
		},
	},

	Form{
		Name: "Half Diminished Seventh",
		pos:  exp(halfExp + nExp + diminishedExp + nExp + "7"),
		add: FormAdd{
			I3: 3,  // minor 3rd
			I5: 6,  // diminished 5th
			I7: 10, // minor 7th
		},
	},

	Form{
		Name: "Diminished Major Seventh",
		pos:  exp(diminishedExp + nExp + majorExp + nExp + "7"),
		add:  FormAdd{
		// TODO
		},
	},

	Form{
		Name: "Augmented Major Seventh",
		pos:  exp(augmentedExp + nExp + majorExp + nExp + "7"),
		add:  FormAdd{
		// TODO
		},
	},

	Form{
		Name: "Augmented Minor Seventh",
		pos:  exp(augmentedExp + nExp + minorExp + nExp + "7"),
		add:  FormAdd{
		// TODO
		},
		omit: FormOmit{},
	},

	Form{
		Name: "Harmonic Seventh",
		pos:  exp(harmonicExp + nExp + "7"),
		add: FormAdd{
			I3: 4,   // major 3rd
			I5: 7,   // perfect 5th
			I7: 969, // harmonic 7th (special value: 969 cents, not semitones)
		},
	},

	Form{
		Name: "Omit Seventh",
		pos:  exp(omitExp + nExp + "7"),
		omit: FormOmit{
			I7, // no 7th
		},
	},

	// Ninth

	Form{
		Name: "Add Ninth",
		pos:  exp("9"),
		add: FormAdd{
			I9: 14, // 9th
		},
	},

	Form{
		Name: "Dominant Ninth",
		pos:  exp(dominantExp + nExp + "9"),
		add: FormAdd{
			I7: 10, // minor 7th
			I9: 14, // dominant 9th
		},
	},

	Form{
		Name: "Major Ninth",
		pos:  exp(majorExp + nExp + "9"),
		add: FormAdd{
			I7: 11, // major 7th
			I9: 14, // dominant 9th
		},
	},

	Form{
		Name: "Minor Ninth",
		pos:  exp(minorExp + nExp + "9"),
		add: FormAdd{
			I7: 10, // minor 7th
			I9: 14, // dominant 9th
		},
	},

	Form{
		Name: "Sharp Ninth",
		pos:  exp(sharpExp + nExp + "9"),
		add: FormAdd{
			I9: 15, // sharp 9th
		},
	},

	Form{
		Name: "Omit Ninth",
		pos:  exp(omitExp + nExp + "9"),
		omit: FormOmit{
			I9, // no 9th
		},
	},

	// Eleventh

	Form{
		Name: "Add Eleventh",
		pos:  exp("11"),
		add: FormAdd{
			I11: 17, // 11th
		},
	},

	Form{
		Name: "Dominant Eleventh",
		pos:  exp(dominantExp + nExp + "11"),
		add: FormAdd{
			I7:  10, // minor 7th
			I9:  14, // dominant 9th
			I11: 17, // dominant 11th
		},
		omit: FormOmit{
			I3, // no 3rd
		},
	},

	Form{
		Name: "Major Eleventh",
		pos:  exp(majorExp + nExp + "11"),
		add: FormAdd{
			I7:  11, // major 7th
			I9:  14, // dominant 9th
			I11: 17, // dominant 11th
		},
	},

	Form{
		Name: "Minor Eleventh",
		pos:  exp(minorExp + nExp + "11"),
		add: FormAdd{
			I3:  3,  // minor 3rd
			I7:  10, // minor 7th
			I9:  14, // dominant 9th
			I11: 17, // dominant 11th
		},
	},

	Form{
		Name: "Omit Eleventh",
		pos:  exp(omitExp + nExp + "11"),
		omit: FormOmit{
			I11,
		},
	},

	// Thirteenth

	Form{
		Name: "Add Thirteenth",
		pos:  exp("13"),
		add: FormAdd{
			I13: 21, // dominant 13th
		},
	},

	Form{
		Name: "Dominant Thirteenth",
		pos:  exp(dominantExp + nExp + "13"),
		add: FormAdd{
			I7:  10, // minor 7th
			I9:  14, // dominant 9th
			I11: 17, // dominant 11th
			I13: 21, // dominant 13th
		},
		omit: FormOmit{
			I3, // no 3rd
		},
	},

	Form{
		Name: "Major Thirteenth",
		pos:  exp(majorExp + nExp + "13"),
		add: FormAdd{
			I3:  4,  // major 3rd
			I7:  11, // major 7th
			I9:  14, // dominant 9th
			I11: 17, // dominant 11th
			I13: 21, // dominant 13th
		},
	},

	Form{
		Name: "Minor Thirteenth",
		pos:  exp(minorExp + nExp + "13"),
		add: FormAdd{
			I3:  3,  // minor 3rd
			I7:  10, // minor 7th
			I9:  14, // dominant 9th
			I11: 17, // dominant 11th
			I13: 21, // dominant 13th
		},
	},

	// Lydian

	/*
		Form{
			Name: "Lydian",
			pos: exp("lyd"),
			// TODO
		},

		Form{
			Name: "Omit Lydian",
			pos: exp(omitExp+nExp+"lyd"),
			// TODO
		},
	*/

	// Specific

	/*	Form{
			Name: "AlphaSpecific",
			pos: exp("alpha"),
			// TODO
		},

		Form{
			Name: "BridgeSpecific",
			pos: exp("bridge"),
			// TODO
		},

		Form{
			Name: "ComplexeSonoreSpecific",
			pos: exp("(complexe|sonore)"),
			// TODO
		},

		Form{
			Name: "DreamSpecific",
			pos: exp("dream"),
			// TODO
		},

		Form{
			Name: "ElektraSpecific",
			pos: exp("elektra"),
			// TODO
		},

		Form{
			Name: "FarbenSpecific",
			pos: exp("farben"),
			// TODO
		},

		Form{
			Name: "GrandmotherSpecific",
			pos: exp("grandmother"),
			// TODO
		},

		Form{
			Name: "MagicSpecific",
			pos: exp("magic"),
			// TODO
		},

		Form{
			Name: "MµSpecific",
			pos: exp("µ"),
			// TODO
		},

		Form{
			Name: "MysticSpecific",
			pos: exp("mystic"),
			// TODO
		},

		Form{
			Name: "NorthernLightsSpecific",
			pos: exp("northern" + nExp + "light"),
			// TODO
		},

		Form{
			Name: "PetrushkaSpecific",
			pos: exp("petrush"),
			// TODO
		},

		Form{
			Name: "PsalmsSpecific",
			pos: exp("psalm"),
			// TODO
		},

		Form{
			Name: "SoWhatSpecific",
			pos: exp("so" + nExp + "what"),
			// TODO
		},

		Form{
			Name: "TristanSpecific",
			pos: exp("tristan"),
			// TODO
		},

		Form{
			Name: "VienneseTrichordSpecific",
			pos: exp("viennese" + nExp + "trichord"),
			// TODO
		},

		// General

		Form{
			Name: "MixedIntervalGeneral",
			pos: exp("mixed" + nExp + "interval"),
			// TODO
		},

		Form{
			Name: "SecundalGeneral",
			pos: exp("secundal"),
			// TODO
		},

		Form{
			Name: "TertianGeneral",
			pos: exp("tertian"),
			// TODO
		},

		Form{
			Name: "QuartalGeneral",
			pos: exp("quartal"),
			// TODO
		},

		Form{
			Name: "SyntheticChordGeneral",
			pos: exp("synthetic"),
			// TODO
		},*/

}

func exp(s string) *regexp.Regexp {
	r, _ := regexp.Compile(s)
	return r
}

func (this *Form) matchPosNegString(s string) bool {
	if this.pos == nil {
		return true
	} else if this.pos.MatchString(s) {
		//if this.neg != nil {
		//	if this.neg.MatchString(s) {
		//		return false
		//	}
		//}
		return true
	} else {
		return false
	}
}

// Build the chord by processing all Forms against the given name.
func (this *Chord) parseForms(name string) {
	var toDelete []Interval
	for _, f := range forms {
		if f.MatchString(name) {
			toDelete = append(toDelete, this.applyForm(f)...)
		}
	}
	for _, t := range toDelete {
		delete(this.Tones, t)
	}
	return
}

func (this *Chord) applyForm(f Form) (toDelete []Interval) {
	for i, c := range f.add {
		// Special handling for harmonic seventh (969 cents marker)
		if c == 969 {
			// Harmonic seventh is approximately 969 cents (~9.69 semitones)
			// We use a custom pitch class note.Bsb for this microtonal interval
			this.Tones[i] = calculateHarmonicSeventh(this.Root)
		} else {
			this.Tones[i], _ = this.Root.Step(c)
		}
	}
	for _, t := range f.omit {
		toDelete = append(toDelete, t)
	}
	return
}

// calculateHarmonicSeventh returns the harmonic seventh pitch class for any root
// The harmonic seventh is approximately 969 cents above the root (31 cents flat of minor 7th)
func calculateHarmonicSeventh(root note.Class) note.Class {
	// Map each root to its corresponding harmonic seventh pitch class
	switch root {
	case note.C:
		return note.Ch7
	case note.Cs:
		return note.Csh7
	case note.D:
		return note.Dh7
	case note.Ds:
		return note.Dsh7
	case note.E:
		return note.Eh7
	case note.F:
		return note.Fh7
	case note.Fs:
		return note.Fsh7
	case note.G:
		return note.Gh7
	case note.Gs:
		return note.Gsh7
	case note.A:
		return note.Ah7
	case note.As:
		return note.Ash7
	case note.B:
		return note.Bh7
	default:
		// Fallback to closest chromatic approximation (10 semitones)
		result, _ := root.Step(10)
		return result
	}
}
