// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
//
// https://en.wikipedia.org/wiki/Chord_(music)
//
// Credit
//
// Charney Kaye
// <hi@charneykaye.com>
// https://charneykaye.com
//
// XJ Music
// https://xj.io
//
package chord

import (
	"github.com/go-music-theory/music-theory/note"
)

// Chord in a particular key
type Chord struct {
	Root      note.Class
	AdjSymbol note.AdjSymbol
	Tones     map[Interval]note.Class
}

// Of a particular key, e.g. Of("C minor 7")
func Of(name string) Chord {
	c := Chord{}
	c.parse(name)
	return c
}

// Notes to obtain the notes from the Chord
func (this *Chord) Notes() (notes []*note.Note) {
	forAllIn(this.Tones, func(class note.Class) {
		notes = append(notes, note.OfClass(class))
	})
	return
}

// Transpose a chord +/- semitones
func (this Chord) Transpose(semitones int) Chord {
	transposedChord := Chord{
		AdjSymbol: this.AdjSymbol,
		Tones:     make(map[Interval]note.Class),
	}
	transposedChord.Root, _ = this.Root.Step(semitones)
	for interval, class := range this.Tones {
		transposedChord.Tones[interval], _ = class.Step(semitones)
	}
	return transposedChord
}

//
// Private
//

func (this *Chord) parse(name string) {
	this.Tones = make(map[Interval]note.Class)

	// determine whether the name is "sharps" or "flats"
	this.AdjSymbol = note.AdjSymbolOf(name)

	// parse the root, and keep the remaining string
	this.Root, name = note.RootAndRemaining(name)

	// parse the chord Form
	this.parseForms(name)
}
