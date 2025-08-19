package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Alignment int

const (
	Left Alignment = iota
	Right
	Center
	Top
	Bottom
)

func AlignmentToProp(a Alignment) props.Text {
	switch a {
	case Left:
		fmt.Println("l")
		return props.Text{
			Align: align.Left,
		}
	case Right:
		fmt.Println("r")
		return props.Text{
			Align: align.Right,
		}
	case Center:
		fmt.Println("c")
		return props.Text{
			Align: align.Center,
		}
	case Top:
		fmt.Println("t")
		return props.Text{
			Align: align.Top,
		}
	case Bottom:
		fmt.Println("b")
		return props.Text{
			Align: align.Bottom,
		}
	default:
		fmt.Println("d")
		// default, redudant statement fight me
		return props.Text{
			Align: align.Left,
		}
	}
}
