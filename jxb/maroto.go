package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/wsand02/jxbscare/parser"
)

func (tsl *TreeShapeListener) EnterMaroto(ctx *parser.MarotoContext) {
	tsl.ColsToAdd = []core.Col{}
	tsl.InsertCounter = 0
}

func (tsl *TreeShapeListener) ExitMaroto(ctx *parser.MarotoContext) {
	tsl.PPdf.AddRow(10, tsl.ColsToAdd...)
	tsl.InsertCounter = 0
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
		tsl.ColsToAdd = append(tsl.ColsToAdd, text.NewCol(tsl.GetMyWidth(), idk.Data))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}

func (tsl *TreeShapeListener) AddStuffRecRowDry(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		// tsl.InsertCounter++
		tsl.AddStuffRecRowDry(idk.Children)
	}
	tsl.InsertCounter++
}

func (tsl *TreeShapeListener) AddStuffRecDirect(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.PPdf.AddAutoRow(text.NewCol(12, idk.Data))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}
