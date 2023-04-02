package goantlr

import (
	"fmt"
	"github.com/shopspring/decimal"
	"goantlr/parser"
)

func (ps *Parser) VisitMax(ctx *parser.MaxContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) < 1 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitMin(ctx *parser.MinContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) < 1 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitSum(ctx *parser.SumContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) < 1 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitAverage(ctx *parser.AverageContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) < 1 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitRound(ctx *parser.RoundContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitAdd(ctx *parser.AddContext) interface{} {
	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}

	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitSub(ctx *parser.SubContext) interface{} {
	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}

	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitMul(ctx *parser.MulContext) interface{} {
	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}

	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitDiv(ctx *parser.DivContext) interface{} {
	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}

	for i, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if i == 1 && ex.GetText() == "0" {
			return ParamErr.WithMsg(ctx.GetStart().GetText() + " The dividend cannot be 0")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitMod(ctx *parser.ModContext) interface{} {
	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}

	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (exe *Executor) VisitMax(ctx *parser.MaxContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	var maxValue decimal.Decimal
	for i, ex := range ctx.AllExpress() {
		value, err := ReflectFloat64(ex.Accept(exe))
		if err != nil {
			exe.addFuncResult(funcName, NewResultWithError(0, index, err))
			return int64(0)
		}
		dValue := decimal.NewFromFloat(value)
		if i == 0 {
			maxValue = dValue
		}
		if dValue.Cmp(maxValue) > 0 {
			maxValue = dValue
		}
	}
	retV, _ := maxValue.Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitMin(ctx *parser.MinContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	var minValue decimal.Decimal
	for i, ex := range ctx.AllExpress() {
		value, err := ReflectFloat64(ex.Accept(exe))
		if err != nil {
			exe.addFuncResult(funcName, NewResultWithError(0, index, err))
			return int64(0)
		}
		dValue := decimal.NewFromFloat(value)
		if i == 0 {
			minValue = dValue
		}
		if dValue.Cmp(minValue) < 0 {
			minValue = dValue
		}
	}
	retV, _ := minValue.Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitAdd(ctx *parser.AddContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	value1, err := ReflectFloat64(ctx.Express(0).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	value2, err := ReflectFloat64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	dValue1 := decimal.NewFromFloat(value1)
	dValue2 := decimal.NewFromFloat(value2)
	retV, _ := dValue1.Add(dValue2).Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitSub(ctx *parser.SubContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	value1, err := ReflectFloat64(ctx.Express(0).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	value2, err := ReflectFloat64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	dValue1 := decimal.NewFromFloat(value1)
	dValue2 := decimal.NewFromFloat(value2)
	retV, _ := dValue1.Sub(dValue2).Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitMul(ctx *parser.MulContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	value1, err := ReflectFloat64(ctx.Express(0).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	value2, err := ReflectFloat64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	dValue1 := decimal.NewFromFloat(value1)
	dValue2 := decimal.NewFromFloat(value2)
	retV, _ := dValue1.Mul(dValue2).Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitDiv(ctx *parser.DivContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	value1, err := ReflectFloat64(ctx.Express(0).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	value2, err := ReflectFloat64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	dValue1 := decimal.NewFromFloat(value1)
	dValue2 := decimal.NewFromFloat(value2)
	if dValue2.IsZero() {
		exe.addFuncResult(funcName, NewResultWithError(0, index, fmt.Errorf("the dividend cannot be 0")))
		return int64(0)
	}
	retV, _ := dValue1.Div(dValue2).Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitMod(ctx *parser.ModContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	value1, err := Interface2Int64(ctx.Express(0).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	value2, err := Interface2Int64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	exe.addFuncResult(funcName, NewResult(value1%value2, index))
	return value1 % value2
}

func (exe *Executor) VisitSum(ctx *parser.SumContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	var sumValue decimal.Decimal
	for _, express := range ctx.AllExpress() {
		value, err := ReflectFloat64(express.Accept(exe))
		if err != nil {
			exe.addFuncResult(funcName, NewResultWithError(0, index, err))
			return int64(0)
		}
		dValue1 := decimal.NewFromFloat(value)
		sumValue = sumValue.Add(dValue1)
	}
	retV, _ := sumValue.Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitAverage(ctx *parser.AverageContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	var sumValue decimal.Decimal
	for _, express := range ctx.AllExpress() {
		value, err := ReflectFloat64(express.Accept(exe))
		if err != nil {
			exe.addFuncResult(funcName, NewResultWithError(0, index, err))
			return int64(0)
		}
		sumValue = sumValue.Add(decimal.NewFromFloat(value))
	}
	v := decimal.NewFromFloat(float64(len(ctx.AllExpress())))
	if v.IsZero() {
		return 0
	}
	retV, _ := sumValue.Div(v).Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}

func (exe *Executor) VisitRound(ctx *parser.RoundContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	fV, err := ReflectFloat64(ctx.Express(0).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	round, err := ReflectFloat64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	retV, _ := decimal.NewFromFloat(fV).Round(int32(round)).Float64()
	exe.addFuncResult(funcName, NewResult(retV, index))
	return retV
}
