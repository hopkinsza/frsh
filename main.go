package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("frsh: ")

	//testVars()
	//os.Exit(0)

	in := bufio.NewReader(os.Stdin)
	//in := os.Stdin

	// using text/strings
	//in := strings.NewReader(src)

	var l = new(Lexer)
	l.Init(in)

	if len(os.Args) >= 2 && os.Args[1] == "lex" {
		testLexer(l)
	} else {
		p := yyNewParser()
		p.Parse(l)
	}
}

func testLexer(l *Lexer) {
	var yylval = new(yySymType)

	for tok := l.Lex(yylval); tok != 0; tok = l.Lex(yylval) {
		fmt.Println("got token:", tok)
		fmt.Printf("yylval: %q\n", yylval)
	}
}
