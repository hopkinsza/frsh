package main

import (
	"fmt"
	"log"
	//"math/big"
)

/*
 * A variable.
 */

type frshVarType int
const (
	frAny frshVarType = iota
	frBool
	frInt
	frRat
	frString
	frObject
)

type variable struct {
	vtype frshVarType
	val interface{}
}

func (v variable) String() string {
	return fmt.Sprintf("%v", v.val)
}

var globals map[string]variable = map[string]variable{}

func printVar(ident string) {
	log.Printf("variable %v: %v\n", ident, globals[ident])
}

func testVars() {
	globals["test"] = variable{ frInt, 10 }
	printVar("test")
}
