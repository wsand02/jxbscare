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

func getReferences(cuvi cv.CV) []core.Row {
	rows := []core.Row{
		row.New(6).Add(col.New(12).Add(
			text.New("References", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
			}),
		)),
	}
	for _, ref := range cuvi.References {
		rows = append(rows,
			row.New(4).Add(
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
			),
			row.New(4).Add(text.NewCol(12, fmt.Sprintf("Phone: %s, Email: %s", ref.Phone, ref.Email), props.Text{
				Align: align.Left,
				Size:  8,
			})),
		)
	}
	rows = appendHR(rows)
	return rows
}
