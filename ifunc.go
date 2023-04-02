package goantlr

// 函数扩展接口
type IFunc interface {
	HelloWorld(callIndex int, funcName string, param1 string, param2 string, formula string) (string, error)
}
