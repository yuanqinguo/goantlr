// Code generated from express.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // express

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type expressParser struct {
	*antlr.BaseParser
}

var expressParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expressParserInit() {
	staticData := &expressParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "", "", "", "'IF'", "'IFS'", "'ELSE'", "'AND'",
		"'OR'", "", "'CONCATSTR'", "'FINDSTR'", "'MAX'", "'MIN'", "'SUM'", "'AVERAGE'",
		"'ROUND'", "'+'", "'-'", "", "", "'%'", "'D'", "'M'", "'Y'", "'TODAY'",
		"'DATE'", "'DATEDIF'", "'DATEADD'", "'HELLO_WORLD'",
	}
	staticData.symbolicNames = []string{
		"", "WHITE_SPACE", "INTEGER", "TEXT", "FLOAT", "COMMA", "LLIMIT", "RLIMIT",
		"IF", "IFS", "ELSE", "AND", "OR", "COMPARATOR", "CONCATSTRING", "FINDSTRING",
		"MAX", "MIN", "SUM", "AVERAGE", "ROUND", "ADD", "SUB", "MUL", "DIV",
		"MOD", "DoDays", "DoMonths", "DoYears", "TODAY", "DATE", "DATEDIF",
		"DATEADD", "Helloworld",
	}
	staticData.ruleNames = []string{
		"parse", "express", "boolExpress", "subBoolExpress",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 246, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 4,
		0, 10, 8, 0, 11, 0, 12, 0, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		32, 8, 1, 10, 1, 12, 1, 35, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 56, 8, 1, 10, 1, 12, 1, 59, 9, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 5, 1, 68, 8, 1, 10, 1, 12, 1, 71, 9, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 82, 8, 1, 10, 1, 12, 1,
		85, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 94, 8, 1, 10,
		1, 12, 1, 97, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 106,
		8, 1, 10, 1, 12, 1, 109, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 118, 8, 1, 10, 1, 12, 1, 121, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 130, 8, 1, 10, 1, 12, 1, 133, 9, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 190, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 207,
		8, 1, 10, 1, 12, 1, 210, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 221, 8, 2, 10, 2, 12, 2, 224, 9, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 233, 8, 2, 10, 2, 12, 2, 236, 9, 2, 1, 2,
		1, 2, 3, 2, 240, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 0, 1, 2, 4, 0, 2,
		4, 6, 0, 0, 280, 0, 9, 1, 0, 0, 0, 2, 189, 1, 0, 0, 0, 4, 239, 1, 0, 0,
		0, 6, 241, 1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 11, 1,
		0, 0, 0, 11, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 1, 1, 0, 0, 0, 13,
		14, 6, 1, -1, 0, 14, 190, 5, 2, 0, 0, 15, 190, 5, 4, 0, 0, 16, 190, 5,
		3, 0, 0, 17, 18, 5, 8, 0, 0, 18, 19, 5, 6, 0, 0, 19, 20, 3, 4, 2, 0, 20,
		21, 5, 5, 0, 0, 21, 22, 3, 2, 1, 0, 22, 23, 5, 5, 0, 0, 23, 24, 3, 2, 1,
		0, 24, 25, 5, 7, 0, 0, 25, 190, 1, 0, 0, 0, 26, 27, 5, 9, 0, 0, 27, 28,
		5, 6, 0, 0, 28, 33, 3, 6, 3, 0, 29, 30, 5, 5, 0, 0, 30, 32, 3, 6, 3, 0,
		31, 29, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1,
		0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 37, 5, 7, 0, 0, 37,
		190, 1, 0, 0, 0, 38, 39, 5, 12, 0, 0, 39, 40, 5, 6, 0, 0, 40, 45, 3, 4,
		2, 0, 41, 42, 5, 5, 0, 0, 42, 44, 3, 4, 2, 0, 43, 41, 1, 0, 0, 0, 44, 47,
		1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 48, 49, 5, 7, 0, 0, 49, 190, 1, 0, 0, 0, 50, 51, 5,
		11, 0, 0, 51, 52, 5, 6, 0, 0, 52, 57, 3, 4, 2, 0, 53, 54, 5, 5, 0, 0, 54,
		56, 3, 4, 2, 0, 55, 53, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0,
		0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61,
		5, 7, 0, 0, 61, 190, 1, 0, 0, 0, 62, 63, 5, 14, 0, 0, 63, 64, 5, 6, 0,
		0, 64, 69, 3, 2, 1, 0, 65, 66, 5, 5, 0, 0, 66, 68, 3, 2, 1, 0, 67, 65,
		1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 72, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 7, 0, 0, 73, 190, 1,
		0, 0, 0, 74, 75, 5, 15, 0, 0, 75, 76, 5, 6, 0, 0, 76, 77, 3, 2, 1, 0, 77,
		78, 5, 5, 0, 0, 78, 83, 3, 2, 1, 0, 79, 80, 5, 5, 0, 0, 80, 82, 3, 2, 1,
		0, 81, 79, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 5, 7, 0, 0,
		87, 190, 1, 0, 0, 0, 88, 89, 5, 16, 0, 0, 89, 90, 5, 6, 0, 0, 90, 95, 3,
		2, 1, 0, 91, 92, 5, 5, 0, 0, 92, 94, 3, 2, 1, 0, 93, 91, 1, 0, 0, 0, 94,
		97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0,
		0, 97, 95, 1, 0, 0, 0, 98, 99, 5, 7, 0, 0, 99, 190, 1, 0, 0, 0, 100, 101,
		5, 17, 0, 0, 101, 102, 5, 6, 0, 0, 102, 107, 3, 2, 1, 0, 103, 104, 5, 5,
		0, 0, 104, 106, 3, 2, 1, 0, 105, 103, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0,
		107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109,
		107, 1, 0, 0, 0, 110, 111, 5, 7, 0, 0, 111, 190, 1, 0, 0, 0, 112, 113,
		5, 18, 0, 0, 113, 114, 5, 6, 0, 0, 114, 119, 3, 2, 1, 0, 115, 116, 5, 5,
		0, 0, 116, 118, 3, 2, 1, 0, 117, 115, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0,
		119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121,
		119, 1, 0, 0, 0, 122, 123, 5, 7, 0, 0, 123, 190, 1, 0, 0, 0, 124, 125,
		5, 19, 0, 0, 125, 126, 5, 6, 0, 0, 126, 131, 3, 2, 1, 0, 127, 128, 5, 5,
		0, 0, 128, 130, 3, 2, 1, 0, 129, 127, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0,
		131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0, 0, 0, 133,
		131, 1, 0, 0, 0, 134, 135, 5, 7, 0, 0, 135, 190, 1, 0, 0, 0, 136, 137,
		5, 20, 0, 0, 137, 138, 5, 6, 0, 0, 138, 139, 3, 2, 1, 0, 139, 140, 5, 5,
		0, 0, 140, 141, 3, 2, 1, 0, 141, 142, 5, 7, 0, 0, 142, 190, 1, 0, 0, 0,
		143, 144, 5, 6, 0, 0, 144, 145, 5, 22, 0, 0, 145, 146, 3, 2, 1, 0, 146,
		147, 5, 7, 0, 0, 147, 190, 1, 0, 0, 0, 148, 149, 5, 6, 0, 0, 149, 150,
		3, 2, 1, 0, 150, 151, 5, 7, 0, 0, 151, 190, 1, 0, 0, 0, 152, 153, 5, 29,
		0, 0, 153, 154, 5, 6, 0, 0, 154, 190, 5, 7, 0, 0, 155, 156, 5, 30, 0, 0,
		156, 157, 5, 6, 0, 0, 157, 158, 3, 2, 1, 0, 158, 159, 5, 5, 0, 0, 159,
		160, 3, 2, 1, 0, 160, 161, 5, 5, 0, 0, 161, 162, 3, 2, 1, 0, 162, 163,
		5, 7, 0, 0, 163, 190, 1, 0, 0, 0, 164, 165, 5, 31, 0, 0, 165, 166, 5, 6,
		0, 0, 166, 167, 3, 2, 1, 0, 167, 168, 5, 5, 0, 0, 168, 169, 3, 2, 1, 0,
		169, 170, 5, 5, 0, 0, 170, 171, 5, 3, 0, 0, 171, 172, 5, 7, 0, 0, 172,
		190, 1, 0, 0, 0, 173, 174, 5, 32, 0, 0, 174, 175, 5, 6, 0, 0, 175, 176,
		3, 2, 1, 0, 176, 177, 5, 5, 0, 0, 177, 178, 3, 2, 1, 0, 178, 179, 5, 5,
		0, 0, 179, 180, 5, 3, 0, 0, 180, 181, 5, 7, 0, 0, 181, 190, 1, 0, 0, 0,
		182, 183, 5, 33, 0, 0, 183, 184, 5, 6, 0, 0, 184, 185, 3, 2, 1, 0, 185,
		186, 5, 5, 0, 0, 186, 187, 3, 2, 1, 0, 187, 188, 5, 7, 0, 0, 188, 190,
		1, 0, 0, 0, 189, 13, 1, 0, 0, 0, 189, 15, 1, 0, 0, 0, 189, 16, 1, 0, 0,
		0, 189, 17, 1, 0, 0, 0, 189, 26, 1, 0, 0, 0, 189, 38, 1, 0, 0, 0, 189,
		50, 1, 0, 0, 0, 189, 62, 1, 0, 0, 0, 189, 74, 1, 0, 0, 0, 189, 88, 1, 0,
		0, 0, 189, 100, 1, 0, 0, 0, 189, 112, 1, 0, 0, 0, 189, 124, 1, 0, 0, 0,
		189, 136, 1, 0, 0, 0, 189, 143, 1, 0, 0, 0, 189, 148, 1, 0, 0, 0, 189,
		152, 1, 0, 0, 0, 189, 155, 1, 0, 0, 0, 189, 164, 1, 0, 0, 0, 189, 173,
		1, 0, 0, 0, 189, 182, 1, 0, 0, 0, 190, 208, 1, 0, 0, 0, 191, 192, 10, 12,
		0, 0, 192, 193, 5, 24, 0, 0, 193, 207, 3, 2, 1, 13, 194, 195, 10, 11, 0,
		0, 195, 196, 5, 23, 0, 0, 196, 207, 3, 2, 1, 12, 197, 198, 10, 10, 0, 0,
		198, 199, 5, 25, 0, 0, 199, 207, 3, 2, 1, 11, 200, 201, 10, 9, 0, 0, 201,
		202, 5, 22, 0, 0, 202, 207, 3, 2, 1, 10, 203, 204, 10, 8, 0, 0, 204, 205,
		5, 21, 0, 0, 205, 207, 3, 2, 1, 9, 206, 191, 1, 0, 0, 0, 206, 194, 1, 0,
		0, 0, 206, 197, 1, 0, 0, 0, 206, 200, 1, 0, 0, 0, 206, 203, 1, 0, 0, 0,
		207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		3, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 212, 3, 2, 1, 0, 212, 213, 5,
		13, 0, 0, 213, 214, 3, 2, 1, 0, 214, 240, 1, 0, 0, 0, 215, 216, 5, 12,
		0, 0, 216, 217, 5, 6, 0, 0, 217, 222, 3, 4, 2, 0, 218, 219, 5, 5, 0, 0,
		219, 221, 3, 4, 2, 0, 220, 218, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 222,
		1, 0, 0, 0, 225, 226, 5, 7, 0, 0, 226, 240, 1, 0, 0, 0, 227, 228, 5, 11,
		0, 0, 228, 229, 5, 6, 0, 0, 229, 234, 3, 4, 2, 0, 230, 231, 5, 5, 0, 0,
		231, 233, 3, 4, 2, 0, 232, 230, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234,
		232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 237, 1, 0, 0, 0, 236, 234,
		1, 0, 0, 0, 237, 238, 5, 7, 0, 0, 238, 240, 1, 0, 0, 0, 239, 211, 1, 0,
		0, 0, 239, 215, 1, 0, 0, 0, 239, 227, 1, 0, 0, 0, 240, 5, 1, 0, 0, 0, 241,
		242, 3, 4, 2, 0, 242, 243, 5, 5, 0, 0, 243, 244, 3, 2, 1, 0, 244, 7, 1,
		0, 0, 0, 16, 11, 33, 45, 57, 69, 83, 95, 107, 119, 131, 189, 206, 208,
		222, 234, 239,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// expressParserInit initializes any static state used to implement expressParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewexpressParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressParserInit() {
	staticData := &expressParserStaticData
	staticData.once.Do(expressParserInit)
}

// NewexpressParser produces a new parser instance for the optional input antlr.TokenStream.
func NewexpressParser(input antlr.TokenStream) *expressParser {
	ExpressParserInit()
	this := new(expressParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &expressParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "express.g4"

	return this
}

// expressParser tokens.
const (
	expressParserEOF          = antlr.TokenEOF
	expressParserWHITE_SPACE  = 1
	expressParserINTEGER      = 2
	expressParserTEXT         = 3
	expressParserFLOAT        = 4
	expressParserCOMMA        = 5
	expressParserLLIMIT       = 6
	expressParserRLIMIT       = 7
	expressParserIF           = 8
	expressParserIFS          = 9
	expressParserELSE         = 10
	expressParserAND          = 11
	expressParserOR           = 12
	expressParserCOMPARATOR   = 13
	expressParserCONCATSTRING = 14
	expressParserFINDSTRING   = 15
	expressParserMAX          = 16
	expressParserMIN          = 17
	expressParserSUM          = 18
	expressParserAVERAGE      = 19
	expressParserROUND        = 20
	expressParserADD          = 21
	expressParserSUB          = 22
	expressParserMUL          = 23
	expressParserDIV          = 24
	expressParserMOD          = 25
	expressParserDoDays       = 26
	expressParserDoMonths     = 27
	expressParserDoYears      = 28
	expressParserTODAY        = 29
	expressParserDATE         = 30
	expressParserDATEDIF      = 31
	expressParserDATEADD      = 32
	expressParserHelloworld   = 33
)

// expressParser rules.
const (
	expressParserRULE_parse          = 0
	expressParserRULE_express        = 1
	expressParserRULE_boolExpress    = 2
	expressParserRULE_subBoolExpress = 3
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = expressParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = expressParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expressParser) Parse() (localctx IParseContext) {
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, expressParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(expressParserINTEGER-2))|(1<<(expressParserTEXT-2))|(1<<(expressParserFLOAT-2))|(1<<(expressParserLLIMIT-2))|(1<<(expressParserIF-2))|(1<<(expressParserIFS-2))|(1<<(expressParserAND-2))|(1<<(expressParserOR-2))|(1<<(expressParserCONCATSTRING-2))|(1<<(expressParserFINDSTRING-2))|(1<<(expressParserMAX-2))|(1<<(expressParserMIN-2))|(1<<(expressParserSUM-2))|(1<<(expressParserAVERAGE-2))|(1<<(expressParserROUND-2))|(1<<(expressParserTODAY-2))|(1<<(expressParserDATE-2))|(1<<(expressParserDATEDIF-2))|(1<<(expressParserDATEADD-2))|(1<<(expressParserHelloworld-2)))) != 0) {
		{
			p.SetState(8)
			p.express(0)
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressContext is an interface to support dynamic dispatch.
type IExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressContext differentiates from other interfaces.
	IsExpressContext()
}

type ExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressContext() *ExpressContext {
	var p = new(ExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = expressParserRULE_express
	return p
}

func (*ExpressContext) IsExpressContext() {}

func NewExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressContext {
	var p = new(ExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = expressParserRULE_express

	return p
}

func (s *ExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressContext) CopyFrom(ctx *ExpressContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AddContext struct {
	*ExpressContext
	op antlr.Token
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *AddContext) GetOp() antlr.Token { return s.op }

func (s *AddContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *AddContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *AddContext) ADD() antlr.TerminalNode {
	return s.GetToken(expressParserADD, 0)
}

func (s *AddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type MaxContext struct {
	*ExpressContext
}

func NewMaxContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MaxContext {
	var p = new(MaxContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *MaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxContext) MAX() antlr.TerminalNode {
	return s.GetToken(expressParserMAX, 0)
}

func (s *MaxContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *MaxContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *MaxContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *MaxContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *MaxContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *MaxContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *MaxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitMax(s)

	default:
		return t.VisitChildren(s)
	}
}

type DateAddContext struct {
	*ExpressContext
	op antlr.Token
}

func NewDateAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateAddContext {
	var p = new(DateAddContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *DateAddContext) GetOp() antlr.Token { return s.op }

func (s *DateAddContext) SetOp(v antlr.Token) { s.op = v }

func (s *DateAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateAddContext) DATEADD() antlr.TerminalNode {
	return s.GetToken(expressParserDATEADD, 0)
}

func (s *DateAddContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *DateAddContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *DateAddContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *DateAddContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *DateAddContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *DateAddContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *DateAddContext) TEXT() antlr.TerminalNode {
	return s.GetToken(expressParserTEXT, 0)
}

func (s *DateAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitDateAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type HelloworldContext struct {
	*ExpressContext
}

func NewHelloworldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HelloworldContext {
	var p = new(HelloworldContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *HelloworldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelloworldContext) Helloworld() antlr.TerminalNode {
	return s.GetToken(expressParserHelloworld, 0)
}

func (s *HelloworldContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *HelloworldContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *HelloworldContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *HelloworldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, 0)
}

func (s *HelloworldContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *HelloworldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitHelloworld(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfsContext struct {
	*ExpressContext
}

func NewIfsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfsContext {
	var p = new(IfsContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *IfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfsContext) IFS() antlr.TerminalNode {
	return s.GetToken(expressParserIFS, 0)
}

func (s *IfsContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *IfsContext) AllSubBoolExpress() []ISubBoolExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubBoolExpressContext); ok {
			len++
		}
	}

	tst := make([]ISubBoolExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubBoolExpressContext); ok {
			tst[i] = t.(ISubBoolExpressContext)
			i++
		}
	}

	return tst
}

func (s *IfsContext) SubBoolExpress(i int) ISubBoolExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubBoolExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubBoolExpressContext)
}

func (s *IfsContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *IfsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *IfsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *IfsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitIfs(s)

	default:
		return t.VisitChildren(s)
	}
}

type RoundContext struct {
	*ExpressContext
}

func NewRoundContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundContext {
	var p = new(RoundContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *RoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundContext) ROUND() antlr.TerminalNode {
	return s.GetToken(expressParserROUND, 0)
}

func (s *RoundContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *RoundContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *RoundContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *RoundContext) COMMA() antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, 0)
}

func (s *RoundContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *RoundContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitRound(s)

	default:
		return t.VisitChildren(s)
	}
}

type FindStringContext struct {
	*ExpressContext
}

func NewFindStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FindStringContext {
	var p = new(FindStringContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *FindStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindStringContext) FINDSTRING() antlr.TerminalNode {
	return s.GetToken(expressParserFINDSTRING, 0)
}

func (s *FindStringContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *FindStringContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *FindStringContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *FindStringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *FindStringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *FindStringContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *FindStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitFindString(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusContext struct {
	*ExpressContext
}

func NewPlusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusContext {
	var p = new(PlusContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *PlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *PlusContext) Express() IExpressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *PlusContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *PlusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitPlus(s)

	default:
		return t.VisitChildren(s)
	}
}

type DateDiffContext struct {
	*ExpressContext
	op antlr.Token
}

func NewDateDiffContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateDiffContext {
	var p = new(DateDiffContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *DateDiffContext) GetOp() antlr.Token { return s.op }

func (s *DateDiffContext) SetOp(v antlr.Token) { s.op = v }

func (s *DateDiffContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateDiffContext) DATEDIF() antlr.TerminalNode {
	return s.GetToken(expressParserDATEDIF, 0)
}

func (s *DateDiffContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *DateDiffContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *DateDiffContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *DateDiffContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *DateDiffContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *DateDiffContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *DateDiffContext) TEXT() antlr.TerminalNode {
	return s.GetToken(expressParserTEXT, 0)
}

func (s *DateDiffContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitDateDiff(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubContext struct {
	*ExpressContext
	op antlr.Token
}

func NewSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubContext {
	var p = new(SubContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *SubContext) GetOp() antlr.Token { return s.op }

func (s *SubContext) SetOp(v antlr.Token) { s.op = v }

func (s *SubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *SubContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *SubContext) SUB() antlr.TerminalNode {
	return s.GetToken(expressParserSUB, 0)
}

func (s *SubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModContext struct {
	*ExpressContext
	op antlr.Token
}

func NewModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModContext {
	var p = new(ModContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *ModContext) GetOp() antlr.Token { return s.op }

func (s *ModContext) SetOp(v antlr.Token) { s.op = v }

func (s *ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *ModContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *ModContext) MOD() antlr.TerminalNode {
	return s.GetToken(expressParserMOD, 0)
}

func (s *ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueOrContext struct {
	*ExpressContext
}

func NewTrueOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueOrContext {
	var p = new(TrueOrContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *TrueOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueOrContext) OR() antlr.TerminalNode {
	return s.GetToken(expressParserOR, 0)
}

func (s *TrueOrContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *TrueOrContext) AllBoolExpress() []IBoolExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolExpressContext); ok {
			len++
		}
	}

	tst := make([]IBoolExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolExpressContext); ok {
			tst[i] = t.(IBoolExpressContext)
			i++
		}
	}

	return tst
}

func (s *TrueOrContext) BoolExpress(i int) IBoolExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressContext)
}

func (s *TrueOrContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *TrueOrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *TrueOrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *TrueOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitTrueOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulContext struct {
	*ExpressContext
	op antlr.Token
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *MulContext) GetOp() antlr.Token { return s.op }

func (s *MulContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *MulContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *MulContext) MUL() antlr.TerminalNode {
	return s.GetToken(expressParserMUL, 0)
}

func (s *MulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type AverageContext struct {
	*ExpressContext
}

func NewAverageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AverageContext {
	var p = new(AverageContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *AverageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AverageContext) AVERAGE() antlr.TerminalNode {
	return s.GetToken(expressParserAVERAGE, 0)
}

func (s *AverageContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *AverageContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *AverageContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *AverageContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *AverageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *AverageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *AverageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitAverage(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConcatStringContext struct {
	*ExpressContext
}

func NewConcatStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatStringContext {
	var p = new(ConcatStringContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *ConcatStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatStringContext) CONCATSTRING() antlr.TerminalNode {
	return s.GetToken(expressParserCONCATSTRING, 0)
}

func (s *ConcatStringContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *ConcatStringContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *ConcatStringContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *ConcatStringContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *ConcatStringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *ConcatStringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *ConcatStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitConcatString(s)

	default:
		return t.VisitChildren(s)
	}
}

type TextContext struct {
	*ExpressContext
}

func NewTextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextContext {
	var p = new(TextContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(expressParserTEXT, 0)
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumContext struct {
	*ExpressContext
}

func NewSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumContext {
	var p = new(SumContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) SUM() antlr.TerminalNode {
	return s.GetToken(expressParserSUM, 0)
}

func (s *SumContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *SumContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *SumContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *SumContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *SumContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *SumContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *SumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitSum(s)

	default:
		return t.VisitChildren(s)
	}
}

type DateContext struct {
	*ExpressContext
}

func NewDateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateContext {
	var p = new(DateContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) DATE() antlr.TerminalNode {
	return s.GetToken(expressParserDATE, 0)
}

func (s *DateContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *DateContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *DateContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *DateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *DateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *DateContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *DateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitDate(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerContext struct {
	*ExpressContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(expressParserINTEGER, 0)
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivContext struct {
	*ExpressContext
	op antlr.Token
}

func NewDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivContext {
	var p = new(DivContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *DivContext) GetOp() antlr.Token { return s.op }

func (s *DivContext) SetOp(v antlr.Token) { s.op = v }

func (s *DivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *DivContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *DivContext) DIV() antlr.TerminalNode {
	return s.GetToken(expressParserDIV, 0)
}

func (s *DivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	*ExpressContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(expressParserFLOAT, 0)
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type TodayContext struct {
	*ExpressContext
}

func NewTodayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TodayContext {
	var p = new(TodayContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *TodayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TodayContext) TODAY() antlr.TerminalNode {
	return s.GetToken(expressParserTODAY, 0)
}

func (s *TodayContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *TodayContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *TodayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitToday(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinContext struct {
	*ExpressContext
}

func NewMinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinContext {
	var p = new(MinContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *MinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinContext) MIN() antlr.TerminalNode {
	return s.GetToken(expressParserMIN, 0)
}

func (s *MinContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *MinContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *MinContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *MinContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *MinContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *MinContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *MinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitMin(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueAndContext struct {
	*ExpressContext
}

func NewTrueAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueAndContext {
	var p = new(TrueAndContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *TrueAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueAndContext) AND() antlr.TerminalNode {
	return s.GetToken(expressParserAND, 0)
}

func (s *TrueAndContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *TrueAndContext) AllBoolExpress() []IBoolExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolExpressContext); ok {
			len++
		}
	}

	tst := make([]IBoolExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolExpressContext); ok {
			tst[i] = t.(IBoolExpressContext)
			i++
		}
	}

	return tst
}

func (s *TrueAndContext) BoolExpress(i int) IBoolExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressContext)
}

func (s *TrueAndContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *TrueAndContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *TrueAndContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *TrueAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitTrueAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfContext struct {
	*ExpressContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(expressParserIF, 0)
}

func (s *IfContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *IfContext) BoolExpress() IBoolExpressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressContext)
}

func (s *IfContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *IfContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *IfContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *IfContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinusContext struct {
	*ExpressContext
}

func NewMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusContext {
	var p = new(MinusContext)

	p.ExpressContext = NewEmptyExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressContext))

	return p
}

func (s *MinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *MinusContext) SUB() antlr.TerminalNode {
	return s.GetToken(expressParserSUB, 0)
}

func (s *MinusContext) Express() IExpressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *MinusContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *MinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expressParser) Express() (localctx IExpressContext) {
	return p.express(0)
}

func (p *expressParser) express(_p int) (localctx IExpressContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, expressParserRULE_express, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(14)
			p.Match(expressParserINTEGER)
		}

	case 2:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Match(expressParserFLOAT)
		}

	case 3:
		localctx = NewTextContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)
			p.Match(expressParserTEXT)
		}

	case 4:
		localctx = NewIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(expressParserIF)
		}
		{
			p.SetState(18)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(19)
			p.BoolExpress()
		}
		{
			p.SetState(20)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(21)
			p.express(0)
		}
		{
			p.SetState(22)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(23)
			p.express(0)
		}
		{
			p.SetState(24)
			p.Match(expressParserRLIMIT)
		}

	case 5:
		localctx = NewIfsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Match(expressParserIFS)
		}
		{
			p.SetState(27)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(28)
			p.SubBoolExpress()
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(29)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(30)
				p.SubBoolExpress()
			}

			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(36)
			p.Match(expressParserRLIMIT)
		}

	case 6:
		localctx = NewTrueOrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(expressParserOR)
		}
		{
			p.SetState(39)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(40)
			p.BoolExpress()
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(41)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(42)
				p.BoolExpress()
			}

			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(48)
			p.Match(expressParserRLIMIT)
		}

	case 7:
		localctx = NewTrueAndContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(expressParserAND)
		}
		{
			p.SetState(51)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(52)
			p.BoolExpress()
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(53)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(54)
				p.BoolExpress()
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(expressParserRLIMIT)
		}

	case 8:
		localctx = NewConcatStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(expressParserCONCATSTRING)
		}
		{
			p.SetState(63)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(64)
			p.express(0)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(65)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(66)
				p.express(0)
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(72)
			p.Match(expressParserRLIMIT)
		}

	case 9:
		localctx = NewFindStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Match(expressParserFINDSTRING)
		}
		{
			p.SetState(75)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(76)
			p.express(0)
		}
		{
			p.SetState(77)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(78)
			p.express(0)
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(79)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(80)
				p.express(0)
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(86)
			p.Match(expressParserRLIMIT)
		}

	case 10:
		localctx = NewMaxContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(expressParserMAX)
		}
		{
			p.SetState(89)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(90)
			p.express(0)
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(91)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(92)
				p.express(0)
			}

			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(98)
			p.Match(expressParserRLIMIT)
		}

	case 11:
		localctx = NewMinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(expressParserMIN)
		}
		{
			p.SetState(101)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(102)
			p.express(0)
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(103)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(104)
				p.express(0)
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(110)
			p.Match(expressParserRLIMIT)
		}

	case 12:
		localctx = NewSumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(112)
			p.Match(expressParserSUM)
		}
		{
			p.SetState(113)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(114)
			p.express(0)
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(115)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(116)
				p.express(0)
			}

			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(122)
			p.Match(expressParserRLIMIT)
		}

	case 13:
		localctx = NewAverageContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(124)
			p.Match(expressParserAVERAGE)
		}
		{
			p.SetState(125)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(126)
			p.express(0)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(127)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(128)
				p.express(0)
			}

			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(134)
			p.Match(expressParserRLIMIT)
		}

	case 14:
		localctx = NewRoundContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.Match(expressParserROUND)
		}
		{
			p.SetState(137)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(138)
			p.express(0)
		}
		{
			p.SetState(139)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(140)
			p.express(0)
		}
		{
			p.SetState(141)
			p.Match(expressParserRLIMIT)
		}

	case 15:
		localctx = NewMinusContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(144)
			p.Match(expressParserSUB)
		}
		{
			p.SetState(145)
			p.express(0)
		}
		{
			p.SetState(146)
			p.Match(expressParserRLIMIT)
		}

	case 16:
		localctx = NewPlusContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(148)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(149)
			p.express(0)
		}
		{
			p.SetState(150)
			p.Match(expressParserRLIMIT)
		}

	case 17:
		localctx = NewTodayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(152)
			p.Match(expressParserTODAY)
		}
		{
			p.SetState(153)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(154)
			p.Match(expressParserRLIMIT)
		}

	case 18:
		localctx = NewDateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(155)
			p.Match(expressParserDATE)
		}
		{
			p.SetState(156)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(157)
			p.express(0)
		}
		{
			p.SetState(158)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(159)
			p.express(0)
		}
		{
			p.SetState(160)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(161)
			p.express(0)
		}
		{
			p.SetState(162)
			p.Match(expressParserRLIMIT)
		}

	case 19:
		localctx = NewDateDiffContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(164)
			p.Match(expressParserDATEDIF)
		}
		{
			p.SetState(165)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(166)
			p.express(0)
		}
		{
			p.SetState(167)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(168)
			p.express(0)
		}
		{
			p.SetState(169)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(170)

			var _m = p.Match(expressParserTEXT)

			localctx.(*DateDiffContext).op = _m
		}
		{
			p.SetState(171)
			p.Match(expressParserRLIMIT)
		}

	case 20:
		localctx = NewDateAddContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(173)
			p.Match(expressParserDATEADD)
		}
		{
			p.SetState(174)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(175)
			p.express(0)
		}
		{
			p.SetState(176)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(177)
			p.express(0)
		}
		{
			p.SetState(178)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(179)

			var _m = p.Match(expressParserTEXT)

			localctx.(*DateAddContext).op = _m
		}
		{
			p.SetState(180)
			p.Match(expressParserRLIMIT)
		}

	case 21:
		localctx = NewHelloworldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(182)
			p.Match(expressParserHelloworld)
		}
		{
			p.SetState(183)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(184)
			p.express(0)
		}
		{
			p.SetState(185)
			p.Match(expressParserCOMMA)
		}
		{
			p.SetState(186)
			p.express(0)
		}
		{
			p.SetState(187)
			p.Match(expressParserRLIMIT)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDivContext(p, NewExpressContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expressParserRULE_express)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(192)

					var _m = p.Match(expressParserDIV)

					localctx.(*DivContext).op = _m
				}
				{
					p.SetState(193)
					p.express(13)
				}

			case 2:
				localctx = NewMulContext(p, NewExpressContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expressParserRULE_express)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(195)

					var _m = p.Match(expressParserMUL)

					localctx.(*MulContext).op = _m
				}
				{
					p.SetState(196)
					p.express(12)
				}

			case 3:
				localctx = NewModContext(p, NewExpressContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expressParserRULE_express)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(198)

					var _m = p.Match(expressParserMOD)

					localctx.(*ModContext).op = _m
				}
				{
					p.SetState(199)
					p.express(11)
				}

			case 4:
				localctx = NewSubContext(p, NewExpressContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expressParserRULE_express)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(201)

					var _m = p.Match(expressParserSUB)

					localctx.(*SubContext).op = _m
				}
				{
					p.SetState(202)
					p.express(10)
				}

			case 5:
				localctx = NewAddContext(p, NewExpressContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expressParserRULE_express)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(204)

					var _m = p.Match(expressParserADD)

					localctx.(*AddContext).op = _m
				}
				{
					p.SetState(205)
					p.express(9)
				}

			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolExpressContext is an interface to support dynamic dispatch.
type IBoolExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExpressContext differentiates from other interfaces.
	IsBoolExpressContext()
}

type BoolExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExpressContext() *BoolExpressContext {
	var p = new(BoolExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = expressParserRULE_boolExpress
	return p
}

func (*BoolExpressContext) IsBoolExpressContext() {}

func NewBoolExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExpressContext {
	var p = new(BoolExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = expressParserRULE_boolExpress

	return p
}

func (s *BoolExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExpressContext) CopyFrom(ctx *BoolExpressContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OrContext struct {
	*BoolExpressContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.BoolExpressContext = NewEmptyBoolExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(expressParserOR, 0)
}

func (s *OrContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *OrContext) AllBoolExpress() []IBoolExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolExpressContext); ok {
			len++
		}
	}

	tst := make([]IBoolExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolExpressContext); ok {
			tst[i] = t.(IBoolExpressContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) BoolExpress(i int) IBoolExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressContext)
}

func (s *OrContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *OrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *OrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	*BoolExpressContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.BoolExpressContext = NewEmptyBoolExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(expressParserAND, 0)
}

func (s *AndContext) LLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserLLIMIT, 0)
}

func (s *AndContext) AllBoolExpress() []IBoolExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolExpressContext); ok {
			len++
		}
	}

	tst := make([]IBoolExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolExpressContext); ok {
			tst[i] = t.(IBoolExpressContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) BoolExpress(i int) IBoolExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressContext)
}

func (s *AndContext) RLIMIT() antlr.TerminalNode {
	return s.GetToken(expressParserRLIMIT, 0)
}

func (s *AndContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expressParserCOMMA)
}

func (s *AndContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, i)
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparatorContext struct {
	*BoolExpressContext
	op antlr.Token
}

func NewComparatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BoolExpressContext = NewEmptyBoolExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressContext))

	return p
}

func (s *ComparatorContext) GetOp() antlr.Token { return s.op }

func (s *ComparatorContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) AllExpress() []IExpressContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressContext); ok {
			len++
		}
	}

	tst := make([]IExpressContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressContext); ok {
			tst[i] = t.(IExpressContext)
			i++
		}
	}

	return tst
}

func (s *ComparatorContext) Express(i int) IExpressContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *ComparatorContext) COMPARATOR() antlr.TerminalNode {
	return s.GetToken(expressParserCOMPARATOR, 0)
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expressParser) BoolExpress() (localctx IBoolExpressContext) {
	this := p
	_ = this

	localctx = NewBoolExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, expressParserRULE_boolExpress)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComparatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.express(0)
		}
		{
			p.SetState(212)

			var _m = p.Match(expressParserCOMPARATOR)

			localctx.(*ComparatorContext).op = _m
		}
		{
			p.SetState(213)
			p.express(0)
		}

	case 2:
		localctx = NewOrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(expressParserOR)
		}
		{
			p.SetState(216)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(217)
			p.BoolExpress()
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(218)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(219)
				p.BoolExpress()
			}

			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(225)
			p.Match(expressParserRLIMIT)
		}

	case 3:
		localctx = NewAndContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(227)
			p.Match(expressParserAND)
		}
		{
			p.SetState(228)
			p.Match(expressParserLLIMIT)
		}
		{
			p.SetState(229)
			p.BoolExpress()
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == expressParserCOMMA {
			{
				p.SetState(230)
				p.Match(expressParserCOMMA)
			}
			{
				p.SetState(231)
				p.BoolExpress()
			}

			p.SetState(236)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(237)
			p.Match(expressParserRLIMIT)
		}

	}

	return localctx
}

// ISubBoolExpressContext is an interface to support dynamic dispatch.
type ISubBoolExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubBoolExpressContext differentiates from other interfaces.
	IsSubBoolExpressContext()
}

type SubBoolExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubBoolExpressContext() *SubBoolExpressContext {
	var p = new(SubBoolExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = expressParserRULE_subBoolExpress
	return p
}

func (*SubBoolExpressContext) IsSubBoolExpressContext() {}

func NewSubBoolExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubBoolExpressContext {
	var p = new(SubBoolExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = expressParserRULE_subBoolExpress

	return p
}

func (s *SubBoolExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *SubBoolExpressContext) CopyFrom(ctx *SubBoolExpressContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SubBoolExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubBoolExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SubIfsContext struct {
	*SubBoolExpressContext
}

func NewSubIfsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubIfsContext {
	var p = new(SubIfsContext)

	p.SubBoolExpressContext = NewEmptySubBoolExpressContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SubBoolExpressContext))

	return p
}

func (s *SubIfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubIfsContext) BoolExpress() IBoolExpressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressContext)
}

func (s *SubIfsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(expressParserCOMMA, 0)
}

func (s *SubIfsContext) Express() IExpressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *SubIfsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expressVisitor:
		return t.VisitSubIfs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expressParser) SubBoolExpress() (localctx ISubBoolExpressContext) {
	this := p
	_ = this

	localctx = NewSubBoolExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, expressParserRULE_subBoolExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewSubIfsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.BoolExpress()
	}
	{
		p.SetState(242)
		p.Match(expressParserCOMMA)
	}
	{
		p.SetState(243)
		p.express(0)
	}

	return localctx
}

func (p *expressParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressContext = nil
		if localctx != nil {
			t = localctx.(*ExpressContext)
		}
		return p.Express_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *expressParser) Express_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
