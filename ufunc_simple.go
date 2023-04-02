package goantlr

import "fmt"

type UFunSimple struct {
	objName string
}

func NewUFunSimple(name string) UFunSimple {
	return UFunSimple{
		objName: name,
	}
}

func (uf UFunSimple) HelloWorld(callIndex int, funcName string, param1 string, param2 string, formula string) (
	string, error) {
	return fmt.Sprintf("funcCallIndex: %d funcName: %s, param1: %s, param2: %s, formula: %s",
		callIndex, funcName, param1, param2, formula), nil
}
