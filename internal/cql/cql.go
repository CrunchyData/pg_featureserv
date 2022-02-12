package cql

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	log "github.com/sirupsen/logrus"
)

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

func TranspileToSQL(cqlStr string) string {
	if len(cqlStr) < 1 {
		return ""
	}
	// Setup the input
	is := antlr.NewInputStream(cqlStr)

	// Create the Lexer
	lexer := NewCqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewCQL(stream)
	tree := parser.CqlFilter()
	// Finally parse the expression (by walking the tree)
	listener := NewCqlListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetSQL()
}

func (l *cqlListener) GetSQL() string {
	return l.sql
}

/*
func (l *cqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
}
*/

func (l *cqlListener) ExitCqlFilter(ctx *CqlFilterContext) {
	exctx := ctx.BooleanValueExpression()
	l.sql = l.result[exctx.GetBaseRuleContext()]
}

func (l *cqlListener) ExitBooleanValueExpression(ctx *BooleanValueExpressionContext) {
	sql := l.result[ctx.BooleanTerm().GetBaseRuleContext()]
	if ctx.OR() != nil {
		expr := l.result[ctx.BooleanValueExpression().GetBaseRuleContext()]
		sql = expr + " OR " + sql
	}
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitBooleanTerm(ctx *BooleanTermContext) {
	sql := l.result[ctx.BooleanFactor().GetBaseRuleContext()]
	if ctx.AND() != nil {
		expr := l.result[ctx.BooleanTerm().GetBaseRuleContext()]
		sql = expr + " AND " + sql
	}
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitBooleanFactor(ctx *BooleanFactorContext) {
	sql := l.result[ctx.BooleanPrimary().GetBaseRuleContext()]
	if ctx.NOT() != nil {
		sql = " NOT " + sql
	}
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {
	var sql string
	if ctx.LEFTPAREN() == nil {
		sql = l.result[ctx.Predicate().GetBaseRuleContext()]
	} else {
		sql = "(" + l.result[ctx.BooleanValueExpression().GetBaseRuleContext()] + ")"
	}
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitPredicate(ctx *PredicateContext) {
	var sql string
	if ctx.BinaryComparisonPredicate() != nil {
		sql = l.result[ctx.BinaryComparisonPredicate().GetBaseRuleContext()]
	} else if ctx.LikePredicate() != nil {
		sql = l.result[ctx.LikePredicate().GetBaseRuleContext()]
	} else if ctx.BetweenPredicate() != nil {
		sql = l.result[ctx.BetweenPredicate().GetBaseRuleContext()]
	} else if ctx.IsNullPredicate() != nil {
		sql = l.result[ctx.IsNullPredicate().GetBaseRuleContext()]
	} else if ctx.InPredicate() != nil {
		sql = l.result[ctx.InPredicate().GetBaseRuleContext()]
	}
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
	expr1 := l.result[ctx.ScalarExpression(0).GetBaseRuleContext()]
	expr2 := l.result[ctx.ScalarExpression(1).GetBaseRuleContext()]
	op := ctx.ComparisonOperator().GetText()
	sql := expr1 + " " + op + " " + expr2
	log.Debug(sql)
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitScalarExpression(ctx *ScalarExpressionContext) {
	var sql string
	if ctx.PropertyName() != nil {
		sql = quotedName(ctx.PropertyName().GetText())
	} else if ctx.CharacterLiteral() != nil {
		sql = quotedText(ctx.CharacterLiteral().GetText())
	} else if ctx.NumericLiteral() != nil {
		sql = ctx.NumericLiteral().GetText()
	} else if ctx.BooleanLiteral() != nil {
		sql = ctx.BooleanLiteral().GetText()
	}
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitLikePredicate(ctx *LikePredicateContext) {
	var sb strings.Builder
	sb.WriteString(quotedName(ctx.PropertyName().GetText()))
	if ctx.NOT() != nil {
		sb.WriteString(" NOT")
	}
	op := " LIKE "
	if ctx.ILIKE() != nil {
		op = " ILIKE "
	}
	sb.WriteString(op)
	str := ctx.CharacterLiteral()
	sb.WriteString(quotedText(str.GetText()))

	l.result[ctx.GetBaseRuleContext()] = sb.String()
}

func (l *cqlListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {
	prop := quotedName(ctx.PropertyName().GetText())
	not := ""
	if ctx.NOT() != nil {
		not = " NOT"
	}
	expr1 := l.result[ctx.ScalarExpression(0).GetBaseRuleContext()]
	expr2 := l.result[ctx.ScalarExpression(1).GetBaseRuleContext()]
	sql := " " + prop + not + " BETWEEN " + expr1 + " AND " + expr2
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {
	prop := quotedName(ctx.PropertyName().GetText())
	not := ""
	if ctx.NOT() != nil {
		not = " NOT"
	}
	sql := " " + prop + " IS" + not + " NULL"
	l.result[ctx.GetBaseRuleContext()] = sql
}

func (l *cqlListener) ExitInPredicate(ctx *InPredicateContext) {
	var sb strings.Builder
	sb.WriteString(quotedName(ctx.PropertyName().GetText()))
	if ctx.NOT() != nil {
		sb.WriteString(" NOT")
	}
	sb.WriteString(" IN (")
	inPredValueList(ctx, &sb)
	sb.WriteString(") ")
	sql := sb.String()
	l.result[ctx.GetBaseRuleContext()] = sql
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
