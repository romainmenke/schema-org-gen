package schemaorggo_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/romainmenke/schema-org-gen/schemaorggo"
)

func ExampleAirline() {
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
	// Output: {"@context":"http://schema.org","@type":"Airline","dissolutionDate":["2002-02-02"],"founder":[{"@context":"http://schema.org","@type":"Person","name":["John"]}],"foundingDate":["2001-01-01"],"name":["Super Sonic"]}

}
