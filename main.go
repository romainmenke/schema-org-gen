package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/romainmenke/schema-org-gen/internal/fetch"
	"github.com/romainmenke/schema-org-gen/internal/gengo"
	"github.com/romainmenke/schema-org-gen/internal/genjson"
	"github.com/romainmenke/schema-org-gen/internal/genphp"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	var (
		app   = kingpin.New("schema-org-gen", "A generator for schema.org types")
		clean = app.Flag("clean", "Fetch the typemap from schema.org before generating.").Bool()

		goCmd   = app.Command("go", "Generate for Go")
		phpCmd  = app.Command("php", "Generate for php")
		jsonCmd = app.Command("json", "Generate JSON examples")

		tm  *typemap.TypeMap
		err error
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	go func() {
		signal_chan := make(chan os.Signal, 1)
		signal.Notify(
			signal_chan,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)

		for {
			s := <-signal_chan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				cancel()
				return
			}
		}

	}()

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	if clean != nil && *clean {
		tm, err = tmFromSchemaOrg(ctx)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		tm, err = tmFromCache(ctx)
		if err != nil {
			log.Println(err)
			return
		}
	}

	switch cmd {
	case goCmd.FullCommand():
		err = gengo.Generate(ctx, tm, "./schemaorggo", "schemaorggo")
		if err != nil {
			log.Fatal(err)
		}

	case jsonCmd.FullCommand():
		err = genjson.Generate(ctx, tm, "./schemaorgjson")
		if err != nil {
			log.Fatal(err)
		}

	case phpCmd.FullCommand():
		err = genphp.Generate(ctx, tm, "./schemaorgphp", "schemaorgphp")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func tmFromCache(ctx context.Context) (*typemap.TypeMap, error) {
	typeMapCacheFile, err := typemap.NewOrExistingTypeMapCacheFile()
	if err != nil {
		return nil, err
	}
	defer typeMapCacheFile.Close()

	dec := json.NewDecoder(typeMapCacheFile)

	tm := &typemap.TypeMap{}
	err = dec.Decode(tm)
	if err == io.EOF {
		return tmFromSchemaOrg(ctx)
	}
	if err != nil {
		return nil, err
	}

	err = typeMapCacheFile.Close()
	if err != nil {
		return nil, err
	}

	if tm.UpdatedAt.IsZero() {
		return tmFromSchemaOrg(ctx)
	}

	return tm, nil
}

func tmFromSchemaOrg(ctx context.Context) (*typemap.TypeMap, error) {

	typeMapCacheFile, err := typemap.NewTypeMapCacheFile()
	if err != nil {
		return nil, err
	}
	defer typeMapCacheFile.Close()

	tm, err := fetch.TypeMap(ctx)
	if err != nil {
		return nil, err
	}

	err = tm.Walk(ctx, fetch.Object)
	if err != nil {
		return nil, err
	}

	enc := json.NewEncoder(typeMapCacheFile)
	err = enc.Encode(tm)
	if err != nil {
		return nil, err
	}

	err = typeMapCacheFile.Close()
	if err != nil {
		return nil, err
	}

	return tm, nil
}
