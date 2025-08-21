package pdf

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/wsand02/jxbscare/cv"
)

func getEducation(cuvi cv.CV) []core.Row {
	rows := []core.Row{
		row.New(6).Add(col.New(12).Add(
			text.New("Education", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
			}),
		)),
	}
	for _, edu := range cuvi.Education {
		rows = append(rows,
			row.New(4).Add(
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
			),
			row.New(4).Add(text.NewCol(12, edu.School, props.Text{
				Align: align.Left,
				Style: fontstyle.Italic,
				Size:  7,
			})),
			row.New(5).Add(text.NewCol(12, edu.Desc, props.Text{
				Align: align.Left,
				Size:  8,
			})),
		)
	}
	rows = appendHR(rows)
	return rows
}
