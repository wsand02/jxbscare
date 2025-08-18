// Code generated from JXB.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JXB

import "github.com/antlr4-go/antlr/v4"

// BaseJXBListener is a complete listener for a parse tree produced by JXBParser.
type BaseJXBListener struct{}

var _ JXBListener = &BaseJXBListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJXBListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJXBListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJXBListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJXBListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseJXBListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseJXBListener) ExitDocument(ctx *DocumentContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJXBListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJXBListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseJXBListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseJXBListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJXBListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJXBListener) ExitBlock(ctx *BlockContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseJXBListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseJXBListener) ExitInsert(ctx *InsertContext) {}

// EnterMaroto is called when production maroto is entered.
func (s *BaseJXBListener) EnterMaroto(ctx *MarotoContext) {}

// ExitMaroto is called when production maroto is exited.
func (s *BaseJXBListener) ExitMaroto(ctx *MarotoContext) {}
