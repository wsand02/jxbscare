package pdf

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
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
	AppendHeader(pdf, cuvi)
	AppendWorkExperience(pdf, cuvi)
	AppendPublications(pdf, cuvi)
	AppendProjects(pdf, cuvi)
	AppendEducation(pdf, cuvi)
	AppendTechnicalSkills(pdf, cuvi)
	AppendLanguages(pdf, cuvi)
	AppendReferences(pdf, cuvi)
	document, err := pdf.Generate()
	if err != nil {
		log.Fatal(err)
	}
	err = document.Save("cv.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

/*
right row
selfie idkf.png
name Anders Andersson

left row
address Gatan 200F, 66633, Atlanta
phone 46777777777
email example@example.com
website example.com
linkedin
github

hr
*/
func AppendHeader(m core.Maroto, cuvi cv.CV) {
	m.AddRow(32,
		col.New(6).Add(
			text.New(fmt.Sprintf("CV: %s", cuvi.AboutMe.Name), props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
			}),
			text.New(fmt.Sprintf("Address: %s", cuvi.AboutMe.Address), props.Text{
				Top:   8,
				Size:  8,
				Align: align.Left,
			}),
			text.New(fmt.Sprintf("Phone: %s", cuvi.AboutMe.Phone), props.Text{
				Top:   12,
				Size:  8,
				Align: align.Left,
			}),
			text.New(fmt.Sprintf("Email: %s", cuvi.AboutMe.Email), props.Text{
				Top:   16,
				Size:  8,
				Align: align.Left,
			}),
			text.New(fmt.Sprintf("Website: %s", cuvi.AboutMe.Website), props.Text{
				Top:   20,
				Size:  8,
				Align: align.Left,
			}),
			text.New(fmt.Sprintf("LinkedIn: %s", cuvi.AboutMe.LinkedIn), props.Text{
				Top:   24,
				Size:  8,
				Align: align.Left,
			}),
			text.New(fmt.Sprintf("Github: %s", cuvi.AboutMe.Github), props.Text{
				Top:   28,
				Size:  8,
				Align: align.Left,
			})),
		col.New(4),
		image.NewFromFileCol(2, cuvi.AboutMe.Selfie, props.Rect{
			Center:  false,
			Percent: 100,
		}))

	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}

/*
job:
title
company
start
end
location
desc
*/
func AppendWorkExperience(m core.Maroto, cuvi cv.CV) {
	m.AddRow(6, col.New(12).Add(
		text.New("Work Experience", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
		}),
	))
	for _, work := range cuvi.WorkExperience {
		m.AddRow(4,
			text.NewCol(10, fmt.Sprintf("%s, %s", work.Title, work.Company), props.Text{
				Align: align.Left,
				Size:  8,
				Style: fontstyle.Bold,
			}),
			text.NewCol(2, fmt.Sprintf("%s - %s", work.Start, work.End), props.Text{
				Align: align.Right,
				Size:  8,
				Style: fontstyle.Italic,
			}),
		)
		m.AddRow(4, text.NewCol(12, work.Location, props.Text{
			Align: align.Left,
			Style: fontstyle.Italic,
			Size:  7,
		}))
		m.AddRow(6, text.NewCol(12, work.Description, props.Text{
			Align: align.Left,
			Size:  8,
		}))
	}
	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}

func AppendPublications(m core.Maroto, cuvi cv.CV) {
	m.AddRow(6, col.New(12).Add(
		text.New("Publications", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
		}),
	))
	for _, pub := range cuvi.Publications {
		m.AddRow(4,
			text.NewCol(10, pub.Title, props.Text{
				Align: align.Left,
				Size:  8,
				Style: fontstyle.Bold,
			}),
			text.NewCol(2, fmt.Sprintf("%d", pub.Year), props.Text{
				Align: align.Right,
				Size:  8,
				Style: fontstyle.Italic,
			}),
		)
		m.AddRow(4, text.NewCol(12, pub.Summary, props.Text{
			Align: align.Left,
			Size:  8,
		}))
	}
	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}

func AppendProjects(m core.Maroto, cuvi cv.CV) {
	m.AddRow(6, col.New(12).Add(
		text.New("Projects", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
		}),
	))
	for _, project := range cuvi.Projects {
		m.AddRow(4,
			text.NewCol(10, project.Name, props.Text{
				Align: align.Left,
				Size:  8,
				Style: fontstyle.Bold,
			}),
			text.NewCol(2, fmt.Sprintf("%s - %s", project.Start, project.End), props.Text{
				Align: align.Right,
				Size:  8,
				Style: fontstyle.Italic,
			}),
		)
		m.AddRow(4, text.NewCol(12, project.Description, props.Text{
			Align: align.Left,
			Size:  8,
		}))
		if len(project.Technologies) > 0 {
			techs := strings.Join(project.Technologies, ", ")
			m.AddRow(5, text.NewCol(12, fmt.Sprintf("Technologies: %s", techs), props.Text{
				Align: align.Left,
				Size:  7,
				Style: fontstyle.Italic,
			}))
		} else {
			m.AddRow(1, col.New(12).Add())
		}
	}
	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}

func AppendEducation(m core.Maroto, cuvi cv.CV) {
	m.AddRow(6, col.New(12).Add(
		text.New("Education", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
		}),
	))
	for _, edu := range cuvi.Education {
		m.AddRow(4,
			text.NewCol(10, edu.Name, props.Text{
				Align: align.Left,
				Size:  8,
				Style: fontstyle.Bold,
			}),
			text.NewCol(2, fmt.Sprintf("%s - %s", edu.Start, edu.End), props.Text{
				Align: align.Right,
				Size:  8,
				Style: fontstyle.Italic,
			}),
		)
		m.AddRow(4, text.NewCol(12, edu.School, props.Text{
			Align: align.Left,
			Style: fontstyle.Italic,
			Size:  7,
		}))
		m.AddRow(5, text.NewCol(12, edu.Desc, props.Text{
			Align: align.Left,
			Size:  8,
		}))
	}
	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}

func AppendTechnicalSkills(m core.Maroto, cuvi cv.CV) {
	m.AddRow(6, col.New(12).Add(
		text.New("Technical Skills", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
		}),
	))
	m.AddRow(4,
		text.NewCol(3, "Programming Languages: ", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  8,
		}),
		text.NewCol(11, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Language), ", "), props.Text{
			Align: align.Left,
			Size:  8,
		}),
	)
	m.AddRow(4,
		text.NewCol(3, "Frameworks: ", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  8,
		}),
		text.NewCol(11, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Framework), ", "), props.Text{
			Align: align.Left,
			Size:  8,
		}),
	)
	m.AddRow(4,
		text.NewCol(3, "Tools: ", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  8,
		}),
		text.NewCol(11, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Tool), ", "), props.Text{
			Align: align.Left,
			Size:  8,
		}),
	)
	m.AddRow(4,
		text.NewCol(3, "Other Skills: ", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  8,
		}),
		text.NewCol(11, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Other), ", "), props.Text{
			Align: align.Left,
			Size:  8,
		}),
	)
	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}

func AppendLanguages(m core.Maroto, cuvi cv.CV) {
	m.AddRow(6, col.New(12).Add(
		text.New("Languages", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
		}),
	))
	for _, lang := range cuvi.Languages {
		m.AddRow(4,
			text.NewCol(10, lang.Name, props.Text{
				Align: align.Left,
				Size:  8,
				Style: fontstyle.Bold,
			}),
			text.NewCol(2, lang.Summary, props.Text{
				Align: align.Right,
				Size:  8,
				Style: fontstyle.Italic,
			}),
		)
	}
	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}

func AppendReferences(m core.Maroto, cuvi cv.CV) {
	m.AddAutoRow(col.New(12).Add(
		text.New("References", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
		}),
	))
	for _, ref := range cuvi.References {
		m.AddRow(4,
			text.NewCol(10, ref.Name, props.Text{
				Align: align.Left,
				Size:  8,
				Style: fontstyle.Bold,
			}),
			text.NewCol(2, ref.Company, props.Text{
				Align: align.Right,
				Size:  8,
				Style: fontstyle.Italic,
			}),
		)
		m.AddRow(4, text.NewCol(12, fmt.Sprintf("Phone: %s, Email: %s", ref.Phone, ref.Email), props.Text{
			Align: align.Left,
			Size:  8,
		}))
	}
	m.AddRow(4, col.New(12).Add())
	m.AddRow(6, line.NewCol(12, props.Line{
		SizePercent: 100,
	}))
}
