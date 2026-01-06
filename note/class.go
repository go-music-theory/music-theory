// In music, a pitch class is a set of all pitches that are a whole number of octaves apart, e.g., the pitch class C consists of the Cs in all octaves.
package note

// Class of pitch for a note (across all octaves)
type Class int

const (
	Nil Class = iota
	C
	Cs
	D
	Ds
	E
	F
	Fs
	G
	Gs
	A
	As
	B
	
	// Microtonal/custom pitches (not part of 12-tone chromatic scale)
	// Harmonic seventh pitches - approximately 969 cents above their respective roots (31 cents below the minor 7th)
	// Notation: The '♮7' suffix is a custom notation indicating "natural harmonic 7th" to distinguish
	// from tempered intervals. Not to be confused with the natural sign (♮) used to cancel accidentals.
	Ch7  // C harmonic 7th - between A and A# (A super-sharp)
	Csh7 // C# harmonic 7th - between A# and B
	Dh7  // D harmonic 7th - between B and C
	Dsh7 // D# harmonic 7th - between C and C#
	Eh7  // E harmonic 7th - between C# and D
	Fh7  // F harmonic 7th - between D and D#
	Fsh7 // F# harmonic 7th - between D# and E
	Gh7  // G harmonic 7th - between E and F
	Gsh7 // G# harmonic 7th - between F and F#
	Ah7  // A harmonic 7th - between F# and G
	Ash7 // A# harmonic 7th - between G and G#
	Bh7  // B harmonic 7th - between G# and A
)

// NameOf a note will return its Class and Octave
func NameOf(text string) (Class, Octave) {
	return baseNameOf(text).Step(baseStepOf(text))
}

// Step from a class to another class, +/- semitones, +/- octave
func (from Class) Step(inc int) (Class, Octave) {
	return stepFrom(from, inc)
}

// String of the note, expressed with Sharps or Flats
func (from Class) String(with AdjSymbol) string {
	return stringOf(from, with)
}

//
// Private
//

func stringOf(from Class, with AdjSymbol) string {
	switch from {
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case A:
		return "A"
	case B:
		return "B"
	}

	if with == Sharp {
		return stringSharpOf(from)
	} else if with == Flat {
		return stringFlatOf(from)
	}

	return "-"
}

func stringSharpOf(from Class) string {
	switch from {
	case Cs:
		return "C#"
	case Ds:
		return "D#"
	case Fs:
		return "F#"
	case Gs:
		return "G#"
	case As:
		return "A#"
	// Harmonic seventh pitches
	case Ch7:
		return "A♮7"
	case Csh7:
		return "A#♮7"
	case Dh7:
		return "B♮7"
	case Dsh7:
		return "C♮7"
	case Eh7:
		return "C#♮7"
	case Fh7:
		return "D♮7"
	case Fsh7:
		return "D#♮7"
	case Gh7:
		return "E♮7"
	case Gsh7:
		return "F♮7"
	case Ah7:
		return "F#♮7"
	case Ash7:
		return "G♮7"
	case Bh7:
		return "G#♮7"
	}
	return "-"
}

func stringFlatOf(from Class) string {
	switch from {
	case Cs:
		return "Db"
	case Ds:
		return "Eb"
	case Fs:
		return "Gb"
	case Gs:
		return "Ab"
	case As:
		return "Bb"
	// Harmonic seventh pitches (flat notation)
	case Ch7:
		return "B♭♭7"
	case Csh7:
		return "B♭7"
	case Dh7:
		return "C♭7"
	case Dsh7:
		return "D♭♭7"
	case Eh7:
		return "D♭7"
	case Fh7:
		return "E♭♭7"
	case Fsh7:
		return "E♭7"
	case Gh7:
		return "F♭7"
	case Gsh7:
		return "G♭♭7"
	case Ah7:
		return "G♭7"
	case Ash7:
		return "A♭♭7"
	case Bh7:
		return "A♭7"
	}
	return "-"
}

func baseNameOf(text string) Class {
	if len(text) > 0 {
		switch text[0:1] {
		case "C":
			return C
		case "D":
			return D
		case "E":
			return E
		case "F":
			return F
		case "G":
			return G
		case "A":
			return A
		case "B":
			return B
		default:
			return Nil
		}
	} else {
		return Nil
	}
}

func baseStepOf(text string) int {
	if len(text) < 2 {
		return 0
	}

	switch AdjSymbolBegin(text[1:]) {
	case Sharp:
		return 1
	case Flat:
		return -1
	default:
		return 0
	}
}

func stepFrom(name Class, inc int) (Class, Octave) {
	if inc > 0 {
		return stepFromUp(name, inc)
	} else if inc < 0 {
		return stepFromDown(name, -inc)
	}
	return name, 0
}

func stepFromUp(name Class, inc int) (Class, Octave) {
	octave := Octave(0)
	for i := 0; i < inc; i++ {
		shift := stepUp[name]
		name = shift.Name
		octave += shift.Octave
	}
	return name, octave
}

func stepFromDown(name Class, inc int) (Class, Octave) {
	octave := Octave(0)
	for i := 0; i < inc; i++ {
		shift := stepDown[name]
		name = shift.Name
		octave += shift.Octave
	}
	return name, octave
}

type step struct {
	Name   Class
	Octave Octave
}

var (
	stepUp = map[Class]step{
		Nil:  step{Nil, 0},
		C:    step{Cs, 0},
		Cs:   step{D, 0},
		D:    step{Ds, 0},
		Ds:   step{E, 0},
		E:    step{F, 0},
		F:    step{Fs, 0},
		Fs:   step{G, 0},
		G:    step{Gs, 0},
		Gs:   step{A, 0},
		A:    step{As, 0},
		As:   step{B, 0},
		B:    step{C, 1},
		// Harmonic seventh pitches step up to their closest chromatic neighbor
		Ch7:  step{As, 0},
		Csh7: step{B, 0},
		Dh7:  step{C, 0},
		Dsh7: step{Cs, 0},
		Eh7:  step{D, 0},
		Fh7:  step{Ds, 0},
		Fsh7: step{E, 0},
		Gh7:  step{F, 0},
		Gsh7: step{Fs, 0},
		Ah7:  step{G, 0},
		Ash7: step{Gs, 0},
		Bh7:  step{A, 0},
	}
	stepDown = map[Class]step{
		Nil:  step{Nil, 0},
		C:    step{B, -1},
		Cs:   step{C, 0},
		D:    step{Cs, 0},
		Ds:   step{D, 0},
		E:    step{Ds, 0},
		F:    step{E, 0},
		Fs:   step{F, 0},
		G:    step{Fs, 0},
		Gs:   step{G, 0},
		A:    step{Gs, 0},
		As:   step{A, 0},
		B:    step{As, 0},
		// Harmonic seventh pitches step down to their closest chromatic neighbor
		Ch7:  step{A, 0},
		Csh7: step{As, 0},
		Dh7:  step{B, 0},
		Dsh7: step{C, 0},
		Eh7:  step{Cs, 0},
		Fh7:  step{D, 0},
		Fsh7: step{Ds, 0},
		Gh7:  step{E, 0},
		Gsh7: step{F, 0},
		Ah7:  step{Fs, 0},
		Ash7: step{G, 0},
		Bh7:  step{Gs, 0},
	}
)
