---
title: "Filtering with CQL"
date:
draft: false
weight: 175
---

The features returned by queries can be filtered using
the `filter` query parameter with an expression written using
the Common Query Language (CQL).
CQL expressions return a value of `true` or `false`.
Only features which evaluate to `true` are returned.
In `pg_featureserv` the filter expression is evaluated by the database,
so it can take advantage of indexes to make filter evaluation very efficient.

This section describes the CQL query language supported by `pg_featureserv`.

## Property and Literal Values

The basic elements of filter expressions are values obtained
from feature collection properties and literals.
Properties are referred to by name, and literals can be
numbers, boolean or text values.

#### Example
```
propname
1.234
true
'a text value'
```

## Comparisons

Values can be compared using conditional operators:
```
a = b   a <> b   a > b   a >= b   a < b   a <= b
```

### Example
```
pop_est >= 1000000
name = 'Finland'
```

## BETWEEN predicate

The `BETWEEN` predicate tests if a value lies in a range defined by start and end values (inclusive):
```
property [NOT] BETWEEN a AND b
```

### Example
```
pop_est BETWEEN 100000 AND 1000000
name NOT BETWEEN 'Chile' AND 'Denmark'
```

## IN predicate
The `IN` predicate tests if a value lies in a list of constant values.
```
property [NOT] IN ( val1, val2, ... )
```

### Example
```
id IN (1,2,3)
name IN ('Chile', 'Kenya', 'Denmark')
```

## LIKE predicate
The `LIKE` predicate tests if a text value matches a pattern.
The character `%` is a wildcard.
(Note that this may need to be URL-encoded as `%25`.)
`ILIKE` can be used for case-independent matching.

```
property [NOT] LIKE | ILIKE pattern
```

### Example
```
name LIKE 'Ch%'
continent ILIKE '%america'
```

## IS NULL predicate
The `IS NULL` predicate tests if a property value is (or is not) null.
```
property IS [NOT] NULL
```

### Example
```
name IS NULL
```

## Boolean combinations
Comparisons and predicates can be combined with the
boolean operators `AND`, `OR` and `NOT`.
Operators are evaluated in the order 'NOT', 'AND', 'OR'.
Evaluation order can be controlled by enclosing
subexpressions in parentheses.

### Example
```
(continent = 'Europe' OR continent = 'Afica') AND pop_est < 1000000
```
