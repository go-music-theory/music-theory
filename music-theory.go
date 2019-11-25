// Notes, Keys, Chords and Scales
//
// There's an example command-line utility `music-theory.go` to demo the libraries, with a `bin/` wrapper.
//
// Build and install `music-theory` to your machine
//
//     make install
//
// Determine a Chord
//
//     $ music-theory chord "Cm nondominant -5 679"
//
//     root: C
//     tones:
//       3: D#
//       6: A
//       7: A#
//       9: D
//
// List known chord-building rules
//
//     $ music-theory chords
//
//     - Basic
//     - Nondominant
//     - Major Triad
//     - Minor Triad
//     - Augmented Triad
//     - Diminished Triad
//     - Suspended Triad
//     - Omit Fifth
//     - Flat Fifth
//     - Add Sixth
//     - Augmented Sixth
//     - Omit Sixth
//     - Add Seventh
//     - Dominant Seventh
//     - Major Seventh
//     - Minor Seventh
//     - Diminished Seventh
//     - Half Diminished Seventh
//     - Diminished Major Seventh
//     - Augmented Major Seventh
//     - Augmented Minor Seventh
//     - Harmonic Seventh
//     - Omit Seventh
//     - Add Ninth
//     - Dominant Ninth
//     - Major Ninth
//     - Minor Ninth
//     - Sharp Ninth
//     - Omit Ninth
//     - Add Eleventh
//     - Dominant Eleventh
//     - Major Eleventh
//     - Minor Eleventh
//     - Omit Eleventh
//     - Add Thirteenth
//     - Dominant Thirteenth
//     - Major Thirteenth
//     - Minor Thirteenth
//
// Determine a Scale
//
//     $ music-theory scale "C aug"
//
//     root: C
//     tones:
//     1: C
//     2: D#
//     3: E
//     4: G
//     5: G#
//     6: B
//
// List known scale-building rules
//
//     $ music-theory scales
//
//     - Default (Major)
//     - Minor
//     - Major
//     - Natural Minor
//     - Diminished
//     - Augmented
//     - Melodic Minor Ascend
//     - Melodic Minor Descend
//     - Harmonic Minor
//     - Ionian
//     - Dorian
//     - Phrygian
//     - Lydian
//     - Mixolydian
//     - Aeolian
//     - Locrian
//
// Determine a key
//
//    $ music-theory key Db
//
//    root: Db
//    mode: Major
//    relative:
//      root: Bb
//      mode: Minor
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
package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/go-music-theory/music-theory/chord"
	"github.com/go-music-theory/music-theory/key"
	"github.com/go-music-theory/music-theory/pitch"
	"github.com/go-music-theory/music-theory/scale"
)

func main() {
	app := app()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(app.Writer, "Error occurred: %v\n", err)
	}
}

func app() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "music-theory"
	app.Usage = "Notes, Keys, Chords and Scales"
	app.Version = "0.0.4"
	app.Authors = []cli.Author{
		{Name: "Charney Kaye", Email: "hi@charneykaye.com"},
	}
	app.Commands = commands
	return app
}

var commands = []cli.Command{

	{ // Build a Chord
		Name:        "chord",
		Aliases:     []string{"c"},
		Usage:       "build a Chord",
		Description: "Chord is a named harmonic set of three or more pitch classes specified by a name, e.g. C or Cm6 or D♭m679-5",
		Action: func(c *cli.Context) {
			name := c.Args().First()
			if len(name) > 0 {
				fmt.Fprintf(c.App.Writer, "%s", chord.Of(name).ToYAML())
			} else {
				// no arguments
				err := cli.ShowCommandHelp(c, "chord")
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error occurred: %v\n", err)
				}
			}
		},
	},

	{ // List all Chords
		Name:        "chords",
		Usage:       "list all known Chords",
		Description: "The Chord DNA is this software is a sequential chain of rules to be executed by matching text in the chord name to its musical implications from the root of the chord.",
		Action: func(c *cli.Context) {
			fmt.Fprintf(c.App.Writer, "%s", chord.ChordFormList.ToYAML())
		},
	},

	{ // Build a Scale
		Name:        "scale",
		Aliases:     []string{"c"},
		Usage:       "build a Scale",
		Description: "Scale is any set of musical notes ordered by fundamental frequency or pitch specified by a name, e.g. C or Cm6 or D♭m679-5",
		Action: func(c *cli.Context) {
			name := c.Args().First()
			if len(name) > 0 {
				fmt.Fprintf(c.App.Writer, "%s", scale.Of(name).ToYAML())
			} else {
				// no arguments
				err := cli.ShowCommandHelp(c, "scale")
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error occurred: %v\n", err)
				}
			}
		},
	},

	{ // List all Scales
		Name:        "scales",
		Usage:       "list all known Scales",
		Description: "The Scale DNA is this software is a sequential chain of rules to be executed by matching text in the scale name to its musical implications from the root of the scale.",
		Action: func(c *cli.Context) {
			fmt.Fprintf(c.App.Writer, "%s", scale.ScaleModeList.ToYAML())
		},
	},

	{ // Find a Key
		Name:        "key",
		Aliases:     []string{"k"},
		Usage:       "find a Key",
		Description: "The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.",
		Action: func(c *cli.Context) {
			name := c.Args().First()
			if len(name) > 0 {
				fmt.Fprintf(c.App.Writer, "%s", key.Of(name).ToYAML())
			} else {
				// no arguments
				err := cli.ShowCommandHelp(c, "key")
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error occurred: %v\n", err)
				}
			}
		},
	},

	{ // Find a Note Pitch
		Name:        "pitch",
		Aliases:     []string{"p"},
		Usage:       "find a note pitch in Hz",
		Description: "The pitch is note frequency described in Hz. Based on standard concert pitch and twelve-tone equal temperament. As an argument, pass a note in international pitch notation.",
		Flags: []cli.Flag{
		    cli.IntFlag{Name: "tuning, t", Value: 440, Usage: "Set the pitch of the root note A 4"},
	    },
		Action: func(c *cli.Context) {
			name := c.Args().First()
			octave := c.Args().Get(1)
			tuning := c.Int("tuning")
			if len(name) > 0 {
				var notePitch string
				var err error
				if len(octave) > 0 {
					notePitch, err = pitch.OfClassAndOctave(name, octave, tuning)
				} else {
					notePitch, err = pitch.OfNote(name, tuning)
				}
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error occurred: %v\n", err)
				}
				fmt.Fprintf(c.App.Writer, "%v\n", notePitch)
			} else {
				// no arguments
				err := cli.ShowCommandHelp(c, "pitch")
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error occurred: %v\n", err)
				}
			}
		},
	},
}
