package chord_test

import (
	"fmt"

	"github.com/go-music-theory/music-theory/chord"
	"github.com/go-music-theory/music-theory/note"
)

// Example demonstrates creating a chord by name
func Example() {
	c := chord.Of("Cmaj7")
	fmt.Printf("Root: %s\n", c.Root.String(c.AdjSymbol))
	fmt.Printf("Tones: %v\n", len(c.Tones))
	// Output:
	// Root: C
	// Tones: 4
}

// ExampleOf demonstrates creating various chords by name
func ExampleOf() {
	// Major chord
	cmaj := chord.Of("C")
	fmt.Printf("%s: %v tones\n", cmaj.Root.String(cmaj.AdjSymbol), len(cmaj.Tones))

	// Minor chord
	cmin := chord.Of("Cm")
	fmt.Printf("%s minor: %v tones\n", cmin.Root.String(cmin.AdjSymbol), len(cmin.Tones))

	// Seventh chord
	c7 := chord.Of("C7")
	fmt.Printf("%s seventh: %v tones\n", c7.Root.String(c7.AdjSymbol), len(c7.Tones))

	// Output:
	// C: 3 tones
	// C minor: 3 tones
	// C seventh: 4 tones
}

// ExampleChord_Notes demonstrates getting notes from a chord
func ExampleChord_Notes() {
	c := chord.Of("C")
	notes := c.Notes()

	for _, n := range notes {
		fmt.Println(n.Class.String(note.Sharp))
	}

	// Output:
	// C
	// E
	// G
}

// ExampleChord_Transpose demonstrates transposing a chord
func ExampleChord_Transpose() {
	c := chord.Of("C")

	// Transpose up by 2 semitones (to D)
	d := c.Transpose(2)
	fmt.Printf("Original: %s\n", c.Root.String(c.AdjSymbol))
	fmt.Printf("Transposed: %s\n", d.Root.String(d.AdjSymbol))

	// Output:
	// Original: C
	// Transposed: D
}

// ExampleOf_slashChord demonstrates creating slash chords
func ExampleOf_slashChord() {
	// Slash chord: C major with E in bass (first inversion)
	c := chord.Of("C/E")
	fmt.Printf("Root: %s\n", c.Root.String(c.AdjSymbol))
	fmt.Printf("Bass: %s\n", c.Bass.String(c.AdjSymbol))

	notes := c.Notes()
	fmt.Printf("Notes (bass first): ")
	for i, n := range notes {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", n.Class.String(note.Sharp))
	}
	fmt.Println()

	// Output:
	// Root: C
	// Bass: E
	// Notes (bass first): E, C, G
}

// ExampleChord_Transpose_slashChord demonstrates transposing a slash chord
func ExampleChord_Transpose_slashChord() {
	// Create a slash chord
	c := chord.Of("C/E")

	// Transpose up by 2 semitones (C/E -> D/F#)
	d := c.Transpose(2)
	fmt.Printf("Original: %s/%s\n", c.Root.String(c.AdjSymbol), c.Bass.String(c.AdjSymbol))
	fmt.Printf("Transposed: %s/%s\n", d.Root.String(d.AdjSymbol), d.Bass.String(d.AdjSymbol))

	// Output:
	// Original: C/E
	// Transposed: D/F#
}
