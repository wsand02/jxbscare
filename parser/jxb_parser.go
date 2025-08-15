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
		"", "'begin'", "'end'", "'insert'",
	}
	staticData.SymbolicNames = []string{
		"", "BEGIN", "END", "INSERT", "KEYWORD", "STRING", "WS", "NL",
	}
	staticData.RuleNames = []string{
		"document", "statement", "assignment", "block", "insert",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 54, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		1, 0, 5, 0, 12, 8, 0, 10, 0, 12, 0, 15, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 3, 1, 22, 8, 1, 1, 2, 1, 2, 5, 2, 26, 8, 2, 10, 2, 12, 2, 29, 9, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 37, 8, 3, 10, 3, 12, 3, 40, 9,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52,
		8, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 0, 54, 0, 13, 1, 0, 0, 0, 2, 21,
		1, 0, 0, 0, 4, 23, 1, 0, 0, 0, 6, 32, 1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10,
		12, 3, 2, 1, 0, 11, 10, 1, 0, 0, 0, 12, 15, 1, 0, 0, 0, 13, 11, 1, 0, 0,
		0, 13, 14, 1, 0, 0, 0, 14, 16, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 16, 17,
		5, 0, 0, 1, 17, 1, 1, 0, 0, 0, 18, 22, 3, 4, 2, 0, 19, 22, 3, 6, 3, 0,
		20, 22, 3, 8, 4, 0, 21, 18, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 20, 1,
		0, 0, 0, 22, 3, 1, 0, 0, 0, 23, 27, 5, 4, 0, 0, 24, 26, 5, 5, 0, 0, 25,
		24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0,
		0, 28, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 5, 7, 0, 0, 31, 5, 1,
		0, 0, 0, 32, 33, 5, 1, 0, 0, 33, 34, 5, 5, 0, 0, 34, 38, 5, 7, 0, 0, 35,
		37, 3, 2, 1, 0, 36, 35, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 0, 0,
		0, 38, 39, 1, 0, 0, 0, 39, 41, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 42,
		5, 2, 0, 0, 42, 43, 5, 5, 0, 0, 43, 44, 5, 7, 0, 0, 44, 7, 1, 0, 0, 0,
		45, 46, 5, 3, 0, 0, 46, 47, 5, 4, 0, 0, 47, 52, 5, 7, 0, 0, 48, 49, 5,
		3, 0, 0, 49, 50, 5, 5, 0, 0, 50, 52, 5, 7, 0, 0, 51, 45, 1, 0, 0, 0, 51,
		48, 1, 0, 0, 0, 52, 9, 1, 0, 0, 0, 5, 13, 21, 27, 38, 51,
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
	JXBParserEOF     = antlr.TokenEOF
	JXBParserBEGIN   = 1
	JXBParserEND     = 2
	JXBParserINSERT  = 3
	JXBParserKEYWORD = 4
	JXBParserSTRING  = 5
	JXBParserWS      = 6
	JXBParserNL      = 7
)

// JXBParser rules.
const (
	JXBParserRULE_document   = 0
	JXBParserRULE_statement  = 1
	JXBParserRULE_assignment = 2
	JXBParserRULE_block      = 3
	JXBParserRULE_insert     = 4
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26) != 0 {
		{
			p.SetState(10)
			p.Statement()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(16)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JXBParserKEYWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Assignment()
		}

	case JXBParserBEGIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Block()
		}

	case JXBParserINSERT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(20)
			p.Insert()
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
		p.SetState(23)
		p.Match(JXBParserKEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JXBParserSTRING {
		{
			p.SetState(24)
			p.Match(JXBParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
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
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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

func (s *BlockContext) AllStatement() []IStatementContext {
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

func (s *BlockContext) Statement(i int) IStatementContext {
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
		p.SetState(32)
		p.Match(JXBParserBEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(JXBParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(JXBParserNL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26) != 0 {
		{
			p.SetState(35)
			p.Statement()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(41)
		p.Match(JXBParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(JXBParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
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
	KEYWORD() antlr.TerminalNode
	NL() antlr.TerminalNode
	STRING() antlr.TerminalNode

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

func (s *InsertContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(JXBParserKEYWORD, 0)
}

func (s *InsertContext) NL() antlr.TerminalNode {
	return s.GetToken(JXBParserNL, 0)
}

func (s *InsertContext) STRING() antlr.TerminalNode {
	return s.GetToken(JXBParserSTRING, 0)
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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(JXBParserINSERT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(JXBParserKEYWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(JXBParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(JXBParserINSERT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(JXBParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(JXBParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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
