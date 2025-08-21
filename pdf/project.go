package pdf

import (
	"fmt"
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

func getProjects(cuvi cv.CV) []core.Row {
	rows := []core.Row{
		row.New(6).Add(col.New(12).Add(
			text.New("Projects", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
			}))),
	}
	for _, project := range cuvi.Projects {
		rows = append(rows,
			row.New(4).Add(
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
			),
			row.New(4).Add(text.NewCol(12, project.Description, props.Text{
				Align: align.Left,
				Size:  8,
			})))

		if len(project.Technologies) > 0 {
			techs := strings.Join(project.Technologies, ", ")
			rows = append(rows, row.New(5).Add(text.NewCol(12, fmt.Sprintf("Technologies: %s", techs), props.Text{
				Align: align.Left,
				Size:  7,
				Style: fontstyle.Italic,
			})))
		} else {
			rows = append(rows, row.New(1).Add(col.New(12).Add()))
		}
	}
	rows = appendHR(rows)
	return rows
}
