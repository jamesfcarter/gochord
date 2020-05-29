package gochord

import (
	"strconv"
	"strings"
)

const (
	FretBoardSize = 4
)

type FretBoard struct {
	FirstFret int
	Unfretted []string
	Frets     [][]string
}

func (fb FretBoard) FretIndex(fret int) int {
	return fret - fb.FirstFret - 1
}

func (fb FretBoard) Strings() int {
	return len(fb.Frets[0])
}

func (fb FretBoard) nutString() string {
	s := "="
	for i := 1; i < fb.Strings(); i++ {
		s = s + "=="
	}
	return s
}

func (fb FretBoard) fretString() string {
	a := make([]string, fb.Strings())
	for i := range a {
		a[i] = "+"
	}
	return strings.Join(a, "-")
}

func (fb FretBoard) fingeringString(index int) string {
	a := make([]string, fb.Strings())
	for i := range a {
		if fb.Frets[index][i] == "" {
			a[i] = "|"
		} else {
			a[i] = fb.Frets[index][i]
		}
	}
	return strings.Join(a, " ")
}

func (fb FretBoard) emptyString() string {
	a := make([]string, fb.Strings())
	for i := range a {
		a[i] = "|"
	}
	return strings.Join(a, " ")
}

func (fb FretBoard) unfrettedString() string {
	a := make([]string, fb.Strings())
	for i := range a {
		a[i] = fb.Unfretted[i]
	}
	return strings.Join(a, " ")
}

func (fb FretBoard) String() string {
	s := fb.unfrettedString() + "\n"
	if fb.FirstFret == 0 {
		s = s + fb.nutString() + "\n"
	} else {
		s = s + fb.fretString() + " " + strconv.Itoa(fb.FirstFret) + "\n"
	}
	for i := range fb.Frets {
		s = s + fb.fingeringString(i) + "\n"
		s = s + fb.fretString() + "\n"
	}
	return s
}
