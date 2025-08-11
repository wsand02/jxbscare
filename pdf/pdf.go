package pdf

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/wsand02/jxbscare/parser"
)

type TreeShapeListener struct {
	*parser.BaseJXBListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func Laboutonmaxxadlatte() {
	// fmt.Printf("%s", os.Args[1])
	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := parser.NewJXBLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJXBParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Document()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
