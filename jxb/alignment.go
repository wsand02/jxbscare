package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/wsand02/jxbscare/parser"
)

type Alignment int

const (
	Left Alignment = iota
	Right
	Center
	Top
	Bottom
)

// fight me
func AlignmentToAlignment(ctx *parser.InsertContext) Alignment {
	if ctx.ALIGNMENT() != nil {
		a := ctx.ALIGNMENT().GetText()
		fmt.Println(a)
		switch a {
		case "left":
			return Left
		case "right":
			return Right
		case "center":
			return Center
		case "top":
			return Top
		case "bottom":
			return Bottom
		default:
			return Left
		}
	}
	return Left
}

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
