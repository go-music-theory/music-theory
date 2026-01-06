// In music theory, a scale is any set of musical notes ordered by fundamental frequency or pitch.
//
// https://en.wikipedia.org/wiki/Scale_(music)
//
// A scale ordered by increasing pitch is an ascending scale, and a scale ordered by decreasing pitch is a descending scale. Some scales contain different pitches when ascending than when descending. For example, the Melodic minor scale.
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
package scale

import (
	"github.com/go-music-theory/music-theory/note"
)

// Scale in a particular key
type Scale struct {
	Root      note.Class
	AdjSymbol note.AdjSymbol
	Tones     map[Interval]note.Class
}

// Of a particular key, e.g. Of("C minor 7")
func Of(name string) Scale {
	c := Scale{}
	c.parse(name)
	return c
}

// Notes to obtain the notes from the Scale
func (this *Scale) Notes() (notes []*note.Note) {
	forAllIn(this.Tones, func(class note.Class) {
		notes = append(notes, note.OfClass(class))
	})
	return
}

//
// Private
//

func (this *Scale) parse(name string) {
	this.Tones = make(map[Interval]note.Class)

	// determine whether the name is "sharps" or "flats"
	this.AdjSymbol = note.AdjSymbolOf(name)

	// parse the root, and keep the remaining string
	this.Root, name = note.RootAndRemaining(name)

	// parse the scale Mode
	this.parseModes(name)
}
