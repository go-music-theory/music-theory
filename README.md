[![Tests](https://github.com/go-music-theory/music-theory/actions/workflows/test.yml/badge.svg)](https://github.com/go-music-theory/music-theory/actions/workflows/test.yml) 
[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0?status.svg)](https://godoc.org/gopkg.in/music-theory.v0) 
[![Go Report Card](https://goreportcard.com/badge/gopkg.in/music-theory.v0)](https://goreportcard.com/report/gopkg.in/music-theory.v0) 
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

[gopkg.in/music-theory.v0](https://gopkg.in/music-theory.v0)

[https://github.com/go-music-theory/music-theory](https://github.com/go-music-theory/music-theory)

## Music theory models in Go

There's an example command-line utility `music-theory.go` to demo the libraries, with a `bin/` wrapper.

To build and install `music-theory` to your machine:

    make install

Then, to calculate the note pitch classes for a specified **Chord**:

    $ music-theory chord "Cm nondominant -5 679"
    
    root: C
    tones:
      3: D#
      6: A
      7: A#
      9: D

To list the names of all the known chord-building rules:

    $ music-theory chords
    
    - Basic
    - Nondominant
    - Major Triad
    - Minor Triad
    - Augmented Triad
    - Diminished Triad
    - Suspended Triad
    - Omit Fifth
    - Flat Fifth
    - Add Sixth
    - Augmented Sixth
    - Omit Sixth
    - Add Seventh
    - Dominant Seventh
    - Major Seventh
    - Minor Seventh
    - Diminished Seventh
    - Half Diminished Seventh
    - Diminished Major Seventh
    - Augmented Major Seventh
    - Augmented Minor Seventh
    - Harmonic Seventh
    - Omit Seventh
    - Add Ninth
    - Dominant Ninth
    - Major Ninth
    - Minor Ninth
    - Sharp Ninth
    - Omit Ninth
    - Add Eleventh
    - Dominant Eleventh
    - Major Eleventh
    - Minor Eleventh
    - Omit Eleventh
    - Add Thirteenth
    - Dominant Thirteenth
    - Major Thirteenth
    - Minor Thirteenth

To calculate the note pitch classes for a specified **Scale**:

    $ music-theory scale "C aug"
    
    root: C
    tones:
      1: C
      2: D#
      3: E
      4: G
      5: G#
      6: B

To list the names of all the known scale-building rules:

    $ music-theory scales
    
    - Default (Major)
    - Minor
    - Major
    - Natural Minor
    - Diminished
    - Augmented
    - Melodic Minor Ascend
    - Melodic Minor Descend
    - Harmonic Minor
    - Ionian
    - Dorian
    - Phrygian
    - Lydian
    - Mixolydian
    - Aeolian
    - Locrian

To determine a key:

    $ music-theory key Db
    
    root: Db
    mode: Major
    relative:
      root: Bb
      mode: Minor

##### Credit

[Nick Charney Kaye](https://charneykaye.com)

[XJ Music](https://xj.io)

## [Note](note/)

A Note is used to represent the relative duration and pitch of a sound.

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/note?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/note) [![Coverage](https://raw.githubusercontent.com/wiki/go-music-theory/music-theory/coverage.svg)](https://raw.githack.com/wiki/go-music-theory/music-theory/coverage.html)

## [Key](key/)

The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/key?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/key) [![Coverage](https://raw.githubusercontent.com/wiki/go-music-theory/music-theory/coverage.svg)](https://raw.githack.com/wiki/go-music-theory/music-theory/coverage.html)

## [Chord](chord/)

In music theory, a chord is any harmonic set of three or more notes that is heard as if sounding simultaneously.

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/chord?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/chord) [![Coverage](https://raw.githubusercontent.com/wiki/go-music-theory/music-theory/coverage.svg)](https://raw.githack.com/wiki/go-music-theory/music-theory/coverage.html)

## [Scale](scale/)

In music theory, a scale is any set of musical notes ordered by fundamental frequency or pitch.

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/scale?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/scale) [![Coverage](https://raw.githubusercontent.com/wiki/go-music-theory/music-theory/coverage.svg)](https://raw.githack.com/wiki/go-music-theory/music-theory/coverage.html)
