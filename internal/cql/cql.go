package cql

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	log "github.com/sirupsen/logrus"
)

func TranspileToSQL(cqlStr string) (string, error) {
	if len(cqlStr) < 1 {
		return "", nil
	}
	// Setup the input
	is := antlr.NewInputStream(cqlStr)

	parseErrors := &CqlErrorListener{}

	// Create the Lexer
	lexer := NewCqlLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(parseErrors)

	//-- create parser
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewCQL(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parseErrors)

	tree := parser.CqlFilter()
	// Finally parse the expression
	listener := NewCqlListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if parseErrors.errorCount > 0 {
		log.Debug("CQL parser error = " + parseErrors.msg)
		msg := syntaxErrorMsg(cqlStr, parseErrors.col)
		err := fmt.Errorf("CQL syntax error: %s", msg)
		return "", err
	}
	return listener.GetSQL(), nil
}

func syntaxErrorMsg(input string, col int) string {
	start := 0
	dots1 := ""
	if col > 30 {
		start = col - 20
		dots1 = "..."
	}

	end := len(input)
	dots2 := ""
	if col < end-30 {
		end = col + 20
		dots2 = "..."
	}

	//-- extract parts and insert error flag
	before := input[start : col-1]
	after := input[col:end]
	msg := "\"" + dots1 + before + " !!>> " + after + dots2 + "\""
	return msg
}

type CqlErrorListener struct {
	*antlr.DefaultErrorListener
	errorCount int
	line       int
	col        int
	msg        string
}

func (l *CqlErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//-- only report first error (but entire input is scanned by parser)
	if l.errorCount >= 1 {
		return
	}
	l.errorCount += 1
	l.line = line
	l.col = column
	l.msg = msg
}

func (l *CqlErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errorCount += 1
}
func (l *CqlErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errorCount += 1
}
func (l *CqlErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	l.errorCount += 1
}

//----------------------------------

type cqlListener struct {
	*BaseCQLListener

	// SQL strings for nodes
	result map[*antlr.BaseRuleContext]string
	// result SQL
	sql string
}

func NewCqlListener() *cqlListener {
	this := new(cqlListener)
	this.result = make(map[*antlr.BaseRuleContext]string)
	return this
}
func (l *cqlListener) GetSQL() string {
	return l.sql
}

// helper function to avoid nil pointer problems
func (l *cqlListener) sqlFrag(ctx antlr.ParserRuleContext) string {
	if ctx == nil {
		return ""
	}
	frag := l.result[ctx.GetBaseRuleContext()]
	//log.Debug("sqlFrag for " + ctx.GetText() + " -> " + frag)
	return frag
}

func (l *cqlListener) saveFrag(ctx antlr.ParserRuleContext, sql string) {
	l.result[ctx.GetBaseRuleContext()] = sql
}

// helper function to avoid nil pointer problems
func getText(ctx antlr.ParserRuleContext) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetText()
}

// helper function to avoid nil pointer problems
func getNodeText(node antlr.TerminalNode) string {
	if node == nil {
		return ""
	}
	return node.GetText()
}

/*
func (l *cqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
}
*/

func (l *cqlListener) ExitCqlFilter(ctx *CqlFilterContext) {
	l.sql = l.sqlFrag(ctx.BooleanValueExpression())
}

func (l *cqlListener) ExitBooleanValueExpression(ctx *BooleanValueExpressionContext) {
	sql := l.sqlFrag(ctx.BooleanTerm())
	if ctx.OR() != nil {
		expr := l.sqlFrag(ctx.BooleanValueExpression())
		sql = expr + " OR " + sql
	}
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitBooleanTerm(ctx *BooleanTermContext) {
	sql := l.sqlFrag(ctx.BooleanFactor())
	if ctx.AND() != nil {
		expr := l.sqlFrag(ctx.BooleanTerm())
		sql = expr + " AND " + sql
	}
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitBooleanFactor(ctx *BooleanFactorContext) {
	sql := l.sqlFrag(ctx.BooleanPrimary())
	if ctx.NOT() != nil {
		sql = " NOT " + sql
	}
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {
	var sql string
	if ctx.LEFTPAREN() == nil {
		sql = l.sqlFrag(ctx.Predicate())
	} else {
		sql = "(" + l.sqlFrag(ctx.BooleanValueExpression()) + ")"
	}
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitPredicate(ctx *PredicateContext) {
	var sql string
	if ctx.BinaryComparisonPredicate() != nil {
		sql = l.sqlFrag(ctx.BinaryComparisonPredicate())
	} else if ctx.LikePredicate() != nil {
		sql = l.sqlFrag(ctx.LikePredicate())
	} else if ctx.BetweenPredicate() != nil {
		sql = l.sqlFrag(ctx.BetweenPredicate())
	} else if ctx.IsNullPredicate() != nil {
		sql = l.sqlFrag(ctx.IsNullPredicate())
	} else if ctx.InPredicate() != nil {
		sql = l.sqlFrag(ctx.InPredicate())
	}
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
	expr1 := l.sqlFrag(ctx.ScalarExpression(0))
	expr2 := l.sqlFrag(ctx.ScalarExpression(1))
	op := getNodeText(ctx.ComparisonOperator())
	sql := expr1 + " " + op + " " + expr2
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitScalarExpression(ctx *ScalarExpressionContext) {
	var sql string
	if ctx.PropertyName() != nil {
		sql = quotedName(getText(ctx.PropertyName()))
	} else if ctx.CharacterLiteral() != nil {
		sql = quotedText(getText(ctx.CharacterLiteral()))
	} else if ctx.NumericLiteral() != nil {
		sql = getText(ctx.NumericLiteral())
	} else if ctx.BooleanLiteral() != nil {
		sql = getText(ctx.BooleanLiteral())
	}
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitLikePredicate(ctx *LikePredicateContext) {
	var sb strings.Builder
	sb.WriteString(quotedName(getText(ctx.PropertyName())))
	if ctx.NOT() != nil {
		sb.WriteString(" NOT")
	}
	op := " LIKE "
	if ctx.ILIKE() != nil {
		op = " ILIKE "
	}
	sb.WriteString(op)
	sb.WriteString(quotedText(getText(ctx.CharacterLiteral())))
	l.saveFrag(ctx, sb.String())
}

func (l *cqlListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {
	prop := quotedName(getText(ctx.PropertyName()))
	not := ""
	if ctx.NOT() != nil {
		not = " NOT"
	}
	expr1 := l.sqlFrag(ctx.ScalarExpression(0))
	expr2 := l.sqlFrag(ctx.ScalarExpression(1))
	sql := " " + prop + not + " BETWEEN " + expr1 + " AND " + expr2
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {
	prop := quotedName(getText(ctx.PropertyName()))
	not := ""
	if ctx.NOT() != nil {
		not = " NOT"
	}
	sql := " " + prop + " IS" + not + " NULL"
	l.saveFrag(ctx, sql)
}

func (l *cqlListener) ExitInPredicate(ctx *InPredicateContext) {
	var sb strings.Builder
	sb.WriteString(quotedName(getText(ctx.PropertyName())))
	if ctx.NOT() != nil {
		sb.WriteString(" NOT")
	}
	sb.WriteString(" IN (")
	inPredValueList(ctx, &sb)
	sb.WriteString(") ")
	sql := sb.String()
	l.saveFrag(ctx, sql)
}

func inPredValueList(ctx *InPredicateContext, sb *strings.Builder) {
	//-- numeric literal list?
	nums := ctx.AllNumericLiteral()
	if len(nums) > 0 {
		for i, num := range nums {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(num.GetText())
		}
		return
	}
	//-- must be string literal list
	strs := ctx.AllCharacterLiteral()
	if len(strs) > 0 {
		for i, s := range strs {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(quotedText(s.GetText()))
		}
	}
}

func toBool(boolText string) bool {
	return strings.ToLower(boolText) == "true"
}

func quotedName(name string) string {
	//-- CQL property names can be quoted
	if strings.HasPrefix(name, "\"") {
		return name
	}
	return "\"" + name + "\""
}

func quotedText(s string) string {
	//TODO: make this better (escape quotes, etc)
	return s
}
