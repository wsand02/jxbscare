package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/wsand02/jxbscare/parser"
)

const (
	HR_TOP_MARGIN = 3
	HR_BOT_MARGIN = 6
	HR_WIDTH_PER  = 100
	BR_Y_SIZE     = 10
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
	tsl.PPdf.AddRow(HR_TOP_MARGIN, col.New(MAROTO_MAX_WIDTH))
	tsl.PPdf.AddRow(HR_BOT_MARGIN, line.NewCol(MAROTO_MAX_WIDTH, props.Line{
		SizePercent: HR_WIDTH_PER,
	}))
}

func (tsl *TreeShapeListener) InsertBR() {
	tsl.PPdf.AddRow(BR_Y_SIZE, col.New(MAROTO_MAX_WIDTH))
}

func (tsl *TreeShapeListener) InsertDirectly(ctx *parser.InsertContext) {
	if ctx.KEYWORD() != nil {
		if tsl.CVData[ctx.KEYWORD().GetText()].MarotoNodeType == "selfie" {
			tsl.PPdf.AddAutoRow(image.NewFromFileCol(3, tsl.CVData[ctx.KEYWORD().GetText()].Data, props.Rect{
				Center:  false,
				Percent: 80,
			}))
		} else {
			tsl.PPdf.AddAutoRow(text.NewCol(MAROTO_MAX_WIDTH, tsl.CVData[ctx.KEYWORD().GetText()].Data, props.Text{
				Align: CtxToAlign(ctx),
			}))
		}

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
