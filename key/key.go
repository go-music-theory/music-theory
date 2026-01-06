// The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
//
// https://en.wikipedia.org/wiki/Key_(music)
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
package key

import (
	"github.com/go-music-theory/music-theory/note"
)

// Of a particular key, e.g. Of("C minor 7")
func Of(name string) Key {
	k := Key{}
	k.parse(name)
	return k
}

// Key is a model of a musical key signature
type Key struct {
	Root      note.Class
	AdjSymbol note.AdjSymbol
	Mode      Mode
}

//
// Private
//

func (this *Key) parse(name string) {
	// determine whether the name is "sharps" or "flats"
	this.AdjSymbol = note.AdjSymbolOf(name)

	// parse the root, and keep the remaining string
	this.Root, name = note.RootAndRemaining(name)

	// parse the key mode
	this.parseMode(name)
}
