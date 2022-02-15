/*
# CQL2 Antlr grammar, with small modifications.
# - Additions: ILIKE

# Build: in this dir: antlr -Dlanguage=Go -package cql CQL.g4 CqlLexer.g4
#
# See examples:
# https://portal.ogc.org/files/96288#cql-bnf
# https://github.com/interactive-instruments/xtraplatform-spatial/tree/master/xtraplatform-cql/src/main/antlr/de/ii/xtraplatform/cql/infra
*/
parser grammar CQL;
options { tokenVocab=CqlLexer;  }

/*
#=============================================================================#
# A CQL filter is a logically connected expression of one or more predicates.
#=============================================================================#
*/
cqlFilter : booleanValueExpression EOF;
booleanValueExpression : booleanTerm | booleanValueExpression OR booleanTerm;
booleanTerm : booleanFactor | booleanTerm AND booleanFactor;
booleanFactor : (NOT)? booleanPrimary;
booleanPrimary : predicate
                | LEFTPAREN booleanValueExpression RIGHTPAREN;

/*
#=============================================================================#
#  CQL supports scalar, spatial, temporal and existence predicates.
#=============================================================================#
*/

predicate : binaryComparisonPredicate
            | likePredicate
            | betweenPredicate
            | isNullPredicate
            | inPredicate
//            | spatialPredicate
//            | temporalPredicate
//            | arrayPredicate
//            | existencePredicate
            ;

/*
#=============================================================================#
# A comparison predicate evaluates if two scalar expression statisfy the
# specified comparison operator.  The comparion operators include an operator
# to evaluate regular expressions (LIKE), a range evaluation operator and
# an operator to test if a scalar expression is NULL or not.
#=============================================================================#
*/
binaryComparisonPredicate : scalarExpression ComparisonOperator scalarExpression;

likePredicate :  propertyName (NOT)? ( LIKE | ILIKE ) characterLiteral;

betweenPredicate : propertyName (NOT)? BETWEEN
                             scalarExpression AND scalarExpression ;
//                             (scalarExpression | temporalExpression) AND (scalarExpression | temporalExpression);

isNullPredicate : propertyName IS (NOT)? NULL;

/*
# A scalar expression is the property name, a chracter literal, a numeric
# literal or a function/method invocation that returns a scalar value.
*/
scalarExpression : propertyName
                    | characterLiteral
                    | numericLiteral
                    | booleanLiteral
//                    | function
                    /*| arithmeticExpression*/;

propertyName: Identifier;

characterLiteral: CharacterStringLiteral;

numericLiteral: NumericLiteral;

booleanLiteral: BooleanLiteral;

/*
#=============================================================================#
# A spatial predicate evaluates if two spatial expressions satisfy the
# specified spatial operator.
#=============================================================================#
*/

spatialPredicate :  SpatialOperator LEFTPAREN geomExpression COMMA geomExpression RIGHTPAREN;

/*
# A geometric expression is a property name of a geometry-valued property,
# a geometric literal (expressed as WKT) or a function that returns a
# geometric value.
*/
geomExpression : propertyName
               | geomLiteral
               /*| function*/;

/*
#=============================================================================#
# Definition of GEOMETRIC literals
#
# NOTE: This is basically BNF that define WKT encoding; it would be nice
#       to instead reference some normative BNF for WKT.
#=============================================================================#
*/

geomLiteral: point
             | linestring
             | polygon
             | multiPoint
             | multiLinestring
             | multiPolygon
             | geometryCollection
             | envelope;

point : POINT LEFTPAREN coordinate RIGHTPAREN;

linestring : LINESTRING linestringDef;

linestringDef: LEFTPAREN coordinate (COMMA coordinate)* RIGHTPAREN;

polygon : POLYGON polygonDef;

polygonDef : LEFTPAREN linestringDef (COMMA linestringDef)* RIGHTPAREN;

multiPoint : MULTIPOINT LEFTPAREN coordinate (COMMA coordinate)* RIGHTPAREN;

multiLinestring : MULTILINESTRING LEFTPAREN linestringDef (COMMA linestringDef)* RIGHTPAREN;

multiPolygon : MULTIPOLYGON LEFTPAREN polygonDef (COMMA polygonDef)* RIGHTPAREN;

geometryCollection : GEOMETRYCOLLECTION LEFTPAREN geomLiteral (COMMA geomLiteral)* RIGHTPAREN;

envelope: ENVELOPE LEFTPAREN westBoundLon COMMA southBoundLat COMMA (minElev COMMA)? eastBoundLon  COMMA northBoundLat (COMMA maxElev)? RIGHTPAREN;

coordinate : xCoord yCoord (zCoord)?;

xCoord : NumericLiteral;

yCoord : NumericLiteral;

zCoord : NumericLiteral;

westBoundLon : NumericLiteral;

eastBoundLon : NumericLiteral;

northBoundLat : NumericLiteral;

southBoundLat : NumericLiteral;

minElev : NumericLiteral;

maxElev : NumericLiteral;


/*
#=============================================================================#
# A temporal predicate evaluates if two temporal expressions satisfy the
# specified temporal operator.
#=============================================================================#
*/
//CHANGE: allow intervals with /
temporalPredicate : temporalExpression (TemporalOperator | ComparisonOperator) temporalExpression;

temporalExpression : propertyName
                   | temporalLiteral
                   /*| function*/;

temporalLiteral: TemporalLiteral;

/*
#=============================================================================#
# The IN predicate
#=============================================================================#
*/
//CHANGE: optional PropertyName for id filters
//CHANGE: added missing comma
/*
inPredicate : (propertyName | function)? (NOT)? IN LEFTPAREN ( characterLiteral |
                                            numericLiteral |
                                            geomLiteral |
                                            temporalLiteral  ) ( COMMA (characterLiteral |
                                                              numericLiteral |
                                                              geomLiteral |
                                                              temporalLiteral)  )* RIGHTPAREN;
*/
inPredicate : propertyName NOT? IN LEFTPAREN (
        characterLiteral (COMMA characterLiteral)*
        | numericLiteral (COMMA numericLiteral)*
    ) RIGHTPAREN;
