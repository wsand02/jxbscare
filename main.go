package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
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

const NOT_FOUND = -1

var tokentab map[string]int = map[string]int{
	"id":     T_ID,
	"text":   T_TEXT,
	";":      T_SCOLON,
	"undef":  T_UNDEF,
	"TERROR": NOT_FOUND,
}

var keywordtab map[string]int = map[string]int{
	"name":      T_NAME,
	"email":     T_EMAIL,
	"photo":     T_PHOTO,
	"papersize": T_PAPER_SIZE,
	"cv":        T_CV,
	"KERROR":    NOT_FOUND,
}

// type tab struct {
// 	text  string
// 	token int
// }

func lex2tok(lex string) int {
	val, ok := keywordtab[lex]
	if ok {
		return val
	}
	val, ok = tokentab[lex]
	if ok {
		return val
	}
	return T_ID // ?, jag översätter ju bara gammal c kod men still
}

func key2tok(lex string) int {
	val, ok := keywordtab[lex]
	if ok {
		return val
	}
	return T_ID
}

func tok2lex(tok int) string {
	for k, v := range keywordtab {
		if v == tok {
			return k
		}
	}
	for k, v := range tokentab {
		if v == tok {
			return k
		}
	}
	return "TERROR"
}

type parserState struct {
	currentToken int // lookahead
	parseOK      bool
	tokens       []int
	indexToken   int
	lexemes      []int
	indexLexeme  int
}

func newParser(path string) parserState {
	p := parserState{indexToken: 0, parseOK: true}
	p.loadCV(path)
	p.currentToken = p.getToken()
	return p
}

func (parser *parserState) loadCV(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open jxb")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading jxb", err)
		}
		fmt.Printf("%c\n", c)
		cint := int(c)
		parser.tokens = append(parser.tokens, cint)
	}
	parser.tokens = append(parser.tokens, 0)
	parser.indexToken = 0
}

func (parser *parserState) getToken() int {
	if parser.indexToken >= len(parser.tokens) {
		return T_EOF
	}

	for parser.tokens[parser.indexToken] != 0 {
		parser.indexLexeme = 0
		if unicode.IsSpace(rune(parser.tokens[parser.indexToken])) {
			parser.indexToken++
		} else if unicode.IsLetter(rune(parser.tokens[parser.indexToken])) {
			for unicode.IsLetter(rune(parser.tokens[parser.indexToken])) || unicode.IsDigit(rune(parser.tokens[parser.indexToken])) {
				// getChar()
			}
			parser.lexemes[parser.indexLexeme] = 0
			return key2tok(parser.lexemes)
		}
	}
	tok := parser.tokens[parser.indexToken]
	parser.indexToken++
	return tok
}

func (parser *parserState) match(t int) {
	if parser.currentToken == t {
		fmt.Printf("Expected token, expected: %s found: %s\n", tok2lex(t), tok2lex(parser.currentToken))
		parser.currentToken = parser.getToken()
	} else {
		parser.parseOK = false
		fmt.Printf("Unexpected token, expected: %s, found: %s\n", tok2lex(t), tok2lex(parser.currentToken))
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

// func (parser *parserState) cvEmail() {
// 	fmt.Println("in cvemail")
// 	parser.match(T_EMAIL)
// 	parser.match(T_ID)
// 	parser.match(T_SCOLON)
// }

// func (parser *parserState) cvPhoto() {
// 	fmt.Println("in cvemail")
// 	parser.match(T_PHOTO)
// 	parser.match(T_ID)
// 	parser.match(T_SCOLON)
// }

// func (parser *parserState) cvPaperSize() {
// 	fmt.Println("in cvemail")
// 	parser.match(T_PAPER_SIZE)
// 	parser.match(T_ID)
// 	parser.match(T_SCOLON)
// }

func (parser *parserState) parse() bool {
	parser.cvHeader()
	for parser.currentToken != T_EOF {
		switch parser.currentToken {
		case T_NAME:
			parser.cvName()
		case T_EMAIL:
			parser.cvName()
		case T_PHOTO:
			parser.cvName()
		case T_PAPER_SIZE:
			parser.cvName()
		}
	}
	return parser.parseOK
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
	pp := newParser("./aboow.jxb")
	if pp.parse() {
		fmt.Println("OK")
	} else {
		fmt.Println("FAILURE")
	}
}
