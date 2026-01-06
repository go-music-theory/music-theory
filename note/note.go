// A Note is used to represent the relative duration and pitch of a sound.
//
// https://en.wikipedia.org/wiki/Musical_note
//
// # Credit
//
// Charney Kaye
// <hi@charneykaye.com>
// https://charneykaye.com
//
// XJ Music
// https://xj.io
package note

// Note models a musical note
type Note struct {
	Class  Class  // Class of pitch
	Octave Octave // Octave #

	Performer string  // Can be used to sort out whose Notes are whose
	Position  float64 // Can be used to represent time within the composition
	Duration  float64 // Can be used to represent time of note duration
	Code      string  // Can be used to store any custom values
}

// Named note returns a Note model
func Named(text string) (n *Note) {
	n = &Note{}

	// First the name, including octave shift.
	n.Class, n.Octave = NameOf(text)

	// Last, add the originally named octave.
	n.Octave += OctaveOf(text)

	return
}

// OfClass pitch returns a Note model
func OfClass(class Class) (n *Note) {
	n = &Note{}
	n.Class = class
	return
}

// ClassNamed returns a pitch Class
func ClassNamed(text string) Class {
	n := Named(text)
	return n.Class
}
