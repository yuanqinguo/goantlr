// Code generated from express.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // express

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by expressParser.
type expressVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by expressParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by expressParser#Add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by expressParser#Max.
	VisitMax(ctx *MaxContext) interface{}

	// Visit a parse tree produced by expressParser#DateAdd.
	VisitDateAdd(ctx *DateAddContext) interface{}

	// Visit a parse tree produced by expressParser#Helloworld.
	VisitHelloworld(ctx *HelloworldContext) interface{}

	// Visit a parse tree produced by expressParser#Ifs.
	VisitIfs(ctx *IfsContext) interface{}

	// Visit a parse tree produced by expressParser#Round.
	VisitRound(ctx *RoundContext) interface{}

	// Visit a parse tree produced by expressParser#FindString.
	VisitFindString(ctx *FindStringContext) interface{}

	// Visit a parse tree produced by expressParser#Plus.
	VisitPlus(ctx *PlusContext) interface{}

	// Visit a parse tree produced by expressParser#DateDiff.
	VisitDateDiff(ctx *DateDiffContext) interface{}

	// Visit a parse tree produced by expressParser#Sub.
	VisitSub(ctx *SubContext) interface{}

	// Visit a parse tree produced by expressParser#Mod.
	VisitMod(ctx *ModContext) interface{}

	// Visit a parse tree produced by expressParser#TrueOr.
	VisitTrueOr(ctx *TrueOrContext) interface{}

	// Visit a parse tree produced by expressParser#Mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by expressParser#Average.
	VisitAverage(ctx *AverageContext) interface{}

	// Visit a parse tree produced by expressParser#ConcatString.
	VisitConcatString(ctx *ConcatStringContext) interface{}

	// Visit a parse tree produced by expressParser#Text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by expressParser#Sum.
	VisitSum(ctx *SumContext) interface{}

	// Visit a parse tree produced by expressParser#Date.
	VisitDate(ctx *DateContext) interface{}

	// Visit a parse tree produced by expressParser#Integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by expressParser#Div.
	VisitDiv(ctx *DivContext) interface{}

	// Visit a parse tree produced by expressParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by expressParser#Today.
	VisitToday(ctx *TodayContext) interface{}

	// Visit a parse tree produced by expressParser#Min.
	VisitMin(ctx *MinContext) interface{}

	// Visit a parse tree produced by expressParser#TrueAnd.
	VisitTrueAnd(ctx *TrueAndContext) interface{}

	// Visit a parse tree produced by expressParser#If.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by expressParser#Minus.
	VisitMinus(ctx *MinusContext) interface{}

	// Visit a parse tree produced by expressParser#Comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by expressParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by expressParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by expressParser#SubIfs.
	VisitSubIfs(ctx *SubIfsContext) interface{}
}
