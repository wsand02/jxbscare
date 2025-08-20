package jxb

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/wsand02/jxbscare/parser"
)

type Aboowlock struct {
	MarotoNodeType string
	Data           string
	Children       map[string]Aboowlock
	// OGCtx          *antlr.ParserRuleContext
}

// func kukFunktion(ctx antlr.ParserRuleContext) Aboowlock {
// 	abowlock := Aboowlock{}
// 	abowlock.OGCtx = &ctx
// 	abowlock.Children = &[]Aboowlock{}
// 	return abowlock
// }

// då måste jag kunna söka det...

// this will be slow but whatever

type TreeShapeListener struct {
	*parser.BaseJXBListener
	CVData          map[string]Aboowlock
	PPdf            core.Maroto
	LastRow         core.Row
	ColAlign        Alignment
	ColsToAdd       []core.Col
	ComponentsToAdd []core.Component
	InsertCounter   float64
	LineSpacing     float64
}

func NewTreeShapeListener(cfg *entity.Config, lineSpacing float64) *TreeShapeListener {
	tsl := new(TreeShapeListener)
	tsl.CVData = make(map[string]Aboowlock)
	tsl.PPdf = maroto.New(cfg)
	tsl.ColAlign = Left
	tsl.InsertCounter = 0.0
	tsl.LineSpacing = lineSpacing
	return tsl
}
