// Code generated from JXB.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JXBLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var JXBLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jxblexerLexerInit() {
	staticData := &JXBLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'begin'", "'end'", "'insert'",
	}
	staticData.SymbolicNames = []string{
		"", "BEGIN", "END", "INSERT", "KEYWORD", "STRING", "WS", "NL",
	}
	staticData.RuleNames = []string{
		"BEGIN", "END", "INSERT", "KEYWORD", "ALPHA", "DIGIT", "SYM", "STRING",
		"WS", "NL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 122, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 96, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 4, 7, 107, 8, 7, 11, 7, 12, 7, 108, 1, 8, 4, 8, 112, 8, 8, 11, 8,
		12, 8, 113, 1, 8, 1, 8, 1, 9, 4, 9, 119, 8, 9, 11, 9, 12, 9, 120, 0, 0,
		10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 0, 11, 0, 13, 0, 15, 5, 17, 6, 19, 7, 1,
		0, 5, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 5, 0, 35, 35, 43, 47, 58, 58,
		63, 64, 95, 95, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 132, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 27,
		1, 0, 0, 0, 5, 31, 1, 0, 0, 0, 7, 95, 1, 0, 0, 0, 9, 97, 1, 0, 0, 0, 11,
		99, 1, 0, 0, 0, 13, 101, 1, 0, 0, 0, 15, 106, 1, 0, 0, 0, 17, 111, 1, 0,
		0, 0, 19, 118, 1, 0, 0, 0, 21, 22, 5, 98, 0, 0, 22, 23, 5, 101, 0, 0, 23,
		24, 5, 103, 0, 0, 24, 25, 5, 105, 0, 0, 25, 26, 5, 110, 0, 0, 26, 2, 1,
		0, 0, 0, 27, 28, 5, 101, 0, 0, 28, 29, 5, 110, 0, 0, 29, 30, 5, 100, 0,
		0, 30, 4, 1, 0, 0, 0, 31, 32, 5, 105, 0, 0, 32, 33, 5, 110, 0, 0, 33, 34,
		5, 115, 0, 0, 34, 35, 5, 101, 0, 0, 35, 36, 5, 114, 0, 0, 36, 37, 5, 116,
		0, 0, 37, 6, 1, 0, 0, 0, 38, 39, 5, 99, 0, 0, 39, 96, 5, 118, 0, 0, 40,
		41, 5, 112, 0, 0, 41, 42, 5, 97, 0, 0, 42, 43, 5, 112, 0, 0, 43, 44, 5,
		101, 0, 0, 44, 45, 5, 114, 0, 0, 45, 46, 5, 115, 0, 0, 46, 47, 5, 105,
		0, 0, 47, 48, 5, 122, 0, 0, 48, 96, 5, 101, 0, 0, 49, 50, 5, 110, 0, 0,
		50, 51, 5, 97, 0, 0, 51, 52, 5, 109, 0, 0, 52, 96, 5, 101, 0, 0, 53, 54,
		5, 115, 0, 0, 54, 55, 5, 101, 0, 0, 55, 56, 5, 108, 0, 0, 56, 57, 5, 102,
		0, 0, 57, 58, 5, 105, 0, 0, 58, 96, 5, 101, 0, 0, 59, 60, 5, 97, 0, 0,
		60, 61, 5, 100, 0, 0, 61, 62, 5, 100, 0, 0, 62, 63, 5, 114, 0, 0, 63, 64,
		5, 101, 0, 0, 64, 65, 5, 115, 0, 0, 65, 96, 5, 115, 0, 0, 66, 67, 5, 112,
		0, 0, 67, 68, 5, 104, 0, 0, 68, 69, 5, 111, 0, 0, 69, 70, 5, 110, 0, 0,
		70, 96, 5, 101, 0, 0, 71, 72, 5, 101, 0, 0, 72, 73, 5, 109, 0, 0, 73, 74,
		5, 97, 0, 0, 74, 75, 5, 105, 0, 0, 75, 96, 5, 108, 0, 0, 76, 77, 5, 119,
		0, 0, 77, 78, 5, 101, 0, 0, 78, 79, 5, 98, 0, 0, 79, 80, 5, 115, 0, 0,
		80, 81, 5, 105, 0, 0, 81, 82, 5, 116, 0, 0, 82, 96, 5, 101, 0, 0, 83, 84,
		5, 116, 0, 0, 84, 85, 5, 105, 0, 0, 85, 86, 5, 116, 0, 0, 86, 87, 5, 108,
		0, 0, 87, 96, 5, 101, 0, 0, 88, 89, 5, 99, 0, 0, 89, 90, 5, 111, 0, 0,
		90, 91, 5, 110, 0, 0, 91, 92, 5, 116, 0, 0, 92, 93, 5, 97, 0, 0, 93, 94,
		5, 99, 0, 0, 94, 96, 5, 116, 0, 0, 95, 38, 1, 0, 0, 0, 95, 40, 1, 0, 0,
		0, 95, 49, 1, 0, 0, 0, 95, 53, 1, 0, 0, 0, 95, 59, 1, 0, 0, 0, 95, 66,
		1, 0, 0, 0, 95, 71, 1, 0, 0, 0, 95, 76, 1, 0, 0, 0, 95, 83, 1, 0, 0, 0,
		95, 88, 1, 0, 0, 0, 96, 8, 1, 0, 0, 0, 97, 98, 7, 0, 0, 0, 98, 10, 1, 0,
		0, 0, 99, 100, 7, 1, 0, 0, 100, 12, 1, 0, 0, 0, 101, 102, 7, 2, 0, 0, 102,
		14, 1, 0, 0, 0, 103, 107, 3, 9, 4, 0, 104, 107, 3, 11, 5, 0, 105, 107,
		3, 13, 6, 0, 106, 103, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 105, 1, 0,
		0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0,
		109, 16, 1, 0, 0, 0, 110, 112, 7, 3, 0, 0, 111, 110, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 116, 6, 8, 0, 0, 116, 18, 1, 0, 0, 0, 117, 119, 7, 4, 0, 0,
		118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 20, 1, 0, 0, 0, 6, 0, 95, 106, 108, 113, 120, 1,
		6, 0, 0,
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

// JXBLexerInit initializes any static state used to implement JXBLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJXBLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JXBLexerInit() {
	staticData := &JXBLexerLexerStaticData
	staticData.once.Do(jxblexerLexerInit)
}

// NewJXBLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJXBLexer(input antlr.CharStream) *JXBLexer {
	JXBLexerInit()
	l := new(JXBLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &JXBLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "JXB.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JXBLexer tokens.
const (
	JXBLexerBEGIN   = 1
	JXBLexerEND     = 2
	JXBLexerINSERT  = 3
	JXBLexerKEYWORD = 4
	JXBLexerSTRING  = 5
	JXBLexerWS      = 6
	JXBLexerNL      = 7
)
