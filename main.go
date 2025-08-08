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
	_      = iota
	T_NAME = 257 + iota
	T_ID
	T_NUMBER
	T_EMAIL
	T_PHOTO
	T_PAPER_SIZE
	T_CV
	T_TEXT
	T_SCOLON
	T_EOF
	T_UNDEF
	T_ERROR
)

const NOT_FOUND = -1

var tokentab map[string]int = map[string]int{
	"id":     T_ID,
	"number": T_NUMBER,
	"text":   T_TEXT,
	"error":  T_ERROR,
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

func key2tok(lex string) rune {
	val, ok := keywordtab[lex]
	if ok {
		return rune(val)
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
	currentToken rune // lookahead
	parseOK      bool
	tokens       []rune
	indexToken   int
	lexemes      map[int]string
	indexLexeme  int
}

func newParser(path string) parserState {
	p := parserState{indexToken: 0, parseOK: true, lexemes: make(map[int]string)}
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
		parser.tokens = append(parser.tokens, c)
	}
	parser.tokens = append(parser.tokens, T_EOF)
	parser.indexToken = 0
}

func (parser *parserState) getChar() {
	parser.lexemes[parser.indexLexeme] += string(parser.tokens[parser.indexToken])
	parser.indexToken++
}

func (parser *parserState) getToken() rune {
	// if parser.indexToken >= len(parser.tokens) {
	// 	return T_EOF
	// }

	for parser.tokens[parser.indexToken] != T_EOF {
		if unicode.IsSpace(parser.tokens[parser.indexToken]) {
			parser.indexToken++
		} else if unicode.IsLetter(parser.tokens[parser.indexToken]) {
			for unicode.IsLetter(parser.tokens[parser.indexToken]) || unicode.IsDigit(parser.tokens[parser.indexToken]) {
				parser.getChar()
			}
			lex := parser.lexemes[parser.indexLexeme]
			fmt.Printf("IsLetter: %s\n", lex)
			parser.indexLexeme++
			return key2tok(lex)
		} else if unicode.IsDigit(parser.tokens[parser.indexToken]) {
			for unicode.IsDigit(parser.tokens[parser.indexToken]) {
				parser.getChar()
			}
			fmt.Printf("IsDigit: %s\n", parser.lexemes[parser.indexLexeme])
			parser.indexLexeme++
			return T_NUMBER
		} else {
			parser.getChar()
			lex := parser.lexemes[parser.indexLexeme]
			fmt.Printf("Misc: %s\n", lex)
			parser.indexLexeme++
			return rune(lex2tok(lex))
		}
	}
	return T_ERROR
}

func (parser *parserState) match(t int) {
	fmt.Printf("%s\n", tok2lex(t))
	if parser.currentToken == rune(t) {
		fmt.Printf("Expected token, expected: %s found: %s\n", tok2lex(t), tok2lex(int(parser.currentToken)))
		parser.currentToken = parser.getToken()
	} else {
		parser.parseOK = false
		fmt.Printf("Unexpected token, expected: %s, found: %s\n", tok2lex(t), tok2lex(int(parser.currentToken)))
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

func (parser *parserState) cvEmail() {
	fmt.Println("in cvemail")
	parser.match(T_EMAIL)
	parser.match(T_ID)
	parser.match(T_SCOLON)
}

func (parser *parserState) cvPhoto() {
	fmt.Println("in cvemail")
	parser.match(T_PHOTO)
	parser.match(T_ID)
	parser.match(T_SCOLON)
}

func (parser *parserState) cvPaperSize() {
	fmt.Println("in cvemail")
	parser.match(T_PAPER_SIZE)
	parser.match(T_ID)
	parser.match(T_SCOLON)
}

func (parser *parserState) parse() bool {
	parser.cvHeader()
	parser.cvName()
	// något likt var dec list kanske?
	// for parser.currentToken != T_EOF {
	// 	switch parser.currentToken {
	// 	case T_NAME:
	// 		parser.cvName()
	// 	case T_EMAIL:
	// 		parser.cvEmail()
	// 	case T_PHOTO:
	// 		parser.cvPhoto()
	// 	case T_PAPER_SIZE:
	// 		parser.cvPaperSize()
	// 	default:
	// 		continue
	// 	}
	// }
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
