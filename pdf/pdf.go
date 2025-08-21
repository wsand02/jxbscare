package pdf

import (
	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/wsand02/jxbscare/cv"
)

const (
	LEFT_MARGIN  = 10
	RIGHT_MARGIN = 10
	TOP_MARGIN   = 15
)

// träd?
// äta träd https://dota2.fandom.com/wiki/Tango

func initMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithLeftMargin(LEFT_MARGIN).
		WithRightMargin(RIGHT_MARGIN).
		WithTopMargin(TOP_MARGIN).
		WithPageSize(pagesize.A4).Build()
	return maroto.New(cfg)
}

func GeneratePDF(input string, output string) {
	pdf := initMaroto()
	cuvi := cv.ParseCV(input)
	rows := []core.Row{}
	rows = appendAllTheThings(rows, cuvi)
	pdf.AddRows(rows...)
	document, err := pdf.Generate()
	if err != nil {
		log.Fatal(err)
	}
	err = document.Save(output)
	if err != nil {
		log.Fatal(err)
	}
}
