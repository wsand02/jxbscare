package jxb

import (
	"fmt"

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
