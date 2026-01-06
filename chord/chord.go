// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
//
// https://en.wikipedia.org/wiki/Chord_(music)
//
// # Credit
//
// Charney Kaye
// <hi@charneykaye.com>
// https://charneykaye.com
//
// XJ Music
// https://xj.io
package chord

import (
	"strings"

	"github.com/go-music-theory/music-theory/note"
)

// Chord in a particular key
type Chord struct {
	Root      note.Class
	AdjSymbol note.AdjSymbol
	Tones     map[Interval]note.Class
	Bass      note.Class // Bass note for slash chords (e.g., C/E has E as bass)
}

// Of a particular key, e.g. Of("C minor 7")
func Of(name string) Chord {
	c := Chord{}
	c.parse(name)
	return c
}

// Notes to obtain the notes from the Chord
func (this *Chord) Notes() (notes []*note.Note) {
	// If there's a bass note (slash chord), add it first
	if this.Bass != note.Nil {
		notes = append(notes, note.OfClass(this.Bass))
	}
	
	// Check if bass note exists in chord tones to avoid duplication
	bassInTones := false
	if this.Bass != note.Nil {
		for _, class := range this.Tones {
			if class == this.Bass {
				bassInTones = true
				break
			}
		}
	}
	
	forAllIn(this.Tones, func(class note.Class) {
		// Skip bass note if it's already added and exists in tones
		if !bassInTones || class != this.Bass {
			notes = append(notes, note.OfClass(class))
		}
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
	
	// Transpose bass note if it exists
	if this.Bass != note.Nil {
		transposedChord.Bass, _ = this.Bass.Step(semitones)
	}
	
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
	this.Bass = note.Nil // Initialize bass note as Nil

	// determine whether the name is "sharps" or "flats"
	this.AdjSymbol = note.AdjSymbolOf(name)

	// Check for slash chord notation (e.g., "C/E" or "Cmaj7/B")
	slashIndex := strings.Index(name, "/")

	if slashIndex != -1 {
		// Parse bass note from slash notation
		bassString := name[slashIndex+1:]
		this.Bass, _ = note.RootAndRemaining(bassString)
		
		// Parse the chord part before the slash
		name = name[:slashIndex]
	}

	// parse the root, and keep the remaining string
	this.Root, name = note.RootAndRemaining(name)

	// parse the chord Form
	this.parseForms(name)
}
