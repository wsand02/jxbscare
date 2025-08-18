package pdf

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/wsand02/jxbscare/jxb"
	"github.com/wsand02/jxbscare/parser"
)

// träd?

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
	listener := jxb.NewTreeShapeListener(cfg)
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
