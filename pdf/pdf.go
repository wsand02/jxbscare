package pdf

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
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
	CVData    map[string]Aboowlock
	PPdf      core.Maroto
	LastRow   core.Row
	LastCol   core.Col
	ColsToAdd []core.Col
}

func NewTreeShapeListener(cfg *entity.Config) *TreeShapeListener {
	tsl := new(TreeShapeListener)
	tsl.CVData = make(map[string]Aboowlock)
	tsl.PPdf = maroto.New(cfg)
	return tsl
}

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
	parent := findEnclosingBlock(ctx)
	if parent == "row" {
		tsl.InsertIntoRow(ctx)
	} else {
		tsl.InsertDirectly(ctx)
	}
}

func (tsl *TreeShapeListener) InsertDirectly(ctx *parser.InsertContext) {
	if ctx.KEYWORD() != nil {
		tsl.PPdf.AddAutoRow(text.NewCol(12, tsl.CVData[ctx.KEYWORD().GetText()].Data))
	} else if ctx.STRING() != nil {
		tsl.AddStuffRecDirect(tsl.CVData[ctx.STRING().GetText()].Children)
	}
}

// todo count all children from parent aka count all children in row...

func (tsl *TreeShapeListener) InsertIntoRow(ctx *parser.InsertContext) {
	if ctx.KEYWORD() != nil {
		tsl.ColsToAdd = append(tsl.ColsToAdd, text.NewCol(12, tsl.CVData[ctx.KEYWORD().GetText()].Data))
	} else if ctx.STRING() != nil {
		tsl.AddStuffRecRow(tsl.CVData[ctx.STRING().GetText()].Children)
	}
}

func (tsl *TreeShapeListener) EnterMaroto(ctx *parser.MarotoContext) {
	tsl.ColsToAdd = []core.Col{}
}

func (tsl *TreeShapeListener) ExitMaroto(ctx *parser.MarotoContext) {
	tsl.PPdf.AddRow(10, tsl.ColsToAdd...)
}

// func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {

// 	tsl.PPdf.AddRow(10,
// 		text.NewCol(10, ctx.GetText()))
// }

func (tsl *TreeShapeListener) AddStuffRecRow(CVData map[string]Aboowlock) {
	width := 12
	numCols := len(CVData)
	if numCols > 0 {
		width = 12 / numCols
	}
	for _, idk := range CVData {
		tsl.ColsToAdd = append(tsl.ColsToAdd, text.NewCol(width, idk.Data))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}

func (tsl *TreeShapeListener) AddStuffRecDirect(CVData map[string]Aboowlock) {
	for _, idk := range CVData {
		tsl.PPdf.AddAutoRow(text.NewCol(12, idk.Data))
		fmt.Println(idk.MarotoNodeType)
		tsl.AddStuffRecRow(idk.Children)
	}
}

func Laboutonmaxxadlatte(filename string, paperSize string) {
	// fmt.Printf("%s", os.Args[1])
	input, _ := antlr.NewFileStream(filename)

	lexer := parser.NewJXBLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJXBParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Document()

	paper := pagesize.Type(strings.ToLower(paperSize))
	fmt.Println(paperSize)

	cfg := config.NewBuilder().WithPageSize(paper).Build()
	listener := NewTreeShapeListener(cfg)
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
