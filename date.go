package goantlr

import (
	"fmt"
	"goantlr/parser"
	"math"
	"strings"
	"time"
)

// Parser
func (ps *Parser) VisitDateDiff(ctx *parser.DateDiffContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	ex := ctx.GetOp()
	if ex == nil {
		return ParamErr.WithErr(fmt.Errorf(`%s valid parameters is ["Y", "M", "D", "YM" "YD", "MD"]`, ctx.GetStart().GetText()))
	}
	op := strings.Trim(ex.GetText(), "\"")
	if op != "Y" && op != "D" && op != "M" && op != "YM" && op != "YD" && op != "MD" {
		return ParamErr.WithErr(fmt.Errorf(`%s valid parameters is ["Y", "M", "D", "YM" "YD", "MD"]`, ctx.GetStart().GetText()))
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

func (ps *Parser) VisitDateAdd(ctx *parser.DateAddContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}

	op := strings.Trim(ctx.GetOp().GetText(), "\"")
	if op != "Y" && op != "D" && op != "M" {
		return ParamErr.WithErr(fmt.Errorf(`%s valid parameters is["Y", "M", "D"]`, ctx.GetStart().GetText()))
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

func (ps *Parser) VisitToday(ctx *parser.TodayContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	return nil
}

func (ps *Parser) VisitDate(ctx *parser.DateContext) interface{} {
	funcName := ctx.GetStart().GetText()
	ps.funcMap[funcName] = true

	if len(ctx.AllExpress()) != 3 {
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

// Executor
func (exe *Executor) toStdDateTime(funcName string, val interface{}, index int64) (time.Time, error) {
	if val == "" || val == nil {
		return time.Time{}, fmt.Errorf("%s, param express error", funcName)
	}

	dateStr, ok := val.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("%s, param error", funcName)
	}
	var dstDate time.Time
	if strings.Contains(dateStr, ":") && !strings.Contains(dateStr, "-") && !strings.Contains(dateStr, "/") {
		if strings.Count(dateStr, ":") > 1 {
			dstDate1, err := time.ParseInLocation(TimeFormatLayer, dateStr, time.Local)
			if err != nil {
				return time.Time{}, fmt.Errorf("%s, param error, %s", funcName, err.Error())
			}
			dstDate = dstDate1
		} else {
			dstDate1, err := time.ParseInLocation(TimeFormatLayer1, dateStr, time.Local)
			if err != nil {
				return time.Time{}, fmt.Errorf("%s, param error, %s", funcName, err.Error())
			}
			dstDate = dstDate1
		}
	} else {
		dstDate1, err := ParseDateTime(dateStr)
		if err != nil {
			return time.Time{}, fmt.Errorf("%s, param error, %s", funcName, err.Error())
		}
		dstDate = dstDate1
	}

	return dstDate, nil
}

func (exe *Executor) VisitDateDiff(ctx *parser.DateDiffContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)

	leftDate, err := exe.toStdDateTime(funcName, ctx.Express(0).Accept(exe), 1)
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}

	rightDate, err := exe.toStdDateTime(funcName, ctx.Express(1).Accept(exe), 2)
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError(0, index, err))
		return int64(0)
	}
	var isNot bool
	if leftDate.After(rightDate) {
		isNot = true
		midDate := leftDate
		leftDate = rightDate
		rightDate = midDate
	}
	var val interface{}
	op := strings.Trim(ctx.GetOp().GetText(), "\"")
	switch op {
	case "Y":
		val = exe.dateDiffYears(leftDate, rightDate)
	case "M":
		val = exe.dateDiffMonths(leftDate, rightDate)
	case "D":
		val = exe.dateDiffDays(leftDate, rightDate)
	case "YM":
		val = exe.dateDiffYearMonths(leftDate, rightDate)
	case "MD":
		val = exe.dateDiffMonthDays(leftDate, rightDate)
	case "YD":
		val = exe.dateDiffYearDays(leftDate, rightDate)
	default:
		val = int64(0)
	}

	if isNot {
		result, _ := ReflectFloat64(val)
		result = 0 - result
		exe.addFuncResult(funcName, NewResult(result, index))
		return result
	}
	exe.addFuncResult(funcName, NewResult(val, index))
	return val
}

func (exe *Executor) dateDiffYears(leftDate time.Time, rightDate time.Time) interface{} {
	leftY, leftM, leftD := leftDate.Date()
	rightY, rightM, rightD := rightDate.Date()
	diffYear, diffMonth, diffDays := leftY-rightY, leftM-rightM, leftD-rightD
	if diffYear == 0 {
		return int64(0)
	}
	result := math.Abs(float64(diffYear))
	if diffMonth > 0 || (diffMonth == 0 && diffDays > 0) {
		return math.Abs(result - 1)
	}
	return math.Abs(float64(diffYear))
}

func (exe *Executor) dateDiffMonths(leftDate time.Time, rightDate time.Time) interface{} {
	leftY, leftM, leftD := leftDate.Date()
	rightY, rightM, rightD := rightDate.Date()
	diffYear, diffMonth, diffDay := leftY-rightY, leftM-rightM, leftD-rightD
	if diffYear == 0 && diffMonth == 0 {
		return int64(0)
	}
	result := math.Abs(float64(diffYear*MonthsPerYear + int(diffMonth)))
	if diffDay > 0 {
		return math.Abs(result - 1)
	}
	return math.Abs(result)
}

func (exe *Executor) dateDiffDays(leftDate time.Time, rightDate time.Time) interface{} {
	return math.Abs(leftDate.Sub(rightDate).Seconds() / SecondsPerDay)
}

func (exe *Executor) dateDiffYearMonths(leftDate time.Time, rightDate time.Time) interface{} {
	_, leftM, leftD := leftDate.Date()
	_, rightM, rightD := rightDate.Date()
	diffMonth, diffDays := leftM-rightM, leftD-rightD
	if diffMonth == 0 {
		if diffDays <= 0 {
			return int64(0)
		} else {
			return int64(11)
		}
	}
	if diffMonth < 0 {
		if diffDays <= 0 {
			return math.Abs(float64(diffMonth))
		} else {
			return math.Abs(float64(diffMonth) + 1)
		}
	}
	if diffMonth > 0 {
		if diffDays <= 0 {
			return math.Abs(12 - float64(diffMonth))
		} else {
			return math.Abs(12 - float64(diffMonth) - 1)
		}
	}
	return math.Abs(float64(0))
}

func (exe *Executor) dateDiffMonthDays(leftDate time.Time, rightDate time.Time) interface{} {
	_, _, leftD := leftDate.Date()
	_, _, rightD := rightDate.Date()
	diffDays := leftD - rightD
	if diffDays <= 0 {
		return math.Abs(float64(diffDays))
	}
	leftDate = rightDate.AddDate(0, -1, 0)
	result, _ := Interface2Int64(exe.dateDiffDays(leftDate, rightDate))
	return math.Abs(float64(result) - float64(diffDays))
}

func (exe *Executor) dateDiffYearDays(leftDate time.Time, rightDate time.Time) interface{} {
	leftY, leftM, leftD := leftDate.Date()
	rightY, rightM, rightD := rightDate.Date()
	diffDay := leftD - rightD
	if diffDay <= 0 {
		rightDate = time.Date(leftY, rightM, rightD, 0, 0, 0, 0, time.Local)
	} else {
		leftDate = time.Date(rightY, leftM, leftD, 0, 0, 0, 0, time.Local)
	}
	if leftDate.After(rightDate) {
		leftDate = time.Date(rightY-1, leftM, leftD, 0, 0, 0, 0, time.Local)
		rightDate = time.Date(rightY, rightM, rightD, 0, 0, 0, 0, time.Local)
	}
	return exe.dateDiffDays(leftDate, rightDate)
}

func (exe *Executor) VisitDateAdd(ctx *parser.DateAddContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)
	date := ctx.Express(0).Accept(exe)
	srcDate, err := exe.toStdDateTime(funcName, date, 1)
	if err != nil {
		return ""
	}
	dret, err := Interface2Int64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError("", index, fmt.Errorf("%s the 2th parameter must be number", ctx.GetText())))
		return ""
	}

	var val interface{}
	op := strings.Trim(ctx.GetOp().GetText(), "\"")
	switch op {
	case "D":
		val = exe.dateAddDays(srcDate, int(dret))
	case "M":
		val = exe.dateAddMonths(srcDate, int(dret))
	case "Y":
		val = exe.dateAddYears(srcDate, int(dret))
	default:
		val = srcDate
	}
	exe.addFuncResult(funcName, NewResult(val, index))
	return val
}

func (exe *Executor) dateAddDays(srcDate time.Time, val int) interface{} {
	return srcDate.AddDate(0, 0, val).Format(DateFormatLayer)
}

func (exe *Executor) dateAddMonths(srcDate time.Time, val int) interface{} {
	return srcDate.AddDate(0, val, 0).Format(DateFormatLayer)
}

func (exe *Executor) dateAddYears(srcDate time.Time, val int) interface{} {
	return srcDate.AddDate(val, 0, 0).Format(DateFormatLayer)
}

func (exe *Executor) VisitToday(ctx *parser.TodayContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)
	val := time.Now().Format(DateFormatLayer)
	exe.addFuncResult(funcName, NewResult(val, index))
	return val
}

func (exe *Executor) VisitDate(ctx *parser.DateContext) interface{} {
	funcName := ctx.GetStart().GetText()
	index := exe.incrFuncCallIndex(funcName)
	yearNum, err := Interface2Int64(ctx.Express(0).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError("", index, fmt.Errorf("%s the 1th parameter must be number", funcName)))
		return ""
	}
	monthNum, err := Interface2Int64(ctx.Express(1).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError("", index, fmt.Errorf("%s the 2th parameter must be number", funcName)))
		return ""
	}
	dayNum, err := Interface2Int64(ctx.Express(2).Accept(exe))
	if err != nil {
		exe.addFuncResult(funcName, NewResultWithError("", index, fmt.Errorf("%s the 3th parameter must be number", funcName)))
		return ""
	}

	val := time.Date(int(yearNum), time.Month(monthNum), int(dayNum), 0, 0, 0, 0, time.Local).Format(DateFormatLayer)
	exe.addFuncResult(funcName, NewResult(val, index))
	return val
}
