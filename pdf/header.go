package pdf

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/wsand02/jxbscare/cv"
)

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
func getHeader(cuvi cv.CV) []core.Row {
	rows := []core.Row{
		row.New(32).Add(
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
			})),
	}

	rows = appendHR(rows)
	return rows
}
