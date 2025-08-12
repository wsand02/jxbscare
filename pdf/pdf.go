package pdf

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/wsand02/jxbscare/parser"
)

type Aboowlock struct {
	Data  string
	Child *Aboowlock
}

type TreeShapeListener struct {
	*parser.BaseJXBListener
	CVData map[string]string
	Blocks []Aboowlock
	PPdf   core.Maroto
}

func NewTreeShapeListener() *TreeShapeListener {
	tsl := new(TreeShapeListener)
	tsl.CVData = make(map[string]string)
	tsl.Blocks = []Aboowlock{}
	tsl.PPdf = maroto.New()
	return tsl
}

func (tsl *TreeShapeListener) EnterAssignment(ctx *parser.AssignmentContext) {
	fmt.Println("IN ENTER ASSIGNMENT")
	ass := ctx.KEYWORD().GetText()
	ctx.GetInvokingState()
	fmt.Println(ass)
	sb := strings.Builder{}
	for idx, val := range ctx.AllSTRING() {
		fmt.Printf(" %s ", val)
		sb.WriteString(val.GetText())
		if idx < len(ctx.AllSTRING())-1 {
			sb.WriteString(" ")
		}
	}
	tsl.CVData[ass] = sb.String()
	fmt.Println()
	fmt.Println(tsl.CVData[ass])
}

func (tsl *TreeShapeListener) EnterBlock(ctx *parser.BlockContext) {
	fmt.Printf("IN BLOCK %s\n", ctx.STRING(0).GetText())
	fmt.Println("BLOCK CONTENTS: ")
	if len(ctx.AllStatement()) <= 0 {
		return
	}
	for _, val := range ctx.AllStatement() {
		fmt.Println(val.GetText())
		if val.Block() != nil {
			fmt.Println("BLOCK IN BLOCK WOAH")
		}
	}
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
