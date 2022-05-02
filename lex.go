package main

import (
	"io"
	"log"
	"math/big"
	"strings"
	"text/scanner"
)

/*
 * The lexer.
 */

type Lexer struct {
	s *scanner.Scanner
	// the previous token, used for semicolon insertion
	prevtok int
}

func (l *Lexer) Init(in io.Reader) {
	l.s = new(scanner.Scanner)
	l.s.Init(in)
	l.s.Whitespace ^= 1 << '\n' // do not skip newlines
	l.s.Error = func(s *scanner.Scanner, msg string) {
		log.Println("lexing error:", msg)
	}
}

// called by parser on parse error
func (l *Lexer) Error(msg string) {
	log.Println("parsing error:", msg)
}

// called by parser to get each new token
func (l *Lexer) Lex(yylval *yySymType) int {
RELEX:
	yylval.v = new(variable).InitNil(true)
	yylval.varname = ""

	//log.Println("position:", l.s.Position)
	tok := l.s.Scan()
	if tok == scanner.EOF {
		// EOF
		return 0
	}
	txt := l.s.TokenText()

	// special case: semicolon insertion
	if tok == '\n' {
		// todo: check for the correct prevtok
		if l.prevtok != ';' {
			return l.setprev(';')
			log.Println("semicolon inserted")
		} else {
			goto RELEX
		}
	}

	switch tok {
	case scanner.Ident:
		// return IDENT or a keyword
		return l.lexIdent(yylval)
	case scanner.Int:
		yylval.v = new(variable).InitInt(true)
		 // Grab a pointer to the real big.Rat value so we can
		 // manipulate it.
		p := yylval.v.val.(*big.Rat)
		_, ok := p.SetString(txt)
		if !ok || !p.IsInt() {
			log.Printf("bad integer %q", txt)
			return l.setprev(INVALID)
		}
		return l.setprev(INT)
	case scanner.Float:
		yylval.v = new(variable).InitRat(true)
		p := yylval.v.val.(*big.Rat)
		_, ok := p.SetString(txt)
		if !ok {
			log.Printf("bad rat %q", txt)
			return l.setprev(INVALID)
		}
		return l.setprev(RAT)
	case scanner.Char:
		if len([]rune(txt)) != 3 {
			log.Printf("bad char %q", txt)
			return l.setprev(INVALID)
		}
		yylval.v = new(variable).InitInt(true)
		p := yylval.v.val.(*big.Rat)
		p.SetInt64(int64([]rune(txt)[1]))
		return l.setprev(INT)
	case scanner.String:
		yylval.v = new(variable).InitString(true)
		yylval.v.val = txt
		return l.setprev(STRING)
	case scanner.RawString:
		yylval.v = new(variable).InitString(true)
		yylval.v.val = txt
		return l.setprev(RAWSTRING)
	default:
		// should only be one rune
		if len([]rune(txt)) != 1 {
			panic("internal lexer error")
		}

		// return some single characters as their ascii value
		if strings.ContainsAny(txt, "+-*/;=") {
			return l.setprev(int(tok))
		} else {
			yylval.v = new(variable).InitString(true)
			yylval.v.val = txt
			return l.setprev(INVALID)
		}
	}
}

func (l *Lexer) lexIdent(yylval *yySymType) int {
	txt := l.s.TokenText()
	switch txt {
	case "true":
		yylval.v = new(variable).InitBool(true)
		yylval.v.val = true
		return l.setprev(BOOL)
	case "false":
		yylval.v = new(variable).InitBool(true)
		yylval.v.val = false
		return l.setprev(BOOL)
	case "if":
		return l.setprev(IF)
	default:
		yylval.varname = txt
		//yylval.v = new(variable).InitString(true)
		//yylval.v.val = txt
		return l.setprev(IDENT)
	}
}

func (l *Lexer) setprev(i int) int {
	l.prevtok = i
	return i
}
