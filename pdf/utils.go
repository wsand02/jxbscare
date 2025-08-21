package pdf

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/wsand02/jxbscare/cv"
)

// appendHR appends a horizontal rule to the provided rows.
func appendHR(rows []core.Row) []core.Row {
	rows = append(rows,
		row.New(4).Add(col.New(12).Add()),
		row.New(6).Add(line.NewCol(12, props.Line{
			SizePercent: 100,
		})))
	return rows
}

func appendAllTheThings(rows []core.Row, cuvi cv.CV) []core.Row {
	rows = append(rows, getHeader(cuvi)...)
	rows = append(rows, getWorkExperience(cuvi)...)
	rows = append(rows, getPublications(cuvi)...)
	rows = append(rows, getProjects(cuvi)...)
	rows = append(rows, getEducation(cuvi)...)
	rows = append(rows, getTechnicalSkills(cuvi)...)
	rows = append(rows, getLanguages(cuvi)...)
	rows = append(rows, getReferences(cuvi)...)
	return rows
}
