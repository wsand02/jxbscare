package pdf

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/wsand02/jxbscare/parser"
)

// träd?
type Aboowlock struct {
	MarotoNodeType string
	Data           string
	Children       map[string]Aboowlock
	// OGCtx          *antlr.ParserRuleContext
}

// func kukFunktion(ctx antlr.ParserRuleContext) Aboowlock {
// 	abowlock := Aboowlock{}
// 	abowlock.OGCtx = &ctx
// 	abowlock.Children = &[]Aboowlock{}
// 	return abowlock
// }

// då måste jag kunna söka det...

// this will be slow but whatever

type TreeShapeListener struct {
	*parser.BaseJXBListener
	CVData map[string]Aboowlock
	PPdf   core.Maroto
}

func NewTreeShapeListener() *TreeShapeListener {
	tsl := new(TreeShapeListener)
	tsl.CVData = make(map[string]Aboowlock)
	tsl.PPdf = maroto.New()
	return tsl
}

func findEnclosingBlock(ctx antlr.ParserRuleContext) string {
	for p := ctx.GetParent(); p != nil; p = p.GetParent() {
		if blockCtx, ok := p.(*parser.BlockContext); ok {
			return blockCtx.STRING(0).GetText() // Return the name of the block
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

func (tsl *TreeShapeListener) EnterAssignment(ctx *parser.AssignmentContext) {
	ass := ctx.KEYWORD().GetText()
	fmt.Println(ass)
	ab := Aboowlock{
		MarotoNodeType: ass,
	}
	sb := strings.Builder{}
	for idx, val := range ctx.AllSTRING() {
		fmt.Printf(" %s ", val)
		sb.WriteString(val.GetText())
		if idx < len(ctx.AllSTRING())-1 {
			sb.WriteString(" ")
		}
	}
	ab.Data = sb.String()
	parent := findEnclosingBlock(ctx)
	if parent == "global" {
		tsl.CVData[ass] = ab
		return
	}
	parentBlock, exists := tsl.CVData[parent]
	if !exists {
		fmt.Printf("Parent block '%s' not found\n", parent)
		return
	}
	if parentBlock.Children == nil {
		parentBlock.Children = make(map[string]Aboowlock)
	}
	parentBlock.Children[ass] = ab
	tsl.CVData[parent] = parentBlock
}

func (tsl *TreeShapeListener) EnterInsert(ctx *parser.InsertContext) {
	fmt.Println(ctx.GetText())
	if ctx.STRING() != nil {
		tsl.AddStuff(tsl.CVData[ctx.STRING().GetText()].Children)
	} else if ctx.KEYWORD() != nil {
		tsl.AddStuff(tsl.CVData)
	}
}

// func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {

// 	tsl.PPdf.AddRow(10,
// 		text.NewCol(10, ctx.GetText()))
// }

func (tsl *TreeShapeListener) AddStuff(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.PPdf.AddAutoRow(text.NewCol(10, idk.Data))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuff(idk.Children)
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
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	document, err := listener.PPdf.Generate()
	if err != nil {
		fmt.Printf("Error generating PDF: %s\n", err)
	}
	err = document.Save("output.pdf")
	if err != nil {
		fmt.Println("whoops")
	}
}
