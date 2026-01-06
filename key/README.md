# Key

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/key?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/key) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/gopkg.in/music-theory.v0/key)

#### A model of a musical key signature.

The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.

[Musical Key on Wikipedia](https://en.wikipedia.org/wiki/Key_(music))

## Features

### Key Parsing

Create keys by name:

```go
k := key.Of("C major")
fmt.Printf("%s %s\n", k.Root.String(k.AdjSymbol), k.Mode)
// Output: C Major

km := key.Of("Am")
fmt.Printf("%s %s\n", km.Root.String(km.AdjSymbol), km.Mode)
// Output: A Minor
```

### Key-Finding Algorithm

The package includes the Krumhansl-Schmuckler key-finding algorithm, which can determine the most likely key from a collection of notes:

```go
// Identify the key from a musical phrase
notes := []note.Class{note.C, note.D, note.E, note.F, note.G, note.A, note.B}
k := key.FindKey(notes)
fmt.Printf("%s %s\n", k.Root.String(k.AdjSymbol), k.Mode)
// Output: C Major
```

The algorithm works by:
1. Calculating the pitch class distribution from the input notes
2. Correlating this distribution with empirically-derived key profiles for all 24 major and minor keys
3. Returning the key with the highest correlation coefficient

For better accuracy when distinguishing between relative major/minor keys, provide notes with emphasis on the tonic (repeat important notes):

```go
notes := []note.Class{note.A, note.A, note.C, note.E, note.A, note.G, note.F, note.E}
k := key.FindKey(notes)
fmt.Printf("%s %s\n", k.Root.String(k.AdjSymbol), k.Mode)
// Output: A Minor
```

[Krumhansl-Schmuckler Key-Finding Algorithm](http://rnhart.net/articles/key-finding/)

##### Credit

[Charney Kaye](https://charneykaye.com)

[XJ Music](https://xj.io)

