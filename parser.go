package goantlr

import (
	parser "goantlr/parser"

	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"regexp"
	"strconv"
	"strings"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	errList []error
}

func NewErrListener() *ErrorListener {
	return new(ErrorListener)
}

func (c *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if c.errList == nil {
		c.errList = make([]error, 0)
	}
	c.errList = append(c.errList, fmt.Errorf("line "+strconv.Itoa(line)+":"+strconv.Itoa(column+1)+" "+msg))
}

func (c *ErrorListener) Err() Errno {
	if c.errList == nil || len(c.errList) == 0 {
		return Pass
	}
	errStr := "error:"
	for _, e := range c.errList {
		errStr = errStr + e.Error() + " "
	}
	regex := "[0-9]+|'(.*?)' "
	compileRegex := regexp.MustCompile(regex)
	matchArr := compileRegex.FindAllString(errStr, -1)
	if len(matchArr) < 3 {
		return SyntaxErr.WithMsg(errStr)
	}
	return SyntaxErr.WithMsg(fmt.Sprintf("line:%s,pos:%s,err:%s", matchArr[0], matchArr[1], matchArr[2]))
}

type Parser struct {
	parser.BaseexpressVisitor

	varsMap    map[string]bool
	funcMap    map[string]bool
	bizFuncMap map[string]bool
}

func NewParser() *Parser {
	return &Parser{
		BaseexpressVisitor: parser.BaseexpressVisitor{},
		varsMap:            make(map[string]bool),
		funcMap:            make(map[string]bool),
		bizFuncMap:         make(map[string]bool),
	}
}

func (ps *Parser) Check(formula string) ([]string, Errno) {
	errListener := NewErrListener()
	iss := antlr.NewInputStream(formula)
	checkLexer := parser.NewexpressLexer(iss)
	checkLexer.AddErrorListener(errListener)
	for {
		t := checkLexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
	if err := errListener.Err(); err.NotOK() {
		return nil, err
	}
	literalNames := checkLexer.GetLiteralNames()
	literalNames = append(literalNames, []string{"*", "/", "÷", "×", ">", "＞", "<", "＜", ">=", "＞＝", "＜＝", "<=", "=", "＝", "≠", "!=", "≥", "≤", "-", "+"}...)
	return literalNames, Pass
}

func (ps *Parser) Parser(formula string) Errno {
	if len(formula) == 0 {
		return Pass
	}

	if literalNames, errNo := ps.Check(formula); errNo.NotOK() {
		return errNo
	} else {
		errNo = ps.CheckLiteralName(formula, literalNames)
		if errNo.NotOK() {
			return errNo
		}
	}

	errNo := ps.CheckLimit(formula)
	if errNo.NotOK() {
		return errNo
	}

	is := antlr.NewInputStream(formula)
	lexer := parser.NewexpressLexer(is)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewexpressParser(tokens)
	errListener := NewErrListener()
	p.AddErrorListener(errListener)

	err := p.Express().Accept(ps)
	if err == nil {
		return errListener.Err()
	}

	return err.(Errno)
}

func (ps *Parser) CheckLimit(formula string) Errno {
	var err error
	matchArr1 := codeRegex1.FindAllString(formula, -1)
	for _, i := range matchArr1 {
		formula = strings.ReplaceAll(formula, i, "")
	}
	var leftCount int
	var rightCount int
	for _, s := range []rune(formula) {
		switch string(s) {
		case ")", "）":
			rightCount++
		case "(", "（":
			leftCount++
		}
		if rightCount > leftCount {
			err = fmt.Errorf("missing left parenthesis")
		}
	}
	if leftCount > rightCount {
		err = fmt.Errorf("missing right parenthesis")
	}
	return SyntaxErr.WithErr(err)
}

func (ps *Parser) CheckLiteralName(formula string, literalNames []string) Errno {
	var err error
	formula = strings.ReplaceAll(formula, " ", "")
	formula = codeRegex2.ReplaceAllString(formula, "")
	matchArr := codeRegex.FindAllString(formula, -1)
	if len(matchArr) > 1 {
		find := strings.Contains(formula, "##")
		if find {
			err = fmt.Errorf("【%s】they need operators between them", strings.Join(matchArr, ","))
		}
	}
	return SyntaxErr.WithErr(err)
}

func (ps *Parser) GetBizFuncList() []string {
	var retList []string
	for k, _ := range ps.bizFuncMap {
		retList = append(retList, k)
	}
	return retList
}

func (ps *Parser) GetFuncList() []string {
	var retList []string
	for k, _ := range ps.funcMap {
		retList = append(retList, k)
	}
	return retList
}

func (ps *Parser) GetVarList() []string {
	var retList []string
	for k, _ := range ps.varsMap {
		retList = append(retList, k)
	}
	return retList
}

func (ps *Parser) VisitParse(ctx *parser.ParseContext) interface{} {
	return ctx.AllExpress()
}

// (100) 正数加了括号处理
func (ps *Parser) VisitPlus(ctx *parser.PlusContext) interface{} {
	if ctx.Express() == nil {
		return ParamMissErr.WithMsg(ctx.GetText())
	}
	return ctx.Express().Accept(ps)
}

// (-100)  负数处理
func (ps *Parser) VisitMinus(ctx *parser.MinusContext) interface{} {
	if ctx.Express() == nil {
		return ParamMissErr.WithMsg(ctx.GetText())
	}
	return ctx.Express().Accept(ps)
}

func (ps *Parser) VisitInteger(ctx *parser.IntegerContext) interface{} {
	return Pass
}

func (ps *Parser) VisitFloat(ctx *parser.FloatContext) interface{} {
	return Pass
}

func (ps *Parser) VisitText(ctx *parser.TextContext) interface{} {
	ps.varsMap[ctx.GetText()] = true
	return Pass
}

func (ps *Parser) VisitIf(ctx *parser.IfContext) interface{} {
	if ctx.BoolExpress() == nil {
		return CondRequiredErr.WithMsg(ctx.GetStart().GetText())
	}
	if ctx.BoolExpress().GetText() == "" {
		return CondEmptyErr.WithMsg(ctx.GetStart().GetText())
	}
	err := ctx.BoolExpress().Accept(ps)
	if err != nil {
		return err
	}
	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllExpress() {
		err = ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitComparator(ctx *parser.ComparatorContext) interface{} {
	if len(ctx.AllExpress()) != 2 {
		return ParamMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllExpress() {
		if ex.GetText() == "" {
			return ParamEmptyErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitTrueOr(ctx *parser.TrueOrContext) interface{} {
	if len(ctx.AllBoolExpress()) < 1 {
		return CondMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllBoolExpress() {
		if ex.GetText() == "" {
			return CondEmptyErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitTrueAnd(ctx *parser.TrueAndContext) interface{} {
	if len(ctx.AllBoolExpress()) < 1 {
		return CondMissErr.WithMsg(ctx.GetStart().GetText())
	}

	for _, ex := range ctx.AllBoolExpress() {
		if ex.GetText() == "" {
			return CondEmptyErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitOr(ctx *parser.OrContext) interface{} {
	if len(ctx.AllBoolExpress()) < 1 {
		return CondMissErr.WithMsg(ctx.GetStart().GetText())
	}
	for _, ex := range ctx.AllBoolExpress() {
		if ex.GetText() == "" {
			return CondEmptyErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitAnd(ctx *parser.AndContext) interface{} {
	if len(ctx.AllBoolExpress()) < 1 {
		return CondMissErr.WithMsg(ctx.GetStart().GetText())
	}

	for _, ex := range ctx.AllBoolExpress() {
		if ex.GetText() == "" {
			return CondEmptyErr.WithMsg(ctx.GetStart().GetText())
		}
		err := ex.Accept(ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Parser) VisitSubIfs(ctx *parser.SubIfsContext) interface{} {
	if ctx.BoolExpress() == nil {
		return CondMissErr.WithMsg(ctx.GetStart().GetText())
	}
	if ctx.BoolExpress().GetText() == "" {
		return CondEmptyErr.WithMsg(ctx.GetStart().GetText())
	}

	if err := ctx.BoolExpress().Accept(ps); err != nil {
		return err
	}

	if ctx.Express() == nil {
		return ParamMissErr.WithMsg(ctx.GetText())
	}
	if ctx.Express().GetText() == "" {
		return ParamEmptyErr.WithMsg(ctx.GetStart().GetText())
	}
	if err := ctx.Express().Accept(ps); err != nil {
		return err
	}
	return nil
}

func (ps *Parser) VisitIfs(ctx *parser.IfsContext) interface{} {
	if len(ctx.AllSubBoolExpress()) < 1 {
		return ParamMissErr.WithMsg(ctx.GetText())
	}

	for _, ex := range ctx.AllSubBoolExpress() {
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
