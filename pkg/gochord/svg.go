package gochord

import (
	"fmt"
	"io"
	"strconv"

	svgo "github.com/ajstarks/svgo"
)

type g struct {
	width  int
	height int
	hChunk int
	vChunk int
}

func (g g) lineStyle(size int) string {
	lineSize := g.width / 100
	if lineSize == 0 {
		lineSize = 1
	}
	lineSize = lineSize * size
	return fmt.Sprintf("stroke:rgb(0,0,0);stroke-width:%dpx", lineSize)
}

func (g g) hUnit(size int) int {
	return g.hChunk * size
}

func (g g) vUnit(size int) int {
	return g.vChunk * size
}

func newG(width, height int) g {
	return g{
		width:  width,
		height: height,
		hChunk: width / 40,
		vChunk: height / 40,
	}
}

func spotDiameter(h, v int) int {
	base := h
	if v < h {
		base = v
	}
	return base * 7 / 20
}

func (c Chord) SVG(width, height int, out io.Writer) {
	svg := svgo.New(out)
	svg.Start(width, height)
	defer svg.End()

	g := newG(width, height)
	leftBorder := g.hUnit(4)
	rightBorder := g.hUnit(4)
	topBorder := g.vUnit(11)
	bottomBorder := g.vUnit(4)

	fb := c.FretBoard()
	stringSpacing := (width - leftBorder - rightBorder) / (c.StringCount() - 1)
	fretSpacing := (height - topBorder - bottomBorder) / len(fb.Frets)

	svg.Text(g.hUnit(4), g.vUnit(6), c.Name, fmt.Sprintf("text-anchor:left;font-size:%dpx", g.hUnit(6)))
	if fb.FirstFret != 0 {
		svg.Text(width-rightBorder+g.hUnit(2), topBorder+g.vUnit(1), strconv.Itoa(fb.FirstFret), fmt.Sprintf("text-anchor:middle;font-size:%dpx", g.hUnit(6)))
	}

	for s := 0; s < c.StringCount(); s++ {
		x := leftBorder + s*stringSpacing
		svg.Line(
			x,
			topBorder,
			x,
			height-bottomBorder,
			g.lineStyle(1),
		)
		svg.Text(x, topBorder-g.vUnit(1), fb.Unfretted[s], fmt.Sprintf("text-anchor:middle;font-size:%dpx", g.hUnit(6)))
	}

	for f := 0; f <= len(fb.Frets); f++ {
		size := 1
		if f == 0 && fb.FirstFret == 0 {
			size = 5
		}
		svg.Line(
			leftBorder,
			topBorder+f*fretSpacing,
			width-rightBorder,
			topBorder+f*fretSpacing,
			g.lineStyle(size),
		)
	}

	diameter := spotDiameter(stringSpacing, fretSpacing)
	for f, fret := range fb.Frets {
		for s, str := range fret {
			if str == "" {
				continue
			}
			svg.Circle(
				leftBorder+s*stringSpacing,
				topBorder+f*fretSpacing+fretSpacing/2,
				diameter,
			)
		}
	}
}
