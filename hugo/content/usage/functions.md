---
title: "Functions"
date:
draft: false
weight: 200
---

A powerful feature of Postgres is the ability to create
[user-defined functions](https://www.postgresql.org/docs/current/xfunc.html).
Functions allow encapsulating complex logic behind a simple
interface (namely that of providing some input arguments
and getting output as a set of records).
This makes them easy to publish via a simple web API.

Functions can execute any data processing that is
possible to perform with Postgres and PostGIS.
They can return either spatial or non-spatial results
(as GeoJSON or plain JSON).
They thus provide a powerful extension to the capabilities of
the `pg_featureserv` API.

Potential uses of functions include:

* Query a spatial database table or view with custom SQL
  (which can include things such as special filters or aggregation)
* Query a non-spatial table or view to return data objects or a summary record
* Generate spatial data controlled by a set of parameters
* Expose a geometric computation,
  by accepting a geometric input value and returning a single record containing the result
* Functions can even be used to update data (as long as appropriate security is in place).

## Publish Database Functions

The service can publish any function which returns a set of rows
using the return type `SETOF record`
or the equivalent (and more standard) `TABLE`
(see the Postgres manual section on [set-returning functions](https://www.postgresql.org/docs/current/xfunc-sql.html#XFUNC-SQL-FUNCTIONS-RETURNING-SET).)
Because there are usually many functions in a Postgres database,
the service only publishes functions defined in the `postgisftw` schema.

A function specifies zero or more input parameters.
An input parameter can be of any Postgres type
which has a cast from a text representation.  This includes the PostGIS geometry
and geography types, which support text representations of
[WKT or WKB](https://postgis.net/docs/manual-3.0/using_postgis_dbmanagement.html#OpenGISWKBWKT).
Input parameter names are exposed as query parameters,
so you should avoid using names which are existing API qeuery parameters.

A function must return a set of records containing one or more
columns, of any Postgres type.
A **spatial function** is one which returns a column of type `geometry` or `geography`.
Output from spatial functions is returned as GeoJSON datasets.
Output from non-spatial functions is returned as JSON datasets.

Geometry values returned by a function can be in any coordinate system,
but must have their SRID set to the appropriate value.
If required, they are reprojected to geographic coordinates (SRID = 4326) in the output GeoJSON.
If geometry is queried from an existing table the SRID may already be set;
otherwise the function should set it explicitly.

The example below illustrates
the basic structure of a spatial set-returning function.
See the [Examples](/examples/) section for further examples.

#### Example of a spatial function

This example is about the simplest possible spatial function,
but it illustrates the most important concepts.
It returns a filtered subset of a table ([ne_50m_admin_0_countries](https://www.naturalearthdata.com/http//www.naturalearthdata.com/download/50m/cultural/ne_50m_admin_0_countries.zip) which is in [EPSG:4326](https://epsg.io/4326)).
The filter in this case is the first letters of the country name.

The `name_prefix` parameter includes a **default value**: this is useful for clients
(like the preview interface for this service)
that read arbitrary function definitions and need a default value to fill into interface fields.

```sql
CREATE OR REPLACE FUNCTION postgisftw.countries_name(
	name_prefix text DEFAULT 'A')
RETURNS TABLE(name text, abbrev text, continent text, geom geometry)
AS $$
BEGIN
	RETURN QUERY
		SELECT t.name::text,
            t.abbrev::text,
            t.continent::text,
            t.geom
    FROM ne.admin_0_countries t
    WHERE t.name ILIKE name_prefix || '%';
END;
$$
LANGUAGE 'plpgsql' STABLE PARALLEL SAFE;

COMMENT ON FUNCTION postgisftw.countries_name IS 'Filters the countries table by the initial letters of the name using the "name_prefix" parameter.';
```

Aspects to note are:

* The function is defined in the `postgisftw` schema.
* it has a single input parameter `name_prefix`, with the `DEFAULT` value 'A'.
* It returns a table (set) of type `(name text, geom geometry)`.
* The function body is a simple `SELECT` query which uses the input parameter as part of a `ILIKE` filter,
  and returns a column list matching the output table definition.
* The geometry values are assumed to carry an SRID specified in the queried table.
* The function "[volatility](https://www.postgresql.org/docs/current/xfunc-volatility.html)" is declared as `STABLE` because within a transaction context multiple calls with the same inputs will return the same outputs. It is not marked as `IMMUTABLE` because changes in the base table can change the outputs over time, even for the same inputs.
* The function is declared as `PARALLEL SAFE` because it doesn't depend on any state that might be altered by making multiple concurrent calls to the function.

The function can be called via the API by providing a value for the `name_prefix` parameter
(which could be omitted, due to the presence of a default value):

```
http://localhost:9000/functions/countries_name/items?name_prefix=T
```

The response is a GeoJSON document containing the 13 countries starting with the letter 'T'.

As with feature collections, available functions can be listed,
and each function can supply metadata describing it.

## List Functions

The path `/functions` returns a JSON document
containing a list of the functions available in the service.

#### Example
```
http://localhost:9000/functions
```

Each listed function is described by a subset of its metadata,
including its id and description.
A list of links provide URLs for accessing:

* `self` - the function metadata
* `alternate` - the function metadata as an HTML view
* `items` - the function data items


## Describe Function metadata

The path `/functions/{funid}` returns a JSON object describing
the metadata for a database function.
`{funid}` is the name of the function.
It is not schema-qualified, because functions
are published from only one schema.

#### Example
```
http://localhost:9000/functions/geonames_geom
```

The response is a JSON document ontaining metadata about the function, including:

* The function description
* A list of the input parameters, described by name, type, description, and default value (if any)
* A list of the properties and their JSON types

A list of links provides URLs for accessing:

* `self` - the function metadata
* `alternate` - the function metadata as an HTML view
* `items` - the data items returned by querying the function
