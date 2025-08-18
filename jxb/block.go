package jxb

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/wsand02/jxbscare/parser"
)

func findEnclosingBlock(ctx antlr.ParserRuleContext) string {
	for p := ctx.GetParent(); p != nil; p = p.GetParent() {
		if blockCtx, ok := p.(*parser.BlockContext); ok {
			return blockCtx.STRING(0).GetText() // Return the name of the block
		}
		if _, ok := p.(*parser.MarotoContext); ok {
			return "row"
		}
	}
	return "global"
}

func (tsl *TreeShapeListener) EnterBlock(ctx *parser.BlockContext) {
	ass := ctx.STRING(0).GetText()
	ab := Aboowlock{
		MarotoNodeType: "block",
	}
	// i framtiden ska då alltså dess arguments vara i data?
	tsl.CVData[ass] = ab
}
