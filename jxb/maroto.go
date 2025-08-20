package jxb

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/wsand02/jxbscare/parser"
)

const MAROTO_MAX_WIDTH = 12

func (tsl *TreeShapeListener) EnterMaroto(ctx *parser.MarotoContext) {
	fmt.Println("ENTER MAROTO")
	tsl.ColsToAdd = []core.Col{}
	colAmount := len(ctx.AllMarcol())
	fmt.Printf("\nCols in Maroto: %d\n", colAmount)
	width := CalcWidth(colAmount)
	for _, xcol := range ctx.AllMarcol() {
		tsl.InsertCounter = 0.0
		tsl.ComponentsToAdd = []core.Component{}
		ins := xcol.Insert()
		if ins.KEYWORD() != nil {
			tsl.ComponentsToAdd = append(tsl.ComponentsToAdd, text.New(tsl.CVData[ins.KEYWORD().GetText()].Data, props.Text{
				Align: AlignmentToProp(AlignmentToAlignment(ins.(*parser.InsertContext))),
				Top:   tsl.InsertCounter,
			}))
		} else if ins.STRING() != nil {
			tsl.MarotoInsertRec(ins.(*parser.InsertContext), tsl.CVData[ins.STRING().GetText()].Children)
		}

		tsl.ColsToAdd = append(tsl.ColsToAdd, col.New(width).Add(tsl.ComponentsToAdd...))
	}
	tsl.PPdf.AddAutoRow(tsl.ColsToAdd...)
}

// differs from surface level AddStufRecRow
func (tsl *TreeShapeListener) MarotoInsertRec(ctx *parser.InsertContext, CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.ComponentsToAdd = append(tsl.ComponentsToAdd, text.New(idk.Data, props.Text{
			Align: CtxToAlign(ctx),
			Top:   tsl.InsertCounter,
		}))
		fmt.Println(idk.MarotoNodeType)
		tsl.MarotoInsertRec(ctx, idk.Children)
		tsl.InsertCounter += tsl.LineSpacing
	}
}

func (tsl *TreeShapeListener) ExitMaroto(ctx *parser.MarotoContext) {
	fmt.Println("EXIT MAROTO")
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

func (tsl *TreeShapeListener) AddStuffRecDirect(ctx *parser.InsertContext, CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.PPdf.AddAutoRow(text.NewCol(MAROTO_MAX_WIDTH, idk.Data, props.Text{
			Align: CtxToAlign(ctx),
		}))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecDirect(ctx, idk.Children)
	}
}
