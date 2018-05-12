package gogen

import (
	"context"
	"strings"
)

func writeExample(ctx context.Context, dir string, packageName string) error {
	f, err := newObjectFile(dir, "aaa_example_test")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(strings.Replace(`
package <package-name>_test

import (
	"fmt"
	"time"

	"github.com/romainmenke/schema-org-gen/<package-name>"
)

func ExampleAirline() {
	// Create an object of the correct Type
	airline := <package-name>.Airline{}

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

	// Some types need to be cast to "<package-name>" variants
	airline.FoundingDate = <package-name>.Date(foundingDate)
	airline.DissolutionDate = <package-name>.Date(dissolutionDate)

	// Define a nested object
	founder := &<package-name>.Person{}
	founder.Name = "John"
	airline.Founder = founder

	// Use "<package-name>.MarshalJSON" instead of "json.Marshal"
	// this ensures "@type" and "@context" are set correctly
	b, err := <package-name>.MarshalJSON(&airline)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	// Output: {"name":"Super Sonic","dissolutionDate":"2002-02-02","foundingDate":"2001-01-01","@context":"http://schema.org","@type":"Airline"}

}
`, "<package-name>", packageName, -1))

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
