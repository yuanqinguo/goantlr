package goantlr

import (
	"fmt"
	"goantlr/parser"
	"strings"
)

// Parser
func (ps *Parser) VisitConcatString(ctx *parser.ConcatStringContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) < 1 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, vx := range ctx.AllExpress() {
		if err := vx.Accept(ps); err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitFindString(ctx *parser.FindStringContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) < 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, vx := range ctx.AllExpress() {
		if vx.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		if err := vx.Accept(ps); err != nil {
			return err
		}
	}
	return nil
}

// Executor
func (exe *Executor) VisitConcatString(ctx *parser.ConcatStringContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	var retVals []string
	for _, vx := range ctx.AllExpress() {
		retVals = append(retVals, fmt.Sprintf("%v", vx.Accept(exe)))
	}
	retStr := strings.Join(retVals, "")
	exe.addFuncResult(funcName, NewResult(retStr, index))
	return retStr
}

func (exe *Executor) VisitFindString(ctx *parser.FindStringContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	fStr := fmt.Sprintf("%v", ctx.Express(0).Accept(exe))
	aStr := fmt.Sprintf("%v", ctx.Express(1).Accept(exe))
	var ret int64 = -1
	if len([]rune(aStr)) < len([]rune(fStr)) {
		return int64(-1)
	}
	var tSize int64
	if len(ctx.AllExpress()) > 2 {
		retSize := ctx.Express(2).Accept(exe)
		temp, err := Interface2Int64(retSize)
		if err != nil {
			exe.addFuncResult(funcName, NewResultWithError(-1, index, fmt.Errorf("%s the 3th parameter must be number", ctx.GetText())))
			return int64(-1)
		}
		if temp < 0 {
			exe.addFuncResult(funcName, NewResultWithError("", index, fmt.Errorf("%s the 3th parameter must be greater than 0", ctx.GetText())))
			return ""
		}
		tSize = temp
	}
	if int64(len([]rune(aStr))) < tSize {
		exe.addFuncResult(funcName, NewResult(-1, index))
		return int64(-1)
	}
	if tSize < 1 {
		tSize = 1
	}
	aStr = string([]rune(aStr)[tSize-1:])
	idx := strings.Index(aStr, fStr)
	if idx != -1 {
		prefix := []byte(aStr)[0:idx]
		rs := []rune(string(prefix))
		ret = int64(len(rs) + int(tSize))
	}

	exe.addFuncResult(funcName, NewResult(ret, index))
	return ret
}
