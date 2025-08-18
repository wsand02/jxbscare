package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/wsand02/jxbscare/parser"
)

func (tsl *TreeShapeListener) EnterInsert(ctx *parser.InsertContext) {
	fmt.Println(ctx.GetText())
	parent := findEnclosingBlock(ctx)
	if parent == "row" {
		tsl.InsertIntoRow(ctx)
	} else {
		tsl.InsertDirectly(ctx)
	}
}

func (tsl *TreeShapeListener) InsertDirectly(ctx *parser.InsertContext) {
	if ctx.KEYWORD() != nil {
		tsl.PPdf.AddAutoRow(text.NewCol(12, tsl.CVData[ctx.KEYWORD().GetText()].Data))
	} else if ctx.STRING() != nil {
		tsl.AddStuffRecDirect(tsl.CVData[ctx.STRING().GetText()].Children)
	}
}

// todo count all children from parent aka count all children in row...

func (tsl *TreeShapeListener) InsertIntoRow(ctx *parser.InsertContext) {
	if ctx.KEYWORD() != nil {
		tsl.ColsToAdd = append(tsl.ColsToAdd, text.NewCol(12, tsl.CVData[ctx.KEYWORD().GetText()].Data))
	} else if ctx.STRING() != nil {
		tsl.AddStuffRecRow(tsl.CVData[ctx.STRING().GetText()].Children)
	}
}

// func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {

// 	tsl.PPdf.AddRow(10,
// 		text.NewCol(10, ctx.GetText()))
// }
