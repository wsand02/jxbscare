package pdf

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/wsand02/jxbscare/parser"
)

var CVData map[string]string = make(map[string]string)

type TreeShapeListener struct {
	*parser.BaseJXBListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (tsl *TreeShapeListener) EnterAssignment(ctx *parser.AssignmentContext) {
	ass := ctx.KEYWORD().GetText()
	fmt.Println(ass)
	sb := strings.Builder{}
	for idx, val := range ctx.AllSTRING() {
		fmt.Printf(" %s ", val)
		sb.WriteString(val.GetText())
		if idx < len(ctx.AllSTRING())-1 {
			sb.WriteString(" ")
		}
	}
	CVData[ass] = sb.String()
	fmt.Println()
	fmt.Println(CVData[ass])
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
