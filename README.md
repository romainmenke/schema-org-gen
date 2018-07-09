# Schema.org Type Generator

[![GoDoc](https://godoc.org/github.com/romainmenke/schema-org-gen/schemaorggo?status.svg)](https://godoc.org/github.com/romainmenke/schema-org-gen/schemaorggo)

* [schema.org](https://schema.org)
* [google structured data](https://developers.google.com/search/docs/guides/intro-structured-data)

This generator is for those who like type safety and do not want to lookup the possible keys for Google structured data and other implementations of schema.org.

# How it works :

The code generator scrapes schema.org collecting a type map, field definitions, data types, comments and notes. All this is converted to one of the supported languages. The result is a package which you import in your own project and use to write `ld+json`'s.

### Generate

```
$ schema-org-gen --help
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

  php
    Generate for php


```

# Using the generated types

### Go

*The generate package is a bit large and it might be better in some cases to copy/paste the types you need.*

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
	airline.Name = []string{"Super Sonic"}

	foundingDate, err := time.Parse("2006-01-02", "2001-01-01")
	if err != nil {
		panic(err)
	}
	dissolutionDate, err := time.Parse("2006-01-02", "2002-02-02")
	if err != nil {
		panic(err)
	}

	// Some types need to be cast to "schemaorggo" variants
	airline.FoundingDate = []schemaorggo.Date{
		schemaorggo.Date(foundingDate),
	}
	airline.DissolutionDate = []schemaorggo.Date{
		schemaorggo.Date(dissolutionDate),
	}

	// Define a nested object
	founder := &schemaorggo.Person{}
	founder.Name = []string{"John"}
	airline.Founder = []*schemaorggo.Person{
		founder,
	}

	b, err := json.Marshal(airline)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	// Output: {"@context":"http://schema.org","@type":"Airline","dissolutionDate":"2002-02-02","founder":{"@context":"http://schema.org","@type":"Person","name":"John"},"foundingDate":"2001-01-01","name":"Super Sonic"}
}
```
------
