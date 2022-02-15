// Generated from CQL.g4 by ANTLR 4.7.

package cql // CQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCQLListener is a complete listener for a parse tree produced by CQL.
type BaseCQLListener struct{}

var _ CQLListener = &BaseCQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCqlFilter is called when production cqlFilter is entered.
func (s *BaseCQLListener) EnterCqlFilter(ctx *CqlFilterContext) {}

// ExitCqlFilter is called when production cqlFilter is exited.
func (s *BaseCQLListener) ExitCqlFilter(ctx *CqlFilterContext) {}

// EnterBooleanValueExpression is called when production booleanValueExpression is entered.
func (s *BaseCQLListener) EnterBooleanValueExpression(ctx *BooleanValueExpressionContext) {}

// ExitBooleanValueExpression is called when production booleanValueExpression is exited.
func (s *BaseCQLListener) ExitBooleanValueExpression(ctx *BooleanValueExpressionContext) {}

// EnterBooleanTerm is called when production booleanTerm is entered.
func (s *BaseCQLListener) EnterBooleanTerm(ctx *BooleanTermContext) {}

// ExitBooleanTerm is called when production booleanTerm is exited.
func (s *BaseCQLListener) ExitBooleanTerm(ctx *BooleanTermContext) {}

// EnterBooleanFactor is called when production booleanFactor is entered.
func (s *BaseCQLListener) EnterBooleanFactor(ctx *BooleanFactorContext) {}

// ExitBooleanFactor is called when production booleanFactor is exited.
func (s *BaseCQLListener) ExitBooleanFactor(ctx *BooleanFactorContext) {}

// EnterBooleanPrimary is called when production booleanPrimary is entered.
func (s *BaseCQLListener) EnterBooleanPrimary(ctx *BooleanPrimaryContext) {}

// ExitBooleanPrimary is called when production booleanPrimary is exited.
func (s *BaseCQLListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseCQLListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseCQLListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBinaryComparisonPredicate is called when production binaryComparisonPredicate is entered.
func (s *BaseCQLListener) EnterBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {}

// ExitBinaryComparisonPredicate is called when production binaryComparisonPredicate is exited.
func (s *BaseCQLListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseCQLListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseCQLListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseCQLListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseCQLListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseCQLListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseCQLListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterScalarExpression is called when production scalarExpression is entered.
func (s *BaseCQLListener) EnterScalarExpression(ctx *ScalarExpressionContext) {}

// ExitScalarExpression is called when production scalarExpression is exited.
func (s *BaseCQLListener) ExitScalarExpression(ctx *ScalarExpressionContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseCQLListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseCQLListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterCharacterLiteral is called when production characterLiteral is entered.
func (s *BaseCQLListener) EnterCharacterLiteral(ctx *CharacterLiteralContext) {}

// ExitCharacterLiteral is called when production characterLiteral is exited.
func (s *BaseCQLListener) ExitCharacterLiteral(ctx *CharacterLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseCQLListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseCQLListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseCQLListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseCQLListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterSpatialPredicate is called when production spatialPredicate is entered.
func (s *BaseCQLListener) EnterSpatialPredicate(ctx *SpatialPredicateContext) {}

// ExitSpatialPredicate is called when production spatialPredicate is exited.
func (s *BaseCQLListener) ExitSpatialPredicate(ctx *SpatialPredicateContext) {}

// EnterGeomExpression is called when production geomExpression is entered.
func (s *BaseCQLListener) EnterGeomExpression(ctx *GeomExpressionContext) {}

// ExitGeomExpression is called when production geomExpression is exited.
func (s *BaseCQLListener) ExitGeomExpression(ctx *GeomExpressionContext) {}

// EnterGeomLiteral is called when production geomLiteral is entered.
func (s *BaseCQLListener) EnterGeomLiteral(ctx *GeomLiteralContext) {}

// ExitGeomLiteral is called when production geomLiteral is exited.
func (s *BaseCQLListener) ExitGeomLiteral(ctx *GeomLiteralContext) {}

// EnterPoint is called when production point is entered.
func (s *BaseCQLListener) EnterPoint(ctx *PointContext) {}

// ExitPoint is called when production point is exited.
func (s *BaseCQLListener) ExitPoint(ctx *PointContext) {}

// EnterLinestring is called when production linestring is entered.
func (s *BaseCQLListener) EnterLinestring(ctx *LinestringContext) {}

// ExitLinestring is called when production linestring is exited.
func (s *BaseCQLListener) ExitLinestring(ctx *LinestringContext) {}

// EnterLinestringDef is called when production linestringDef is entered.
func (s *BaseCQLListener) EnterLinestringDef(ctx *LinestringDefContext) {}

// ExitLinestringDef is called when production linestringDef is exited.
func (s *BaseCQLListener) ExitLinestringDef(ctx *LinestringDefContext) {}

// EnterPolygon is called when production polygon is entered.
func (s *BaseCQLListener) EnterPolygon(ctx *PolygonContext) {}

// ExitPolygon is called when production polygon is exited.
func (s *BaseCQLListener) ExitPolygon(ctx *PolygonContext) {}

// EnterPolygonDef is called when production polygonDef is entered.
func (s *BaseCQLListener) EnterPolygonDef(ctx *PolygonDefContext) {}

// ExitPolygonDef is called when production polygonDef is exited.
func (s *BaseCQLListener) ExitPolygonDef(ctx *PolygonDefContext) {}

// EnterMultiPoint is called when production multiPoint is entered.
func (s *BaseCQLListener) EnterMultiPoint(ctx *MultiPointContext) {}

// ExitMultiPoint is called when production multiPoint is exited.
func (s *BaseCQLListener) ExitMultiPoint(ctx *MultiPointContext) {}

// EnterMultiLinestring is called when production multiLinestring is entered.
func (s *BaseCQLListener) EnterMultiLinestring(ctx *MultiLinestringContext) {}

// ExitMultiLinestring is called when production multiLinestring is exited.
func (s *BaseCQLListener) ExitMultiLinestring(ctx *MultiLinestringContext) {}

// EnterMultiPolygon is called when production multiPolygon is entered.
func (s *BaseCQLListener) EnterMultiPolygon(ctx *MultiPolygonContext) {}

// ExitMultiPolygon is called when production multiPolygon is exited.
func (s *BaseCQLListener) ExitMultiPolygon(ctx *MultiPolygonContext) {}

// EnterGeometryCollection is called when production geometryCollection is entered.
func (s *BaseCQLListener) EnterGeometryCollection(ctx *GeometryCollectionContext) {}

// ExitGeometryCollection is called when production geometryCollection is exited.
func (s *BaseCQLListener) ExitGeometryCollection(ctx *GeometryCollectionContext) {}

// EnterEnvelope is called when production envelope is entered.
func (s *BaseCQLListener) EnterEnvelope(ctx *EnvelopeContext) {}

// ExitEnvelope is called when production envelope is exited.
func (s *BaseCQLListener) ExitEnvelope(ctx *EnvelopeContext) {}

// EnterCoordinate is called when production coordinate is entered.
func (s *BaseCQLListener) EnterCoordinate(ctx *CoordinateContext) {}

// ExitCoordinate is called when production coordinate is exited.
func (s *BaseCQLListener) ExitCoordinate(ctx *CoordinateContext) {}

// EnterXCoord is called when production xCoord is entered.
func (s *BaseCQLListener) EnterXCoord(ctx *XCoordContext) {}

// ExitXCoord is called when production xCoord is exited.
func (s *BaseCQLListener) ExitXCoord(ctx *XCoordContext) {}

// EnterYCoord is called when production yCoord is entered.
func (s *BaseCQLListener) EnterYCoord(ctx *YCoordContext) {}

// ExitYCoord is called when production yCoord is exited.
func (s *BaseCQLListener) ExitYCoord(ctx *YCoordContext) {}

// EnterZCoord is called when production zCoord is entered.
func (s *BaseCQLListener) EnterZCoord(ctx *ZCoordContext) {}

// ExitZCoord is called when production zCoord is exited.
func (s *BaseCQLListener) ExitZCoord(ctx *ZCoordContext) {}

// EnterWestBoundLon is called when production westBoundLon is entered.
func (s *BaseCQLListener) EnterWestBoundLon(ctx *WestBoundLonContext) {}

// ExitWestBoundLon is called when production westBoundLon is exited.
func (s *BaseCQLListener) ExitWestBoundLon(ctx *WestBoundLonContext) {}

// EnterEastBoundLon is called when production eastBoundLon is entered.
func (s *BaseCQLListener) EnterEastBoundLon(ctx *EastBoundLonContext) {}

// ExitEastBoundLon is called when production eastBoundLon is exited.
func (s *BaseCQLListener) ExitEastBoundLon(ctx *EastBoundLonContext) {}

// EnterNorthBoundLat is called when production northBoundLat is entered.
func (s *BaseCQLListener) EnterNorthBoundLat(ctx *NorthBoundLatContext) {}

// ExitNorthBoundLat is called when production northBoundLat is exited.
func (s *BaseCQLListener) ExitNorthBoundLat(ctx *NorthBoundLatContext) {}

// EnterSouthBoundLat is called when production southBoundLat is entered.
func (s *BaseCQLListener) EnterSouthBoundLat(ctx *SouthBoundLatContext) {}

// ExitSouthBoundLat is called when production southBoundLat is exited.
func (s *BaseCQLListener) ExitSouthBoundLat(ctx *SouthBoundLatContext) {}

// EnterMinElev is called when production minElev is entered.
func (s *BaseCQLListener) EnterMinElev(ctx *MinElevContext) {}

// ExitMinElev is called when production minElev is exited.
func (s *BaseCQLListener) ExitMinElev(ctx *MinElevContext) {}

// EnterMaxElev is called when production maxElev is entered.
func (s *BaseCQLListener) EnterMaxElev(ctx *MaxElevContext) {}

// ExitMaxElev is called when production maxElev is exited.
func (s *BaseCQLListener) ExitMaxElev(ctx *MaxElevContext) {}

// EnterTemporalPredicate is called when production temporalPredicate is entered.
func (s *BaseCQLListener) EnterTemporalPredicate(ctx *TemporalPredicateContext) {}

// ExitTemporalPredicate is called when production temporalPredicate is exited.
func (s *BaseCQLListener) ExitTemporalPredicate(ctx *TemporalPredicateContext) {}

// EnterTemporalExpression is called when production temporalExpression is entered.
func (s *BaseCQLListener) EnterTemporalExpression(ctx *TemporalExpressionContext) {}

// ExitTemporalExpression is called when production temporalExpression is exited.
func (s *BaseCQLListener) ExitTemporalExpression(ctx *TemporalExpressionContext) {}

// EnterTemporalLiteral is called when production temporalLiteral is entered.
func (s *BaseCQLListener) EnterTemporalLiteral(ctx *TemporalLiteralContext) {}

// ExitTemporalLiteral is called when production temporalLiteral is exited.
func (s *BaseCQLListener) ExitTemporalLiteral(ctx *TemporalLiteralContext) {}

// EnterInPredicate is called when production inPredicate is entered.
func (s *BaseCQLListener) EnterInPredicate(ctx *InPredicateContext) {}

// ExitInPredicate is called when production inPredicate is exited.
func (s *BaseCQLListener) ExitInPredicate(ctx *InPredicateContext) {}
