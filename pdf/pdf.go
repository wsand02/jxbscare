package pdf

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/wsand02/jxbscare/cv"
	"github.com/wsand02/jxbscare/jxb"
	"github.com/wsand02/jxbscare/parser"
)

const (
	LEFT_MARGIN  = 10
	RIGHT_MARGIN = 10
	TOP_MARGIN   = 15
)

// träd?
// äta träd https://dota2.fandom.com/wiki/Tango

func Laboutonmaxxadlatte(filename string, paperSize string, lineSpacing float64) {
	input, _ := antlr.NewFileStream(filename)

	lexer := parser.NewJXBLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJXBParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Document()

	paper := pagesize.Type(strings.ToLower(paperSize))
	fmt.Println(paperSize)

	cfg := config.NewBuilder().
		WithLeftMargin(LEFT_MARGIN).
		WithRightMargin(RIGHT_MARGIN).
		WithTopMargin(TOP_MARGIN).
		WithPageSize(paper).Build()
	listener := jxb.NewTreeShapeListener(cfg, lineSpacing)
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

func GeneratePDF() {
	cfg := config.NewBuilder().
		WithLeftMargin(LEFT_MARGIN).
		WithRightMargin(RIGHT_MARGIN).
		WithTopMargin(TOP_MARGIN).
		WithPageSize(pagesize.A4).Build()
	pdf := maroto.New(cfg)
	cuvi := cv.ParseCV("cv.toml")
	rows := []core.Row{}
	rows = appendAllTheThings(rows, cuvi)
	pdf.AddRows(rows...)
	document, err := pdf.Generate()
	if err != nil {
		log.Fatal(err)
	}
	err = document.Save("cv.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
