// Code generated from JXB.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JXB

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JXBParser struct {
	*antlr.BaseParser
}

var JXBParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jxbParserInit() {
	staticData := &JXBParserStaticData
	staticData.LiteralNames = []string{
		"", "'.'", "'begin'", "'end'", "'insert'", "", "", "'row'", "'col'",
		"'hr'",
	}
	staticData.SymbolicNames = []string{
		"", "", "BEGIN", "END", "INSERT", "KEYWORD", "ALIGNMENT", "ROW", "COL",
		"HR", "STRING", "WS", "NL",
	}
	staticData.RuleNames = []string{
		"document", "statement", "assignment", "block", "insert", "maroto",
		"marcol",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 71, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 27, 8, 1, 1, 2, 1, 2, 5, 2, 31,
		8, 2, 10, 2, 12, 2, 34, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		42, 8, 3, 10, 3, 12, 3, 45, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 3, 4, 54, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 60, 8, 5, 10, 5, 12, 5,
		63, 9, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6,
		8, 10, 12, 0, 1, 2, 0, 5, 5, 9, 10, 71, 0, 17, 1, 0, 0, 0, 2, 26, 1, 0,
		0, 0, 4, 28, 1, 0, 0, 0, 6, 37, 1, 0, 0, 0, 8, 50, 1, 0, 0, 0, 10, 57,
		1, 0, 0, 0, 12, 67, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 1, 0, 0, 0,
		16, 19, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 20, 1,
		0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 21, 5, 0, 0, 1, 21, 1, 1, 0, 0, 0, 22,
		27, 3, 4, 2, 0, 23, 27, 3, 6, 3, 0, 24, 27, 3, 8, 4, 0, 25, 27, 3, 10,
		5, 0, 26, 22, 1, 0, 0, 0, 26, 23, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 25,
		1, 0, 0, 0, 27, 3, 1, 0, 0, 0, 28, 32, 5, 5, 0, 0, 29, 31, 5, 10, 0, 0,
		30, 29, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1,
		0, 0, 0, 33, 35, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 36, 5, 12, 0, 0, 36,
		5, 1, 0, 0, 0, 37, 38, 5, 2, 0, 0, 38, 39, 5, 10, 0, 0, 39, 43, 5, 12,
		0, 0, 40, 42, 3, 4, 2, 0, 41, 40, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41,
		1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0,
		46, 47, 5, 3, 0, 0, 47, 48, 5, 10, 0, 0, 48, 49, 5, 12, 0, 0, 49, 7, 1,
		0, 0, 0, 50, 51, 5, 4, 0, 0, 51, 53, 7, 0, 0, 0, 52, 54, 5, 6, 0, 0, 53,
		52, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 5, 12,
		0, 0, 56, 9, 1, 0, 0, 0, 57, 61, 5, 7, 0, 0, 58, 60, 3, 12, 6, 0, 59, 58,
		1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 64, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 5, 1, 0, 0, 65, 66, 5,
		12, 0, 0, 66, 11, 1, 0, 0, 0, 67, 68, 5, 8, 0, 0, 68, 69, 3, 8, 4, 0, 69,
		13, 1, 0, 0, 0, 6, 17, 26, 32, 43, 53, 61,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JXBParserInit initializes any static state used to implement JXBParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJXBParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JXBParserInit() {
	staticData := &JXBParserStaticData
	staticData.once.Do(jxbParserInit)
}

// NewJXBParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJXBParser(input antlr.TokenStream) *JXBParser {
	JXBParserInit()
	this := new(JXBParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JXBParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JXB.g4"

	return this
}

// JXBParser tokens.
const (
	JXBParserEOF       = antlr.TokenEOF
	JXBParserT__0      = 1
	JXBParserBEGIN     = 2
	JXBParserEND       = 3
	JXBParserINSERT    = 4
	JXBParserKEYWORD   = 5
	JXBParserALIGNMENT = 6
	JXBParserROW       = 7
	JXBParserCOL       = 8
	JXBParserHR        = 9
	JXBParserSTRING    = 10
	JXBParserWS        = 11
	JXBParserNL        = 12
)

// JXBParser rules.
const (
	JXBParserRULE_document   = 0
	JXBParserRULE_statement  = 1
	JXBParserRULE_assignment = 2
	JXBParserRULE_block      = 3
	JXBParserRULE_insert     = 4
	JXBParserRULE_maroto     = 5
	JXBParserRULE_marcol     = 6
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JXBParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(JXBParserEOF, 0)
}

func (s *DocumentContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *JXBParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JXBParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&180) != 0 {
		{
			p.SetState(14)
			p.Statement()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(JXBParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Block() IBlockContext
	Insert() IInsertContext
	Maroto() IMarotoContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JXBParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) Insert() IInsertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertContext)
}

func (s *StatementContext) Maroto() IMarotoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarotoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMarotoContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *JXBParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JXBParserRULE_statement)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JXBParserKEYWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Assignment()
		}

	case JXBParserBEGIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Block()
		}

	case JXBParserINSERT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Insert()
		}

	case JXBParserROW:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(25)
			p.Maroto()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KEYWORD() antlr.TerminalNode
	NL() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JXBParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(JXBParserKEYWORD, 0)
}

func (s *AssignmentContext) NL() antlr.TerminalNode {
	return s.GetToken(JXBParserNL, 0)
}

func (s *AssignmentContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(JXBParserSTRING)
}

func (s *AssignmentContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(JXBParserSTRING, i)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *JXBParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JXBParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(JXBParserKEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JXBParserSTRING {
		{
			p.SetState(29)
			p.Match(JXBParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Match(JXBParserNL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEGIN() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	END() antlr.TerminalNode
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JXBParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(JXBParserBEGIN, 0)
}

func (s *BlockContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(JXBParserSTRING)
}

func (s *BlockContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(JXBParserSTRING, i)
}

func (s *BlockContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(JXBParserNL)
}

func (s *BlockContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(JXBParserNL, i)
}

func (s *BlockContext) END() antlr.TerminalNode {
	return s.GetToken(JXBParserEND, 0)
}

func (s *BlockContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *JXBParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JXBParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(JXBParserBEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(JXBParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(JXBParserNL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JXBParserKEYWORD {
		{
			p.SetState(40)
			p.Assignment()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(46)
		p.Match(JXBParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(JXBParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(JXBParserNL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInsertContext is an interface to support dynamic dispatch.
type IInsertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSERT() antlr.TerminalNode
	NL() antlr.TerminalNode
	KEYWORD() antlr.TerminalNode
	HR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	ALIGNMENT() antlr.TerminalNode

	// IsInsertContext differentiates from other interfaces.
	IsInsertContext()
}

type InsertContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertContext() *InsertContext {
	var p = new(InsertContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_insert
	return p
}

func InitEmptyInsertContext(p *InsertContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_insert
}

func (*InsertContext) IsInsertContext() {}

func NewInsertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertContext {
	var p = new(InsertContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JXBParserRULE_insert

	return p
}

func (s *InsertContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertContext) INSERT() antlr.TerminalNode {
	return s.GetToken(JXBParserINSERT, 0)
}

func (s *InsertContext) NL() antlr.TerminalNode {
	return s.GetToken(JXBParserNL, 0)
}

func (s *InsertContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(JXBParserKEYWORD, 0)
}

func (s *InsertContext) HR() antlr.TerminalNode {
	return s.GetToken(JXBParserHR, 0)
}

func (s *InsertContext) STRING() antlr.TerminalNode {
	return s.GetToken(JXBParserSTRING, 0)
}

func (s *InsertContext) ALIGNMENT() antlr.TerminalNode {
	return s.GetToken(JXBParserALIGNMENT, 0)
}

func (s *InsertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.EnterInsert(s)
	}
}

func (s *InsertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.ExitInsert(s)
	}
}

func (p *JXBParser) Insert() (localctx IInsertContext) {
	localctx = NewInsertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JXBParserRULE_insert)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(JXBParserINSERT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1568) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JXBParserALIGNMENT {
		{
			p.SetState(52)
			p.Match(JXBParserALIGNMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(55)
		p.Match(JXBParserNL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMarotoContext is an interface to support dynamic dispatch.
type IMarotoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ROW() antlr.TerminalNode
	NL() antlr.TerminalNode
	AllMarcol() []IMarcolContext
	Marcol(i int) IMarcolContext

	// IsMarotoContext differentiates from other interfaces.
	IsMarotoContext()
}

type MarotoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarotoContext() *MarotoContext {
	var p = new(MarotoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_maroto
	return p
}

func InitEmptyMarotoContext(p *MarotoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_maroto
}

func (*MarotoContext) IsMarotoContext() {}

func NewMarotoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarotoContext {
	var p = new(MarotoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JXBParserRULE_maroto

	return p
}

func (s *MarotoContext) GetParser() antlr.Parser { return s.parser }

func (s *MarotoContext) ROW() antlr.TerminalNode {
	return s.GetToken(JXBParserROW, 0)
}

func (s *MarotoContext) NL() antlr.TerminalNode {
	return s.GetToken(JXBParserNL, 0)
}

func (s *MarotoContext) AllMarcol() []IMarcolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMarcolContext); ok {
			len++
		}
	}

	tst := make([]IMarcolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMarcolContext); ok {
			tst[i] = t.(IMarcolContext)
			i++
		}
	}

	return tst
}

func (s *MarotoContext) Marcol(i int) IMarcolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarcolContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMarcolContext)
}

func (s *MarotoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarotoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarotoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.EnterMaroto(s)
	}
}

func (s *MarotoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.ExitMaroto(s)
	}
}

func (p *JXBParser) Maroto() (localctx IMarotoContext) {
	localctx = NewMarotoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JXBParserRULE_maroto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(JXBParserROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JXBParserCOL {
		{
			p.SetState(58)
			p.Marcol()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(JXBParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(JXBParserNL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMarcolContext is an interface to support dynamic dispatch.
type IMarcolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COL() antlr.TerminalNode
	Insert() IInsertContext

	// IsMarcolContext differentiates from other interfaces.
	IsMarcolContext()
}

type MarcolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarcolContext() *MarcolContext {
	var p = new(MarcolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_marcol
	return p
}

func InitEmptyMarcolContext(p *MarcolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JXBParserRULE_marcol
}

func (*MarcolContext) IsMarcolContext() {}

func NewMarcolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarcolContext {
	var p = new(MarcolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JXBParserRULE_marcol

	return p
}

func (s *MarcolContext) GetParser() antlr.Parser { return s.parser }

func (s *MarcolContext) COL() antlr.TerminalNode {
	return s.GetToken(JXBParserCOL, 0)
}

func (s *MarcolContext) Insert() IInsertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertContext)
}

func (s *MarcolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarcolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarcolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.EnterMarcol(s)
	}
}

func (s *MarcolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JXBListener); ok {
		listenerT.ExitMarcol(s)
	}
}

func (p *JXBParser) Marcol() (localctx IMarcolContext) {
	localctx = NewMarcolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JXBParserRULE_marcol)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(JXBParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Insert()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
