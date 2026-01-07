# GitHub Copilot Instructions for music-theory

## Project Overview

This is a Go library for music theory modeling, providing structures and operations for:
- **Notes**: Pitch classes and their representations
- **Chords**: Harmonic sets of three or more notes
- **Scales**: Sets of musical notes ordered by pitch
- **Keys**: Groups of pitches upon which compositions are based

The project includes a CLI tool (`music-theory.go`) for demonstrating the library capabilities.

## Project Structure

```
music-theory/
├── note/          # Note pitch classes, octaves, and pitch calculations
├── chord/         # Chord construction, intervals, and techniques
├── scale/         # Scale modes and construction
├── key/           # Key determination and relationships
├── music-theory.go # Main CLI application
└── bin/           # CLI wrapper scripts
```

## Code Style and Conventions

### Go Formatting
- Use `go fmt` for all Go code formatting (tabs for indentation, 4-space tab width)
- Follow standard Go naming conventions (PascalCase for exported, camelCase for unexported)
- Use meaningful variable names that reflect music theory concepts

### EditorConfig Settings
- Go files: tabs for indentation (size 4)
- Other files: 2 spaces for indentation
- LF line endings, UTF-8 encoding
- Trim trailing whitespace (except markdown)
- Insert final newline

### Documentation
- Use GoDoc-style comments for all exported types, functions, and methods
- Include usage examples in documentation comments where helpful
- Reference music theory terminology accurately

## Testing

### Test Framework
- Use `gopkg.in/stretchr/testify.v1/assert` for assertions
- Test files follow the `*_test.go` naming convention
- Place tests in the same package as the code being tested

### Test Patterns
```go
func TestFunctionName(t *testing.T) {
    result := FunctionToTest()
    assert.Equal(t, expected, result)
}
```

### Running Tests
- Run all tests: `make test` or `go test ./...`
- Tests must pass before building the application
- Aim for comprehensive test coverage (this project has 100% coverage in core packages)

## Building and Development

### Commands
- **Format code**: `make fmt` or `go fmt ./...`
- **Run tests**: `make test` or `go test ./...`
- **Build**: `make` or `make all` (runs deps, test, and build)
- **Install**: `make install` (installs to `/usr/local/bin`)

### Dependencies
- Go 1.14+ (CI tests on Go 1.20, 1.21, 1.22)
- Install dependencies: `go get -v -t ./...`
- Dependencies are managed via `go.mod`

### CLI Usage Examples
```bash
# Determine a chord
music-theory chord "Cm nondominant -5 679"

# List all chords
music-theory chords

# Determine a scale
music-theory scale "C aug"

# Find a key
music-theory key Db

# Calculate pitch frequency
music-theory pitch A4
```

## Music Theory Domain Guidelines

### Terminology
- Use standard Western music theory terminology
- **Root**: The fundamental note of a chord or scale
- **Pitch Class**: The set of all pitches that share the same note name (e.g., all C's)
- **Interval**: The distance between two pitches
- **Mode**: A type of musical scale (Ionian, Dorian, Phrygian, etc.)

### Data Structures
- Notes are represented as classes with optional octave information
- Chords and scales use maps of intervals to notes (`map[Interval]note.Note`)
- Keys include mode information and relative key relationships

### Naming Patterns
- Use `Of()` functions for constructing objects from names (e.g., `chord.Of("Cm")`)
- Use `Named()` functions for note construction (e.g., `note.Named("C")`)
- Use `ToYAML()` methods for human-readable output

## Code Quality

### Before Submitting Code
1. Run `go fmt ./...` to format code
2. Run `go test ./...` to ensure all tests pass
3. Ensure new code has corresponding tests
4. Update documentation comments for public APIs
5. Verify CLI examples work if modifying the main application

### Adding New Features
- Music theory features should be accurate and follow standard conventions
- Consider edge cases (enharmonic equivalents, different tunings, etc.)
- Add tests that validate the music theory correctness
- Update README.md if adding new CLI commands or major features

## Common Pitfalls to Avoid

- Don't hardcode tuning frequencies (use the tuning parameter system)
- Don't assume octaves are always specified (notes can be pitch classes without octaves)
- Be careful with enharmonic equivalents (C# vs Db)
- Maintain the existing YAML output format for CLI consistency
- Don't break the existing API for library users (this is a published package)
