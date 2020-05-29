package gochord

import (
	"testing"
)

func TestFingerBoardString(t *testing.T) {
	for _, test := range []struct {
		name     string
		chord    Chord
		expected string
	}{
		{
			name:  "A",
			chord: NewChord("A", "x02220"),
			expected: `
x o       o
===========
| | | | | |
+-+-+-+-+-+
| | * * * |
+-+-+-+-+-+
| | | | | |
+-+-+-+-+-+
| | | | | |
+-+-+-+-+-+
`,
		},
		{
			name:  "F",
			chord: NewChord("F", "133211"),
			expected: `
           
===========
* | | | * *
+-+-+-+-+-+
| | | * | |
+-+-+-+-+-+
| * * | | |
+-+-+-+-+-+
| | | | | |
+-+-+-+-+-+
`,
		},
		{
			name:  "A barre",
			chord: NewChord("A", "577655"),
			expected: `
           
+-+-+-+-+-+ 3
| | | | | |
+-+-+-+-+-+
* | | | * *
+-+-+-+-+-+
| | | * | |
+-+-+-+-+-+
| * * | | |
+-+-+-+-+-+
`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			fb := test.chord.FretBoard()
			str := fb.String()
			if "\n"+str != test.expected {
				t.Errorf(
					"expected:%s\ngot:\n%s\n",
					test.expected,
					str,
				)
			}
		})
	}
}
