package scale_test

import (
	"fmt"

	"gopkg.in/music-theory.v0/note"
	"gopkg.in/music-theory.v0/scale"
)

// Example demonstrates creating a scale by name
func Example() {
	s := scale.Of("C major")
	fmt.Printf("Root: %s\n", s.Root.String(s.AdjSymbol))
	fmt.Printf("Tones: %v\n", len(s.Tones))
	// Output:
	// Root: C
	// Tones: 7
}

// ExampleOf demonstrates creating various scales by name
func ExampleOf() {
	// Major scale
	cmaj := scale.Of("C major")
	fmt.Printf("%s major: %v tones\n", cmaj.Root.String(cmaj.AdjSymbol), len(cmaj.Tones))
	
	// Minor scale
	cmin := scale.Of("C minor")
	fmt.Printf("%s minor: %v tones\n", cmin.Root.String(cmin.AdjSymbol), len(cmin.Tones))
	
	// Augmented scale
	caug := scale.Of("C aug")
	fmt.Printf("%s augmented: %v tones\n", caug.Root.String(caug.AdjSymbol), len(caug.Tones))
	
	// Output:
	// C major: 7 tones
	// C minor: 7 tones
	// C augmented: 6 tones
}

// ExampleScale_Notes demonstrates getting notes from a scale
func ExampleScale_Notes() {
	s := scale.Of("C major")
	notes := s.Notes()
	
	for i, n := range notes {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(n.Class.String(note.Sharp))
	}
	fmt.Println()
	
	// Output: C, D, E, F, G, A, B
}
