package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wsand02/jxbscare/pdf"
)

const (
	paperSizeDesc      string  = "The size of the paper the pdf will be rendered with."
	defaultPaperSize   string  = "a4"
	lineSpacingDesc    string  = "Space between lines for insert statement rendering."
	defaultLineSpacing float64 = 4.0
)

func main() {
	paperSize := flag.String("paperSize", defaultPaperSize, paperSizeDesc)
	flag.StringVar(paperSize, "p", defaultPaperSize, paperSizeDesc)
	lineSpacing := flag.Float64("lineSpacing", defaultLineSpacing, lineSpacingDesc)
	flag.Parse()

	inputJxb := flag.Arg(0)
	if len(inputJxb) == 0 {
		log.Fatal("No input provided")
	}
	fmt.Println(*paperSize)
	pdf.Laboutonmaxxadlatte(inputJxb, *paperSize, *lineSpacing)
}
