// Package main implements a command-line utility for music
package main

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

type mockWriter struct {
	written []byte
}

func (fw *mockWriter) Write(p []byte) (n int, err error) {
	if fw.written == nil {
		fw.written = p
	} else {
		fw.written = append(fw.written, p...)
	}

	return len(p), nil
}

func (fw *mockWriter) GetWritten() (w string) {
	return string(fw.written[:len(fw.written)])
}

func TestMusicTheory(t *testing.T) {
	app := app()
	w := &mockWriter{}
	app.Writer = w

	err := app.Run([]string{"cmd",
		"chord", "G minor 7",
	})

	assert.Nil(t, err)
	assert.Equal(t, "root: G\ntones:\n  1: G\n  3: Bb\n  5: D\n  7: F\n", w.GetWritten())
}

func TestHelpCmd(t *testing.T) {
	app := app()
	w := &mockWriter{}
	app.Writer = w

	err := app.Run([]string{"cmd",
		"-h",
	})

	assert.Nil(t, err)
	assert.Contains(t, w.GetWritten(), "USAGE")
	assert.Contains(t, w.GetWritten(), "[global options] command [command options] [arguments...]")
}

func TestNotePitch_TwoArguments(t *testing.T) {
	app := app()
	w := &mockWriter{}
	app.Writer = w

	err := app.Run([]string{"cmd",
		"pitch", "A", "4",
	})

	assert.Nil(t, err)
	assert.Equal(t, "440.00Hz\n", w.GetWritten())
}

func TestNotePitch_OneArgument(t *testing.T) {
	app := app()
	w := &mockWriter{}
	app.Writer = w

	err := app.Run([]string{"cmd",
		"pitch", "A4",
	})

	assert.Nil(t, err)
	assert.Equal(t, "440.00Hz\n", w.GetWritten())
}
