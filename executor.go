package goantlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"goantlr/parser"
	"strconv"
	"strings"
)

type Executor struct {
	parser.BaseexpressVisitor

	dataMap          map[string]interface{}
	funcCallIndexMap map[string]int       // 函数的执行次数索引值
	funcResultMaps   map[string][]*Result // 函数执行的结果值列表，一个函数可能在一个公式规则中被执行多次（函数嵌套）
	uFuncHandle      IFunc                // 一个接口，若需扩展函数，请实现该接口
}

func NewExecutor(uFuncHandle IFunc, dataMap map[string]interface{}) *Executor {
	return &Executor{
		uFuncHandle:      uFuncHandle,
		dataMap:          dataMap,
		funcCallIndexMap: make(map[string]int),
		funcResultMaps:   make(map[string][]*Result),
	}
}

func (exe *Executor) addFuncResult(funName string, result *Result) {
	resultList := make([]*Result, 0)
	if l, ok := exe.funcResultMaps[funName]; ok {
		resultList = l
	}
	resultList = append(resultList, result)
	exe.funcResultMaps[funName] = resultList
}
func (exe *Executor) incrFuncCallIndex(funcName string) int {
	index := 0
	if _, ok := exe.funcCallIndexMap[funcName]; ok {
		index = index + 1
		exe.funcCallIndexMap[funcName] = index
	} else {
		exe.funcCallIndexMap[funcName] = index
	}
	return index
}

func (exe *Executor) CalculateForDebug(input string) {
	val, err := exe.Calculate(input)
	if err != nil {
		fmt.Println("input:", input, "error:", err)
		return
	}
	fmt.Println("input:", input, " output:", val)
}

func (exe *Executor) Calculate(input string) (interface{}, error) {
	if len(input) == 0 {
		return input, nil
	}
	is := antlr.NewInputStream(input)
	lexer := parser.NewexpressLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	ps := parser.NewexpressParser(tokens)
	retVal := ps.Express().Accept(exe)
	if err, ok := retVal.(error); ok {
		return 0, err
	}
	return retVal, nil
}

func (exe *Executor) VisitParse(ctx *parser.ParseContext) interface{} {
	return ctx.AllExpress()
}

// (100) 正数加了括号处理
func (exe *Executor) VisitPlus(ctx *parser.PlusContext) interface{} {
	return ctx.Express().Accept(exe)
}

// -100  负数处理
func (exe *Executor) VisitMinus(ctx *parser.MinusContext) interface{} {
	val := ctx.Express().Accept(exe)
	retV, err := ReflectFloat64(val)
	if err != nil {
		return 0
	}
	return -retV
}

func (exe *Executor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		fmt.Println("VisitInteger ParseInt, text:", ctx.GetText(), " err:", err)
		return int64(0)
	}

	return val
}

func (exe *Executor) VisitFloat(ctx *parser.FloatContext) interface{} {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		fmt.Println("VisitFloat ParseFloat, text:", ctx.GetText(), " err:", err)
		return float64(0)
	}

	return val
}

func (exe *Executor) VisitText(ctx *parser.TextContext) interface{} {
	retStr := ctx.GetText()
	matchArr := codeRegex.FindAllString(retStr, -1)
	retStr = strings.ReplaceAll(retStr, `#`, "")
	if len(matchArr) > 1 {
		// 在所有的VisitXX方法中，执行Accept后，最后均将递归到此方法
		// 此方法进行去除#后，为dataMap中的key，若命中dataMap中的值，则作为结果值返回，用于运算
		if v, ok := exe.dataMap[retStr]; ok {
			return v
		}
	}

	// 若没有命中 #xxx# 的情况，表明非变量，而是输入的字符串
	retStr = strings.ReplaceAll(retStr, `"`, "")
	retStr = strings.ReplaceAll(retStr, `'`, "")
	retStr = strings.ReplaceAll(retStr, `“`, "")

	return retStr
}

func (exe *Executor) VisitIf(ctx *parser.IfContext) interface{} {
	retFlag := ctx.BoolExpress().Accept(exe)
	retv := false
	if retFlag != nil {
		retv, _ = retFlag.(bool)
	}

	if retv {
		return ctx.Express(0).Accept(exe)
	}
	return ctx.Express(1).Accept(exe)
}

func (exe *Executor) VisitComparator(ctx *parser.ComparatorContext) interface{} {
	left := ctx.Express(0).Accept(exe)
	right := ctx.Express(1).Accept(exe)
	opType := ctx.GetOp().GetText()
	switch opType {
	case ">", "＞", "<", "＜", ">=", "＞＝", "<=", "＜＝", "≥", "≤":
		if IsNumber(left) != IsNumber(right) {
			return false
		}
		if IsNumber(left) {
			vLeft, err := ReflectFloat64(left)
			if err != nil {
				return false
			}
			vRight, err1 := ReflectFloat64(right)
			if err1 != nil {
				return false
			}
			switch opType {
			case ">", "＞":
				return vLeft > vRight
			case "<", "＜":
				return vLeft < vRight
			case ">=", "≥", "＞＝":
				return vLeft >= vRight
			case "<=", "≤", "＜＝":
				return vLeft <= vRight
			default:
				return false
			}
		} else {
			vLeft := fmt.Sprintf("%v", left)
			vRight := fmt.Sprintf("%v", right)
			switch opType {
			case ">", "＞":
				return vLeft > vRight
			case "<", "＜":
				return vLeft < vRight
			case ">=", "≥", "＞＝":
				return vLeft >= vRight
			case "<=", "≤", "＜＝":
				return vLeft <= vRight
			default:
				return false
			}
		}
	case "=", "＝":
		return fmt.Sprintf("%v", left) == fmt.Sprintf("%v", right)
	case "!=", "≠":
		return fmt.Sprintf("%v", left) != fmt.Sprintf("%v", right)
	default:
		return false
	}
}

func (exe *Executor) VisitTrueOr(ctx *parser.TrueOrContext) interface{} {
	var (
		retFlag interface{}
		retv    = true
		retv1   = false
	)
	for _, ex := range ctx.AllBoolExpress() {
		retFlag = ex.Accept(exe)
		if retFlag != nil {
			retv, _ = retFlag.(bool)
		}
		retv = retv || retv1
		retv1 = retv
	}
	if retv {
		return true
	}
	return false
}

func (exe *Executor) VisitTrueAnd(ctx *parser.TrueAndContext) interface{} {
	var (
		retFlag interface{}
		retv    = true
		retv1   = true
	)
	for _, ex := range ctx.AllBoolExpress() {
		retFlag = ex.Accept(exe)
		if retFlag != nil {
			retv, _ = retFlag.(bool)
		}
		retv = retv && retv1
		retv1 = retv
	}
	if retv {
		return true
	}
	return false
}

func (exe *Executor) VisitOr(ctx *parser.OrContext) interface{} {
	var (
		retFlag interface{}
		retv    = true
		retv1   = false
	)
	for _, ex := range ctx.AllBoolExpress() {
		retFlag = ex.Accept(exe)
		if retFlag != nil {
			retv, _ = retFlag.(bool)
		}
		retv = retv || retv1
		retv1 = retv
	}

	return retv
}

func (exe *Executor) VisitAnd(ctx *parser.AndContext) interface{} {
	var (
		retFlag interface{}
		retv    = true
		retv1   = true
	)
	for _, ex := range ctx.AllBoolExpress() {
		retFlag = ex.Accept(exe)
		if retFlag != nil {
			retv, _ = retFlag.(bool)
		}
		retv = retv && retv1
		retv1 = retv
	}
	return retv
}

func (exe *Executor) VisitSubIfs(ctx *parser.SubIfsContext) interface{} {
	retV := ctx.BoolExpress().Accept(exe)
	boolRet, ok := retV.(bool)
	if ok && boolRet {
		return ctx.Express().Accept(exe)
	}
	return nil
}

func (exe *Executor) VisitIfs(ctx *parser.IfsContext) interface{} {
	for _, ex := range ctx.AllSubBoolExpress() {
		v := ex.Accept(exe)
		if v != nil {
			return v
		}
	}
	return float64(0)
}
