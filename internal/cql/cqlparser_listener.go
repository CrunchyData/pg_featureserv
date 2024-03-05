// Generated from CQLParser.g4 by ANTLR 4.7.

package cql // CQLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CQLParserListener is a complete listener for a parse tree produced by CQLParser.
type CQLParserListener interface {
	antlr.ParseTreeListener

	// EnterCqlFilter is called when entering the cqlFilter production.
	EnterCqlFilter(c *CqlFilterContext)

	// EnterBoolExprParen is called when entering the BoolExprParen production.
	EnterBoolExprParen(c *BoolExprParenContext)

	// EnterBoolExprAnd is called when entering the BoolExprAnd production.
	EnterBoolExprAnd(c *BoolExprAndContext)

	// EnterBoolExprNot is called when entering the BoolExprNot production.
	EnterBoolExprNot(c *BoolExprNotContext)

	// EnterBoolExprTerm is called when entering the BoolExprTerm production.
	EnterBoolExprTerm(c *BoolExprTermContext)

	// EnterBoolExprOr is called when entering the BoolExprOr production.
	EnterBoolExprOr(c *BoolExprOrContext)

	// EnterBooleanTerm is called when entering the booleanTerm production.
	EnterBooleanTerm(c *BooleanTermContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterPredicateBinaryComp is called when entering the PredicateBinaryComp production.
	EnterPredicateBinaryComp(c *PredicateBinaryCompContext)

	// EnterPredicateLike is called when entering the PredicateLike production.
	EnterPredicateLike(c *PredicateLikeContext)

	// EnterPredicateBetween is called when entering the PredicateBetween production.
	EnterPredicateBetween(c *PredicateBetweenContext)

	// EnterPredicateIn is called when entering the PredicateIn production.
	EnterPredicateIn(c *PredicateInContext)

	// EnterPredicateIsNull is called when entering the PredicateIsNull production.
	EnterPredicateIsNull(c *PredicateIsNullContext)

	// EnterBinaryComparisonPredicate is called when entering the binaryComparisonPredicate production.
	EnterBinaryComparisonPredicate(c *BinaryComparisonPredicateContext)

	// EnterIsLikePredicate is called when entering the isLikePredicate production.
	EnterIsLikePredicate(c *IsLikePredicateContext)

	// EnterIsBetweenPredicate is called when entering the isBetweenPredicate production.
	EnterIsBetweenPredicate(c *IsBetweenPredicateContext)

	// EnterIsInListPredicate is called when entering the isInListPredicate production.
	EnterIsInListPredicate(c *IsInListPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterScalarExpr is called when entering the ScalarExpr production.
	EnterScalarExpr(c *ScalarExprContext)

	// EnterScalarVal is called when entering the ScalarVal production.
	EnterScalarVal(c *ScalarValContext)

	// EnterScalarParen is called when entering the ScalarParen production.
	EnterScalarParen(c *ScalarParenContext)

	// EnterLiteralName is called when entering the LiteralName production.
	EnterLiteralName(c *LiteralNameContext)

	// EnterLiteralString is called when entering the LiteralString production.
	EnterLiteralString(c *LiteralStringContext)

	// EnterLiteralNumeric is called when entering the LiteralNumeric production.
	EnterLiteralNumeric(c *LiteralNumericContext)

	// EnterLiteralBoolean is called when entering the LiteralBoolean production.
	EnterLiteralBoolean(c *LiteralBooleanContext)

	// EnterLiteralTemporal is called when entering the LiteralTemporal production.
	EnterLiteralTemporal(c *LiteralTemporalContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterCharacterLiteral is called when entering the characterLiteral production.
	EnterCharacterLiteral(c *CharacterLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterTemporalLiteral is called when entering the temporalLiteral production.
	EnterTemporalLiteral(c *TemporalLiteralContext)

	// EnterSpatialPredicate is called when entering the spatialPredicate production.
	EnterSpatialPredicate(c *SpatialPredicateContext)

	// EnterDistancePredicate is called when entering the distancePredicate production.
	EnterDistancePredicate(c *DistancePredicateContext)

	// EnterGeomExpression is called when entering the geomExpression production.
	EnterGeomExpression(c *GeomExpressionContext)

	// EnterGeomLiteral is called when entering the geomLiteral production.
	EnterGeomLiteral(c *GeomLiteralContext)

	// EnterPoint is called when entering the point production.
	EnterPoint(c *PointContext)

	// EnterPointList is called when entering the pointList production.
	EnterPointList(c *PointListContext)

	// EnterLinestring is called when entering the linestring production.
	EnterLinestring(c *LinestringContext)

	// EnterPolygon is called when entering the polygon production.
	EnterPolygon(c *PolygonContext)

	// EnterPolygonDef is called when entering the polygonDef production.
	EnterPolygonDef(c *PolygonDefContext)

	// EnterMultiPoint is called when entering the multiPoint production.
	EnterMultiPoint(c *MultiPointContext)

	// EnterMultiLinestring is called when entering the multiLinestring production.
	EnterMultiLinestring(c *MultiLinestringContext)

	// EnterMultiPolygon is called when entering the multiPolygon production.
	EnterMultiPolygon(c *MultiPolygonContext)

	// EnterGeometryCollection is called when entering the geometryCollection production.
	EnterGeometryCollection(c *GeometryCollectionContext)

	// EnterEnvelope is called when entering the envelope production.
	EnterEnvelope(c *EnvelopeContext)

	// EnterCoordList is called when entering the coordList production.
	EnterCoordList(c *CoordListContext)

	// EnterCoordinate is called when entering the coordinate production.
	EnterCoordinate(c *CoordinateContext)

	// ExitCqlFilter is called when exiting the cqlFilter production.
	ExitCqlFilter(c *CqlFilterContext)

	// ExitBoolExprParen is called when exiting the BoolExprParen production.
	ExitBoolExprParen(c *BoolExprParenContext)

	// ExitBoolExprAnd is called when exiting the BoolExprAnd production.
	ExitBoolExprAnd(c *BoolExprAndContext)

	// ExitBoolExprNot is called when exiting the BoolExprNot production.
	ExitBoolExprNot(c *BoolExprNotContext)

	// ExitBoolExprTerm is called when exiting the BoolExprTerm production.
	ExitBoolExprTerm(c *BoolExprTermContext)

	// ExitBoolExprOr is called when exiting the BoolExprOr production.
	ExitBoolExprOr(c *BoolExprOrContext)

	// ExitBooleanTerm is called when exiting the booleanTerm production.
	ExitBooleanTerm(c *BooleanTermContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitPredicateBinaryComp is called when exiting the PredicateBinaryComp production.
	ExitPredicateBinaryComp(c *PredicateBinaryCompContext)

	// ExitPredicateLike is called when exiting the PredicateLike production.
	ExitPredicateLike(c *PredicateLikeContext)

	// ExitPredicateBetween is called when exiting the PredicateBetween production.
	ExitPredicateBetween(c *PredicateBetweenContext)

	// ExitPredicateIn is called when exiting the PredicateIn production.
	ExitPredicateIn(c *PredicateInContext)

	// ExitPredicateIsNull is called when exiting the PredicateIsNull production.
	ExitPredicateIsNull(c *PredicateIsNullContext)

	// ExitBinaryComparisonPredicate is called when exiting the binaryComparisonPredicate production.
	ExitBinaryComparisonPredicate(c *BinaryComparisonPredicateContext)

	// ExitIsLikePredicate is called when exiting the isLikePredicate production.
	ExitIsLikePredicate(c *IsLikePredicateContext)

	// ExitIsBetweenPredicate is called when exiting the isBetweenPredicate production.
	ExitIsBetweenPredicate(c *IsBetweenPredicateContext)

	// ExitIsInListPredicate is called when exiting the isInListPredicate production.
	ExitIsInListPredicate(c *IsInListPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitScalarExpr is called when exiting the ScalarExpr production.
	ExitScalarExpr(c *ScalarExprContext)

	// ExitScalarVal is called when exiting the ScalarVal production.
	ExitScalarVal(c *ScalarValContext)

	// ExitScalarParen is called when exiting the ScalarParen production.
	ExitScalarParen(c *ScalarParenContext)

	// ExitLiteralName is called when exiting the LiteralName production.
	ExitLiteralName(c *LiteralNameContext)

	// ExitLiteralString is called when exiting the LiteralString production.
	ExitLiteralString(c *LiteralStringContext)

	// ExitLiteralNumeric is called when exiting the LiteralNumeric production.
	ExitLiteralNumeric(c *LiteralNumericContext)

	// ExitLiteralBoolean is called when exiting the LiteralBoolean production.
	ExitLiteralBoolean(c *LiteralBooleanContext)

	// ExitLiteralTemporal is called when exiting the LiteralTemporal production.
	ExitLiteralTemporal(c *LiteralTemporalContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitCharacterLiteral is called when exiting the characterLiteral production.
	ExitCharacterLiteral(c *CharacterLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitTemporalLiteral is called when exiting the temporalLiteral production.
	ExitTemporalLiteral(c *TemporalLiteralContext)

	// ExitSpatialPredicate is called when exiting the spatialPredicate production.
	ExitSpatialPredicate(c *SpatialPredicateContext)

	// ExitDistancePredicate is called when exiting the distancePredicate production.
	ExitDistancePredicate(c *DistancePredicateContext)

	// ExitGeomExpression is called when exiting the geomExpression production.
	ExitGeomExpression(c *GeomExpressionContext)

	// ExitGeomLiteral is called when exiting the geomLiteral production.
	ExitGeomLiteral(c *GeomLiteralContext)

	// ExitPoint is called when exiting the point production.
	ExitPoint(c *PointContext)

	// ExitPointList is called when exiting the pointList production.
	ExitPointList(c *PointListContext)

	// ExitLinestring is called when exiting the linestring production.
	ExitLinestring(c *LinestringContext)

	// ExitPolygon is called when exiting the polygon production.
	ExitPolygon(c *PolygonContext)

	// ExitPolygonDef is called when exiting the polygonDef production.
	ExitPolygonDef(c *PolygonDefContext)

	// ExitMultiPoint is called when exiting the multiPoint production.
	ExitMultiPoint(c *MultiPointContext)

	// ExitMultiLinestring is called when exiting the multiLinestring production.
	ExitMultiLinestring(c *MultiLinestringContext)

	// ExitMultiPolygon is called when exiting the multiPolygon production.
	ExitMultiPolygon(c *MultiPolygonContext)

	// ExitGeometryCollection is called when exiting the geometryCollection production.
	ExitGeometryCollection(c *GeometryCollectionContext)

	// ExitEnvelope is called when exiting the envelope production.
	ExitEnvelope(c *EnvelopeContext)

	// ExitCoordList is called when exiting the coordList production.
	ExitCoordList(c *CoordListContext)

	// ExitCoordinate is called when exiting the coordinate production.
	ExitCoordinate(c *CoordinateContext)
}
