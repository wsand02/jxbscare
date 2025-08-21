package pdf

import (
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/wsand02/jxbscare/cv"
)

func getTechnicalSkills(cuvi cv.CV) []core.Row {
	rows := []core.Row{
		row.New(6).Add(col.New(12).Add(
			text.New("Technical Skills", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
			}),
		)),
	}
	rows = append(rows,
		row.New(4).Add(
			text.NewCol(3, "Programming Languages: ", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
				Size:  8,
			}),
			text.NewCol(9, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Language), ", "), props.Text{
				Align: align.Left,
				Size:  8,
			}),
		),
		row.New(4).Add(
			text.NewCol(3, "Frameworks: ", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
				Size:  8,
			}),
			text.NewCol(9, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Framework), ", "), props.Text{
				Align: align.Left,
				Size:  8,
			}),
		),
		row.New(4).Add(
			text.NewCol(3, "Tools: ", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
				Size:  8,
			}),
			text.NewCol(9, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Tool), ", "), props.Text{
				Align: align.Left,
				Size:  8,
			}),
		),
		row.New(4).Add(
			text.NewCol(3, "Other Skills: ", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
				Size:  8,
			}),
			text.NewCol(9, strings.Join(cuvi.GetTechnicalSkillsByCategory(cv.Other), ", "), props.Text{
				Align: align.Left,
				Size:  8,
			}),
		),
	)
	rows = appendHR(rows)
	return rows
}
