package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/romainmenke/schema-org-gen/internal/ast"
	"github.com/romainmenke/schema-org-gen/internal/fetch"
	"github.com/romainmenke/schema-org-gen/internal/gengo"
	"github.com/romainmenke/schema-org-gen/internal/genphp"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	var (
		app = kingpin.New("schema-org-gen", "A generator for schema.org types")

		goCmd  = app.Command("go", "Generate for Go")
		phpCmd = app.Command("php", "Generate for php")

		tm  ast.Typemap
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

	tm, err = fetch.TypemapRDF(ctx)
	if err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case goCmd.FullCommand():
		err = gengo.Generate(ctx, tm, "./schemaorggo", "schemaorggo")
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
