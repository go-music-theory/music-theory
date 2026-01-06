package key_test

import (
	"fmt"

	"github.com/go-music-theory/music-theory/key"
	"github.com/go-music-theory/music-theory/note"
)

// Example demonstrates creating a key by name
func Example() {
	k := key.Of("C")
	fmt.Printf("Root: %s\n", k.Root.String(k.AdjSymbol))
	fmt.Printf("Mode: %s\n", k.Mode)
	// Output:
	// Root: C
	// Mode: Major
}

// ExampleOf demonstrates creating various keys by name
func ExampleOf() {
	// Major key
	c := key.Of("C")
	fmt.Printf("%s %s\n", c.Root.String(c.AdjSymbol), c.Mode)

	// Minor key
	cm := key.Of("Cm")
	fmt.Printf("%s %s\n", cm.Root.String(cm.AdjSymbol), cm.Mode)

	// Key with flats
	db := key.Of("Db")
	fmt.Printf("%s %s\n", db.Root.String(db.AdjSymbol), db.Mode)

	// Output:
	// C Major
	// C Minor
	// Db Major
}

// ExampleKey_RelativeMinor demonstrates getting the relative minor key
func ExampleKey_RelativeMinor() {
	c := key.Of("C")
	rel := c.RelativeMinor()

	fmt.Printf("%s %s -> %s %s\n",
		c.Root.String(c.AdjSymbol), c.Mode,
		rel.Root.String(rel.AdjSymbol), rel.Mode)

	// Output: C Major -> A Minor
}

// ExampleFindKey demonstrates the Krumhansl-Schmuckler key-finding algorithm
func ExampleFindKey() {
	// Identify the key from a C major scale
	notes := []note.Class{note.C, note.D, note.E, note.F, note.G, note.A, note.B}
	k := key.FindKey(notes)
	fmt.Printf("%s %s\n", k.Root.String(k.AdjSymbol), k.Mode)

	// Identify the key from musical phrase with emphasis on the tonic
	notes = []note.Class{note.A, note.A, note.C, note.E, note.A, note.G, note.F, note.E}
	k = key.FindKey(notes)
	fmt.Printf("%s %s\n", k.Root.String(k.AdjSymbol), k.Mode)

	// Output:
	// C Major
	// A Minor
}
