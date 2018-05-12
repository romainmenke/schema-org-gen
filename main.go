package main

import (
	"context"
	"encoding/json"
	"io"

	"github.com/romainmenke/schema-org-gen/internal/fetch"
	"github.com/romainmenke/schema-org-gen/internal/gogen"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

var verboseLog = true

func main() {
	useCacheRun()
	// cleanRun()
}

func useCacheRun() {

	typeMapCacheFile, err := typemap.NewOrExistingTypeMapCacheFile()
	if err != nil {
		panic(err)
	}
	defer typeMapCacheFile.Close()

	dec := json.NewDecoder(typeMapCacheFile)

	tm := &typemap.TypeMap{}
	err = dec.Decode(tm)
	if err == io.EOF {
		cleanRun()
		return
	}
	if err != nil {
		panic(err)
	}

	err = typeMapCacheFile.Close()
	if err != nil {
		panic(err)
	}

	if tm.UpdatedAt.IsZero() {
		cleanRun()
		return
	}

	err = gogen.Generate(context.Background(), tm, "./schemaorggo", "schemaorggo")
	if err != nil {
		panic(err)
	}
}

func cleanRun() {

	typeMapCacheFile, err := typemap.NewTypeMapCacheFile()
	if err != nil {
		panic(err)
	}
	defer typeMapCacheFile.Close()

	tm, err := fetch.TypeMap(context.Background())
	if err != nil {
		panic(err)
	}

	err = tm.Walk(context.Background(), fetch.Object)
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(typeMapCacheFile)
	err = enc.Encode(tm)
	if err != nil {
		panic(err)
	}

	err = typeMapCacheFile.Close()
	if err != nil {
		panic(err)
	}

	err = gogen.Generate(context.Background(), tm, "./schemaorggo", "schemaorggo")
	if err != nil {
		panic(err)
	}
}
