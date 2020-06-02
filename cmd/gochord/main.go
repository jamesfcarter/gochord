package main

import (
	"fmt"
	"os"

	"github.com/jamesfcarter/gochord/pkg/gochord"
)

func main() {
	chord := gochord.NewChord(os.Args[1], os.Args[2])
	fmt.Println(chord)
	f, _ := os.Create("out.svg")
	defer f.Close()
	chord.SVG(500, 700, f)
}
