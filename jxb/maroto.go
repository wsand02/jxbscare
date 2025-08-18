package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/wsand02/jxbscare/parser"
)

func (tsl *TreeShapeListener) EnterMaroto(ctx *parser.MarotoContext) {
	tsl.ColsToAdd = []core.Col{}
}

func (tsl *TreeShapeListener) ExitMaroto(ctx *parser.MarotoContext) {
	tsl.PPdf.AddRow(10, tsl.ColsToAdd...)
}

func (tsl *TreeShapeListener) AddStuffRecRow(CVData map[string]Aboowlock) {
	width := 12
	numCols := len(CVData)
	if numCols > 0 {
		width = 12 / numCols
	}
	for _, idk := range CVData {
		tsl.ColsToAdd = append(tsl.ColsToAdd, text.NewCol(width, idk.Data))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}

func (tsl *TreeShapeListener) AddStuffRecDirect(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.PPdf.AddAutoRow(text.NewCol(12, idk.Data))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}
