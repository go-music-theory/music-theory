package note_test

import (
	"fmt"

	"gopkg.in/music-theory.v0/note"
)

// Example demonstrates creating a note by name
func Example() {
	n := note.Named("C")
	fmt.Println(n.Class.String(note.Sharp))
	// Output: C
}

// ExampleNamed demonstrates creating a note from a name string
func ExampleNamed() {
	c := note.Named("C")
	fmt.Println(c.Class.String(note.Sharp))

	ds := note.Named("D#")
	fmt.Println(ds.Class.String(note.Sharp))

	// Output:
	// C
	// D#
}

// ExampleOfClass demonstrates creating a note from a note class
func ExampleOfClass() {
	c := note.OfClass(note.C)
	fmt.Println(c.Class.String(note.Sharp))

	// Output: C
}

// ExampleClassNamed demonstrates getting a note class from a name
func ExampleClassNamed() {
	c := note.ClassNamed("C")
	fmt.Println(c.String(note.Sharp))

	// Output: C
}

// ExampleClass_String demonstrates converting a note class to a string
func ExampleClass_String() {
	// C# as sharp
	cs := note.Cs
	fmt.Println(cs.String(note.Sharp))

	// C# as flat (Db)
	fmt.Println(cs.String(note.Flat))

	// Output:
	// C#
	// Db
}
