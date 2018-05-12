package schemaorggo_test

import (
	"fmt"
	"time"

	"github.com/romainmenke/schema-org-gen/schemaorggo"
)

func ExampleAirline() {
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
