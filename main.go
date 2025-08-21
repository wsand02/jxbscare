package main

import (
	"flag"
	"log"

	"github.com/wsand02/jxbscare/pdf"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("Usage: jxbscare <input> <output file path>")
	}
	input := flag.Arg(0)
	output := flag.Arg(1)
	pdf.GeneratePDF(input, output)
}
