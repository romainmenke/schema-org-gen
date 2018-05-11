package main

import (
	"context"
	"encoding/json"
	"io"
)

var verboseLog = true

func main() {
	// useCacheRun()
	cleanRun()
}

func useCacheRun() {

	typeMapCacheFile, err := newOrExistingTypeMapCacheFile()
	if err != nil {
		panic(err)
	}
	defer typeMapCacheFile.Close()

	dec := json.NewDecoder(typeMapCacheFile)

	tm := &TypeMap{}
	err = dec.Decode(tm)
	if err == io.EOF {
		cleanRun()
		return
	}
	if err != nil {
		panic(err)
	}

	if tm.UpdatedAt.IsZero() {
		cleanRun()
		return
	}

	err = writeGoDataTypes(context.Background())
	if err != nil {
		panic(err)
	}

	err = tm.walk(context.Background(), goTypeFile)
	if err != nil {
		panic(err)
	}

	err = typeMapCacheFile.Close()
	if err != nil {
		panic(err)
	}
}

func cleanRun() {

	typeMapCacheFile, err := newTypeMapCacheFile()
	if err != nil {
		panic(err)
	}
	defer typeMapCacheFile.Close()

	tm, err := getTypeMap(context.Background())
	if err != nil {
		panic(err)
	}

	err = tm.walk(context.Background(), getObject)
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

	err = writeGoDataTypes(context.Background())
	if err != nil {
		panic(err)
	}

	err = tm.walk(context.Background(), goTypeFile)
	if err != nil {
		panic(err)
	}
}
