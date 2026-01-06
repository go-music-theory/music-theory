package chord_test

import (
	"fmt"

	"gopkg.in/music-theory.v0/chord"
	"gopkg.in/music-theory.v0/note"
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
