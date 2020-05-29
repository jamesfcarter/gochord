package gochord

import (
	"strconv"
	"strings"
)

const (
	Open  = 0
	Muted = -1
)

type String struct {
	Finger string
	Fret   int
}

type Chord struct {
	Name    string
	Capo    int
	Strings []String
}

func IsFretted(fret int) bool {
	return fret != Open && fret != Muted
}

func NewChord(name, layout string) Chord {
	chrs := strings.Split(layout, "")
	strs := make([]String, len(chrs))
	for i, c := range chrs {
		switch c {
		case "o":
			strs[i].Fret = Open
		case "x":
			strs[i].Fret = Muted
		default:
			f, _ := strconv.Atoi(c)
			strs[i].Fret = f
			strs[i].Finger = "*"
		}
	}
	return Chord{
		Name:    name,
		Capo:    0,
		Strings: strs,
	}
}

func (c Chord) StringCount() int {
	return len(c.Strings)
}

func (c Chord) HighestFret() int {
	highest := c.Capo
	for _, s := range c.Strings {
		if !IsFretted(s.Fret) {
			continue
		}
		if s.Fret > highest {
			highest = s.Fret
		}
	}
	return highest
}

func (c Chord) LowestFret() int {
	lowest := c.Capo
	for _, s := range c.Strings {
		if !IsFretted(s.Fret) {
			continue
		}
		if s.Fret < lowest || lowest == 0 {
			lowest = s.Fret
		}
	}
	return lowest
}

func (c Chord) FirstFret() int {
	if c.Capo > 0 || c.HighestFret()-FretBoardSize < c.Capo {
		return c.Capo
	}
	if c.HighestFret()-FretBoardSize < c.LowestFret() {
		return c.HighestFret() - FretBoardSize
	}
	return c.LowestFret() - 1
}

func (c Chord) FretBoard() FretBoard {
	fb := FretBoard{
		FirstFret: c.FirstFret(),
		Unfretted: make([]string, c.StringCount()),
		Frets:     make([][]string, FretBoardSize),
	}
	for i := range fb.Frets {
		fb.Frets[i] = make([]string, c.StringCount())
	}
	for i, s := range c.Strings {
		switch s.Fret {
		case Open:
			fb.Unfretted[i] = "o"
		case Muted:
			fb.Unfretted[i] = "x"
		default:
			fb.Unfretted[i] = " "
			// FIXME: check index
			fb.Frets[fb.FretIndex(s.Fret)][i] = s.Finger
		}
	}
	return fb
}
