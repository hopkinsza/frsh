%{
package main

import (
	"fmt"
	//"math/big"
)

//Bool bool
//Rat *big.Rat
//String string

//%token <String> IDENT
//%token <Bool> BOOL
//%token <Rat> INT
//%token <Rat> RAT
//%token <String> STRING
//%token <String> RAWSTRING
//%token <String> INVALID

%}

// bool, rat, string, or object
%union {
	val interface{}
}

%token <val> IDENT
%token <val> BOOL
%token <val> INT
%token <val> RAT
%token <val> STRING
%token <val> RAWSTRING
%token <val> INVALID

// keywords
%token IF

%type <val> literal

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
