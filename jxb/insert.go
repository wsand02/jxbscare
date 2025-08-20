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
		tsl.InsertHR()
		return
	}
	if ctx.BR() != nil {
		tsl.InsertBR()
		return
	}
	tsl.InsertDirectly(ctx)
}

func (tsl *TreeShapeListener) InsertHR() {
	tsl.PPdf.AddRow(10, col.New(MAROTO_MAX_WIDTH))
	tsl.PPdf.AddRow(20, line.NewCol(MAROTO_MAX_WIDTH))
}

func (tsl *TreeShapeListener) InsertBR() {
	tsl.PPdf.AddRow(10, col.New(MAROTO_MAX_WIDTH))
}

func (tsl *TreeShapeListener) InsertDirectly(ctx *parser.InsertContext) {
	if ctx.KEYWORD() != nil {
		tsl.PPdf.AddAutoRow(text.NewCol(MAROTO_MAX_WIDTH, tsl.CVData[ctx.KEYWORD().GetText()].Data, props.Text{
			Align: CtxToAlign(ctx),
		}))
	} else if ctx.STRING() != nil {
		tsl.AddStuffRecDirect(ctx, tsl.CVData[ctx.STRING().GetText()].Children)
	}
}

func CalcWidth(col int) int {
	width := MAROTO_MAX_WIDTH
	if col > 0 {
		width = MAROTO_MAX_WIDTH / col
	}
	fmt.Printf("Width: %d / %d = %d \n", MAROTO_MAX_WIDTH, col, width)
	return width
}
