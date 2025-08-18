package jxb

import (
	"fmt"
	"strings"

	"github.com/wsand02/jxbscare/parser"
)

func (tsl *TreeShapeListener) EnterAssignment(ctx *parser.AssignmentContext) {
	ass := ctx.KEYWORD().GetText()
	fmt.Println(ass)
	ab := Aboowlock{
		MarotoNodeType: ass,
	}
	sb := strings.Builder{}
	for idx, val := range ctx.AllSTRING() {
		fmt.Printf(" %s ", val)
		sb.WriteString(val.GetText())
		if idx < len(ctx.AllSTRING())-1 {
			sb.WriteString(" ")
		}
	}
	ab.Data = sb.String()
	parent := findEnclosingBlock(ctx)
	if parent == "global" {
		tsl.CVData[ass] = ab
		return
	}
	parentBlock, exists := tsl.CVData[parent]
	if !exists {
		fmt.Printf("Parent block '%s' not found\n", parent)
		return
	}
	if parentBlock.Children == nil {
		parentBlock.Children = make(map[string]Aboowlock)
	}
	parentBlock.Children[ass] = ab
	tsl.CVData[parent] = parentBlock
}
