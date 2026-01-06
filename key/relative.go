// The relative minor of a major key has the same key signature and starts down a minor third (or equivalently up a major sixth); for example, the relative minor of G major is E minor. Similarly the relative major of a minor key starts up a minor third (or down a major sixth); for example, the relative major of F minor is A♭ major.
package key

import (
	"github.com/go-music-theory/music-theory/note"
)

func (k Key) RelativeMinor() (rk Key) {
	rk = k
	if rk.Mode == Major {
		rk.Mode = Minor
		rk.AdjSymbol = note.Flat
		rk.Root, _ = rk.Root.Step(-3)
	}
	return
}

func (k Key) RelativeMajor() (rk Key) {
	rk = k
	if rk.Mode == Minor {
		rk.Mode = Major
		rk.AdjSymbol = note.Sharp
		rk.Root, _ = rk.Root.Step(3)
	}
	return
}
