package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/wsand02/jxbscare/parser"
)

func (tsl *TreeShapeListener) EnterInsert(ctx *parser.InsertContext) {
	fmt.Println(ctx.GetText())
	p := findEnclosingBlock(ctx)
	if p != "global" {
		return // handled elsewhere in EnterMaroto
	}
	if ctx.HR() != nil {
		tsl.InsertHR(ctx)
	}
	tsl.InsertDirectly(ctx)
}

func (tsl *TreeShapeListener) InsertHR(ctx *parser.InsertContext) {
	tsl.PPdf.AddRow(10, col.New(12))
	tsl.PPdf.AddRow(20, line.NewCol(12))
}

func (tsl *TreeShapeListener) InsertDirectly(ctx *parser.InsertContext) {
	if ctx.KEYWORD() != nil {
		tsl.PPdf.AddAutoRow(text.NewCol(12, tsl.CVData[ctx.KEYWORD().GetText()].Data, props.Text{
			Align: CtxToAlign(ctx),
		}))
	} else if ctx.STRING() != nil {
		tsl.AddStuffRecDirect(ctx, tsl.CVData[ctx.STRING().GetText()].Children)
	}
}

// func (tsl *TreeShapeListener) GetMyWidth() int {
// 	width := 12
// 	numCols := tsl.InsertCounter
// 	if numCols > 0 {
// 		width = 12 / numCols
// 	}
// 	fmt.Printf("Width: 12 / %d = %d \n", numCols, width)
// 	return width
// }

func CalcWidth(col int) int {
	width := MAROTO_MAX_WIDTH
	if col > 0 {
		width = MAROTO_MAX_WIDTH / col
	}
	fmt.Printf("Width: %d / %d = %d \n", MAROTO_MAX_WIDTH, col, width)
	return width
}

// todo count all children from parent aka count all children in row...

// func (tsl *TreeShapeListener) InsertIntoRow(ctx *parser.InsertContext) {
// 	if ctx.KEYWORD() != nil {
// 		tsl.ColsToAdd = append(tsl.ColsToAdd, text.NewCol(tsl.GetMyWidth(), tsl.CVData[ctx.KEYWORD().GetText()].Data, AlignmentToProp(tsl.ColAlign)))
// 	} else if ctx.STRING() != nil {
// 		tsl.AddStuffRecRow(tsl.CVData[ctx.STRING().GetText()].Children)
// 	}
// }

// func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {

// 	tsl.PPdf.AddRow(10,
// 		text.NewCol(10, ctx.GetText()))
// }
