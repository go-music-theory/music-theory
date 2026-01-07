// Chords are expressed in readable strings, e.g. CMb5b7 or Cm679-5
package chord

import (
	"github.com/go-music-theory/music-theory/note"
	"gopkg.in/yaml.v2"
)

func (c Chord) ToYAML() string {
	spec := specFrom(c)
	out, _ := yaml.Marshal(spec)
	return string(out[:])
}

//
// Private
//

func specFrom(c Chord) specChord {
	s := specChord{}
	s.Root = c.Root.String(c.AdjSymbol)
	s.Tones = make(map[int]string)
	for i, t := range c.Tones {
		s.Tones[int(i)] = t.String(c.AdjSymbol)
	}
	// Include bass note if present (slash chord)
	if c.Bass != note.Nil {
		s.Bass = c.Bass.String(c.AdjSymbol)
	}
	return s
}

type specChord struct {
	Root  string
	Bass  string `yaml:"bass,omitempty"`
	Tones map[int]string
}
