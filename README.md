# Schema.org Type Generator

[![GoDoc](https://godoc.org/github.com/romainmenke/schema-org-gen/schemaorggo?status.svg)](https://godoc.org/github.com/romainmenke/schema-org-gen/schemaorggo)

* [schema.org](https://schema.org)
* [google structured data](https://developers.google.com/search/docs/guides/intro-structured-data)

This generator is for those who like type safety and do not want to lookup the possible keys for Google structured data and other implementations of schema.org.

# How it works :

The code generator scrapes schema.org collecting a type map, field definitions, data types, comments and notes. All this is converted to one of the supported languages. The result is a package which you import in your own project and use to write `ld+json`'s.

### Generate

```
usage: schema-org-gen [<flags>] <command> [<args> ...]

A generator for schema.org types

Flags:
  --help   Show context-sensitive help (also try --help-long and --help-man).
  --clean  Fetch the typemap from schema.org before generating.

Commands:
  help [<command>...]
    Show help.

  go
    Generate for Go

  json
    Generate JSON examples
```

# Using the generated types

### Go

`go get github.com/romainmenke/schema-org-gen/schemaorggo`

```go
package main

import (
	"fmt"
	"time"

	"github.com/romainmenke/schema-org-gen/schemaorggo"
)

func main() {
	// Create an object of the correct Type
	airline := schemaorggo.Airline{}

	// Set some properties
	airline.Name = "Super Sonic"

	foundingDate, err := time.Parse("2006-01-02", "2001-01-01")
	if err != nil {
		panic(err)
	}
	dissolutionDate, err := time.Parse("2006-01-02", "2002-02-02")
	if err != nil {
		panic(err)
	}

	// Some types need to be cast to "schemaorggo" variants
	airline.FoundingDate = schemaorggo.Date(foundingDate)
	airline.DissolutionDate = schemaorggo.Date(dissolutionDate)

	// Define a nested object
	founder := &schemaorggo.Person{}
	founder.Name = "John"
	airline.Founder = founder

	// Use "schemaorggo.MarshalJSON" instead of "json.Marshal"
	// this ensures "@type" and "@context" are set correctly
	b, err := schemaorggo.MarshalJSON(&airline)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	// Output: {"name":"Super Sonic","dissolutionDate":"2002-02-02","founder":{"name":"John","@context":"http://schema.org","@type":"Person","birthDate":null,"deathDate":null},"foundingDate":"2001-01-01","@context":"http://schema.org","@type":"Airline"}

}
```
