package key_test

import (
	"fmt"

	"gopkg.in/music-theory.v0/key"
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
