package main

import (
	"fmt"
	"log"
	"math/big"
)

/*
 * A variable.
 */

type frshDataType int
const (
	frNil frshDataType = iota
	frBool
	frInt
	frRat
	frString
	frObject
)

type variable struct {
	constant bool
	holdsAny bool
	vtype frshDataType
	val interface{}
}

// variable inits
func (v variable) InitNil(any bool) variable {
	v.constant = false
	v.holdsAny = any
	v.vtype = frNil
	v.val = nil
	return v
}
func (v variable) InitBool(any bool) variable {
	v.constant = false
	v.holdsAny = any
	v.vtype = frBool
	v.val = false
	return v
}
func (v variable) InitInt(any bool) variable {
	v.constant = false
	v.holdsAny = any
	v.vtype = frInt
	v.val = new(big.Rat)
	return v
}
func (v variable) InitRat(any bool) variable {
	v.constant = false
	v.holdsAny = any
	v.vtype = frRat
	v.val = new(big.Rat)
	return v
}
func (v variable) InitString(any bool) variable {
	v.constant = false
	v.holdsAny = any
	v.vtype = frString
	v.val = ""
	return v
}
func (v variable) InitObject(any bool) variable {
	v.constant = false
	v.holdsAny = any
	v.vtype = frObject
	v.val = map[string]variable{}
	return v
}

func (v variable) String() string {
	return fmt.Sprintf("%v", v.val)
}

/*
 * Global variables.
 */

var globals map[string]variable = map[string]variable{}
// var scopestack = []

func printVar(ident string) {
	log.Printf("variable %v: %v\n", ident, globals[ident])
}

func testVars() {
	globals["test"] = variable{ false, true, frInt, 10 }
	printVar("test")
}
