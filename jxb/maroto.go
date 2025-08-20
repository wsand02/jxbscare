package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/wsand02/jxbscare/parser"
)

func (tsl *TreeShapeListener) EnterMaroto(ctx *parser.MarotoContext) {
	tsl.ColsToAdd = []core.Col{}
	colAmount := len(ctx.AllMarcol())
	width := CalcWidth(colAmount)
	for _, xcol := range ctx.AllMarcol() {
		tsl.ComponentsToAdd = []core.Component{}
		for _, ins := range xcol.AllInsert() {
			if ins.KEYWORD() != nil {
				tsl.ComponentsToAdd = append(tsl.ComponentsToAdd, text.New(tsl.CVData[ins.KEYWORD().GetText()].Data, AlignmentToProp(tsl.ColAlign)))
			} else if ins.STRING() != nil {
				tsl.MarotoInsertRec(tsl.CVData[ins.STRING().GetText()].Children)
			}

		}
		tsl.ColsToAdd = append(tsl.ColsToAdd, col.New(width).Add(tsl.ComponentsToAdd...))
	}
	tsl.PPdf.AddAutoRow(tsl.ColsToAdd...)
}

// differs from surface level AddStufRecRow
func (tsl *TreeShapeListener) MarotoInsertRec(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.ComponentsToAdd = append(tsl.ComponentsToAdd, text.New(idk.Data, AlignmentToProp(tsl.ColAlign)))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}

func (tsl *TreeShapeListener) ExitMaroto(ctx *parser.MarotoContext) {
	tsl.PPdf.AddRow(10, tsl.ColsToAdd...)
	tsl.InsertCounter = 0
	tsl.ColsToAdd = []core.Col{}
	tsl.ComponentsToAdd = []core.Component{}
}

// func (tsl *TreeShapeListener) GetChildCountFromParent(ctx *parser.MarotoContext) {
// 	for p := ctx.GetParent(); p != nil; p = p.GetParent() {
// 		if _, ok := p.(*parser.MarotoContext); ok {

// 		}
// 	}
// 	return "global"
// }

func (tsl *TreeShapeListener) AddStuffRecRow(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.ColsToAdd = append(tsl.ColsToAdd, text.NewCol(tsl.GetMyWidth(), idk.Data, AlignmentToProp(tsl.ColAlign)))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}

func (tsl *TreeShapeListener) AddStuffRecDirect(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.PPdf.AddAutoRow(text.NewCol(12, idk.Data, AlignmentToProp(tsl.ColAlign)))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}
