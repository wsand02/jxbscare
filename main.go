package main

import (
	"fmt"
	"log"
)

const (
	_ = iota
	T_NAME
	T_ID
	T_EMAIL
	T_PHOTO
	T_PAPER_SIZE
	T_CV
	T_TEXT
	T_SCOLON
	T_EOF
	T_UNDEF
)

func tokToString(t int) string {
	switch t {
	case 1:
		return "T_NAME"
	case 2:
		return "T_ID"
	case 3:
		return "T_EMAIL"
	case 4:
		return "T_PHOTO"
	case 5:
		return "T_PAPER_SIZE"
	case 6:
		return "T_CV"
	case 7:
		return "T_TEXT"
	case 8:
		return "T_SCOLON"
	case 9:
		return "T_EOF"
	default:
		return "T_UNDEF"
	}
}

type parserState struct {
	currentToken int
	parseOK      bool
	tokens       []int
	index        int
}

func newParser(tt []int) parserState {
	p := parserState{index: 0, parseOK: true, tokens: tt}
	p.currentToken = p.getToken()
	return p
}

func (parser *parserState) getToken() int {
	if parser.index >= len(parser.tokens) {
		return T_EOF
	}
	tok := parser.tokens[parser.index]
	parser.index++
	return tok
}

func (parser *parserState) match(t int) {
	fmt.Printf("Expected token, expected: %s found: %s\n", tokToString(t), tokToString(parser.currentToken))
	if parser.currentToken == t {
		parser.currentToken = parser.getToken()
	} else {
		parser.parseOK = false
		log.Fatalf("Unexpected token, expected: %d, found: %d", tokToString(t), tokToString(parser.currentToken))
	}
}

func (parser *parserState) cvHeader() {
	fmt.Println("in cvhead")
	parser.match(T_CV)
	parser.match(T_ID)
	parser.match(T_SCOLON)
}

func (parser *parserState) cvName() {
	fmt.Println("in cvname")
	parser.match(T_NAME)
	parser.match(T_ID)
	parser.match(T_SCOLON)
}

// Parsed data structure
type CV struct {
	Name      string
	Email     string
	Photo     string
	PaperSize string
	CVName    string
}

func main() {
	pp := newParser([]int{T_CV, T_ID, T_SCOLON, T_NAME, T_ID, T_SCOLON})
	pp.cvHeader()
	pp.cvName()
	if pp.parseOK {
		fmt.Println("OK")
	} else {
		fmt.Println("FAILURE")
	}
}
