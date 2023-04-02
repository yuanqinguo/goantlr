package goantlr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatorVisitor_Calculate(t *testing.T) {
	uFunHandle := NewUFunSimple("name-hello")
	calcVisitor := NewExecutor(uFunHandle, nil)

	// 扩展函数示例
	calcVisitor.CalculateForDebug("HELLO_WORLD('参数1', '参数2')")

	calcVisitor.CalculateForDebug(`DATEDIF("2017-12-3","2015-10-7","MD")`)
	calcVisitor.CalculateForDebug(`DATEDIF("2015-10-7","2017-12-9","MD")`)
	calcVisitor.CalculateForDebug(`DATEDIF("2015-10-7","2017-7-3","MD")`)
	calcVisitor.CalculateForDebug(`DATEDIF("2015-10-7","2017-7-12","MD")`)
	calcVisitor.CalculateForDebug(`OR(31>3.1)`)
	calcVisitor.CalculateForDebug("1+(-2)")
	calcVisitor.CalculateForDebug("1-(-2)")
	calcVisitor.CalculateForDebug("MAX(10, 99, 100, 102)")
	calcVisitor.CalculateForDebug("MIN(10, 99, 100, 103)")
	calcVisitor.CalculateForDebug("MAX(88, MAX(10, 99), 100)")
	calcVisitor.CalculateForDebug("SUM(1 * 10, 1 + 2 - 1, 1)")
	calcVisitor.CalculateForDebug("AVERAGE(10, 99, 100, 103)")
	calcVisitor.CalculateForDebug("ROUND(10.1544, 1)")
	calcVisitor.CalculateForDebug("DATEDIF(\"1993-04-01\",\"2022-09-15\",\"YD\")")
	calcVisitor.CalculateForDebug(`FINDSTR("上海","深圳人民的夜晚，上海人民的上海滩")`)

	//
	// calcVisitor.CalculateForDebug(`DATEADD("2019-02-05", -1, D)`)
	//calcVisitor.CalculateForDebug(`DATEADD("2019-02-05", (-1), "Y")`)
	//calcVisitor.CalculateForDebug(`DATEADD("2019-02-05", -1, Y)`)
	//calcVisitor.CalculateForDebug(`DATEADD("2019-02-05", 1, D)`)
	//calcVisitor.CalculateForDebug(`DATEADD("2019-02-05", 1, M)`)
	//calcVisitor.CalculateForDebug(`DATEADD("2019-02-05", 1, Y)`)

	//calcVisitor.CalculateForDebug(`DATEDIF(TODAY(), "1998-07-25", D)`)
	//calcVisitor.CalculateForDebug(`DATEDIF(TODAY(), "1998-07-25", M)`)
	//calcVisitor.CalculateForDebug(`DATEDIF(TODAY(), "1998-07-25", Y)`)
	//calcVisitor.CalculateForDebug(`EOMONTH(TODAY(), 1)`)
	//calcVisitor.CalculateForDebug(`DATE(2022+0, 8-10, 1-20)`)
	//calcVisitor.CalculateForDebug(`OR(40001>0, (-1)>0)`)

}

func TestPlusMinus(t *testing.T) {
	// 负数需要使用括号包裹, 负号与减号token字符冲突
	var tests = []struct {
		Input string
		Want  int
	}{
		{
			Input: "-2",
			Want:  -2,
		},
		{
			Input: "1-3",
			Want:  -2,
		},
		{
			Input: "1-(-1)",
			Want:  2,
		},
		{
			Input: "1-2*2",
			Want:  -3,
		},
		{
			Input: "1-2*3",
			Want:  -5,
		},
		{
			Input: "(-1)+3",
			Want:  2,
		},
		{
			Input: "1-(-(-3))",
			Want:  -2,
		},
		{
			Input: "(-1)-(2+3)",
			Want:  -6,
		},
	}
	calcVisitor := NewExecutor(nil, nil)
	for _, v := range tests {
		t.Run(v.Input, func(t *testing.T) {
			result, err := calcVisitor.Calculate(v.Input)
			assert.NoError(t, err)

			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", v.Want) {
				t.Errorf("want %v got %v", v.Want, result)
			}
		})
	}
}
