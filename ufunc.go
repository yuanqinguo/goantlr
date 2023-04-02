package goantlr

import (
	"errors"
	"goantlr/parser"
)

// Parser
func (ps *Parser) VisitHelloworld(ctx *parser.HelloworldContext) interface{} {
	// TODO: check param
	return "hello world"
}

// Executor
func (exe *Executor) VisitHelloworld(ctx *parser.HelloworldContext) interface{} {
	funcName := ctx.GetStart().GetText()
	var (
		first  interface{}
		second interface{}
		index  = 0
	)

	index = exe.incrFuncCallIndex(funcName)
	first = ctx.Express(0).Accept(exe)
	if err, ok := first.(error); ok {
		exe.addFuncResult(funcName, NewResultWithError(nil, index, err))
		return err
	}
	second = ctx.Express(1).Accept(exe)
	if err, ok := second.(error); ok {
		exe.addFuncResult(funcName, NewResultWithError(nil, index, err))
		return err
	}
	if exe.uFuncHandle == nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, errors.New("uFuncHandle is nil")))
		return float64(0)
	}
	ret, err := exe.uFuncHandle.HelloWorld(index, funcName, first.(string), second.(string), ctx.GetText())
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return err
	}

	exe.addFuncResult(funcName, NewResult(ret, index))
	return ret
}
