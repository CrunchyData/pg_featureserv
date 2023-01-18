/*
# CQL2 Antlr grammar, with small modifications.
# - Additions: ILIKE

# Build: in this dir: antlr -Dlanguage=Go -package cql CQLParser.g4 CqlLexer.g4
#
# Standard and examples:
# https://docs.ogc.org/DRAFTS/21-065.html#cql2-bnf
# https://github.com/interactive-instruments/xtraplatform-spatial/tree/master/xtraplatform-cql/src/main/antlr/de/ii/xtraplatform/cql/infra
*/
parser grammar CQLParser;
options { tokenVocab=CqlLexer; contextSuperClass=CqlContext; }

/*============================================================================
# A CQL filter is a logically connected expression of one or more predicates.
#============================================================================*/

cqlFilter : booleanExpression EOF;
booleanExpression : booleanTerm ( OR booleanTerm )?;
booleanTerm : booleanFactor ( AND booleanFactor )?;
booleanFactor : ( NOT )? booleanPrimary;
booleanPrimary : predicate
                | LEFTPAREN booleanExpression RIGHTPAREN;

/*============================================================================
#  CQL supports scalar, spatial, temporal and existence predicates.
#============================================================================*/

predicate : comparisonPredicate
          | spatialPredicate
          | distancePredicate
//          | temporalPredicate
//          | arrayPredicate
          ;

/*============================================================================
# A comparison predicate evaluates if two scalar expression statisfy the
# specified comparison operator.  The comparion operators include an operator
# to evaluate regular expressions (LIKE), a range evaluation operator and
# an operator to test if a scalar expression is NULL or not.
#============================================================================*/

comparisonPredicate : binaryComparisonPredicate
                    | isLikePredicate
                    | isBetweenPredicate
                    | isInListPredicate
                    | isNullPredicate;

binaryComparisonPredicate : scalarExpression ComparisonOperator scalarExpression;

isLikePredicate :  propertyName (NOT)? ( LIKE | ILIKE ) characterLiteral;

isBetweenPredicate : scalarExpression (NOT)? BETWEEN
                             scalarExpression AND scalarExpression ;

isInListPredicate : propertyName NOT? IN LEFTPAREN (
        characterLiteral (COMMA characterLiteral)*
        | numericLiteral (COMMA numericLiteral)*
    ) RIGHTPAREN;

isNullPredicate : propertyName IS (NOT)? NULL;

/*============================================================================
# Scalar expressions
#
# Note: does not enforce type consistency.
# That occurs when transpiled expression is evaluated.
#============================================================================*/

scalarExpression : scalarValue
                    | LEFTPAREN scalarExpression RIGHTPAREN
                    | scalarExpression ArithmeticOperator scalarExpression
                    ;

scalarValue : propertyName
            | characterLiteral
            | numericLiteral
            | booleanLiteral
            | temporalLiteral
//                   | function
             ;

propertyName: Identifier;
characterLiteral: CharacterStringLiteral;
numericLiteral: NumericLiteral;
booleanLiteral: BooleanLiteral;
temporalLiteral: TemporalLiteral;

/*============================================================================
# A spatial predicate evaluates if two spatial expressions satisfy the
# specified spatial operator.
#============================================================================*/

spatialPredicate :  SpatialOperator LEFTPAREN geomExpression COMMA geomExpression RIGHTPAREN;

distancePredicate :  DistanceOperator LEFTPAREN geomExpression COMMA geomExpression COMMA NumericLiteral RIGHTPAREN;

/*
# A geometric expression is a property name of a geometry-valued property,
# a geometric literal (expressed as WKT) or a function that returns a
# geometric value.
*/
geomExpression : propertyName
               | geomLiteral
               /*| function*/;

/*============================================================================
# Definition of GEOMETRIC literals
#============================================================================*/

geomLiteral: point
             | linestring
             | polygon
             | multiPoint
             | multiLinestring
             | multiPolygon
             | geometryCollection
             | envelope;

point : POINT pointList;
pointList : LEFTPAREN coordinate RIGHTPAREN;
linestring : LINESTRING coordList;
polygon : POLYGON polygonDef;
polygonDef : LEFTPAREN coordList (COMMA coordList)* RIGHTPAREN;
multiPoint : MULTIPOINT LEFTPAREN pointList (COMMA pointList)* RIGHTPAREN;
multiLinestring : MULTILINESTRING LEFTPAREN coordList (COMMA coordList)* RIGHTPAREN;
multiPolygon : MULTIPOLYGON LEFTPAREN polygonDef (COMMA polygonDef)* RIGHTPAREN;
geometryCollection : GEOMETRYCOLLECTION LEFTPAREN geomLiteral (COMMA geomLiteral)* RIGHTPAREN;

envelope: ENVELOPE LEFTPAREN NumericLiteral COMMA NumericLiteral COMMA NumericLiteral  COMMA NumericLiteral RIGHTPAREN;

coordList: LEFTPAREN coordinate (COMMA coordinate)* RIGHTPAREN;
coordinate : NumericLiteral NumericLiteral;
