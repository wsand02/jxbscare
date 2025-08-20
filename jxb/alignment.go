package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
)

type Alignment int

const (
	Left Alignment = iota
	Right
	Center
	Top
	Bottom
)

// yeah idfk
func AlignmentToProp(a Alignment) align.Type {
	switch a {
	case Left:
		fmt.Println("l")
		return align.Left
	case Right:
		fmt.Println("r")
		return align.Right
	case Center:
		fmt.Println("c")
		return align.Center
	case Top:
		fmt.Println("t")
		return align.Top
	case Bottom:
		fmt.Println("b")
		return align.Bottom
	default:
		fmt.Println("d")
		// default, redudant statement fight me
		return align.Left
	}
}
