package main

import (
	"github.com/wsand02/jxbscare/pdf"
)

const (
	paperSizeDesc      string  = "The size of the paper the pdf will be rendered with."
	defaultPaperSize   string  = "a4"
	lineSpacingDesc    string  = "Space between lines for insert statement rendering."
	defaultLineSpacing float64 = 4.0
)

func main() {
	pdf.GeneratePDF()
}
