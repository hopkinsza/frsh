%{
package main

import (
	"fmt"
	//"math/big"
)

%}

%union {
	v variable
	varname string
}

%token IDENT
%token BOOL
%token INT
%token RAT
%token STRING
%token RAWSTRING
%token INVALID

// keywords
%token IF

%%

stmts:
	| stmts stmt ';'
	;

stmt:
	| simple_stmt
	//| complex_stmt
	;

simple_stmt:
	onetok
	| assignment
	;

assignment:
	IDENT '=' literal
	{
		fmt.Printf("assign %#v = %#v", $1, $3)
	}
	;

literal:
	BOOL
	| INT
	| RAT
	| STRING
	| RAWSTRING
	;

onetok: IDENT
	{
		fmt.Println("> got ident:", $1)
	}
	| BOOL
	{
		fmt.Println("> got ident:", $1)
	}
	| INT
	{
		fmt.Println("> got int:", $1)
	}
	| RAT
	{
		fmt.Println("> got rat:", $1)
	}
	| STRING
	{
		fmt.Println("> got string:", $1)
	}
	| RAWSTRING
	{
		fmt.Println("> got rawstring:", $1)
	}
	| INVALID
	{
		fmt.Println("> got invalid:", $1)
	}
	| IF
	{
		fmt.Println("> got IF")
	}
	;

%%
