package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
	"runtime"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "setbench"
	app.Usage = "Benchmarks different properties of set implementations."

	memplotFlags, memplotAction := memplotCommand()

	app.Commands = []cli.Command{
		{
			Name:   "memplot",
			Usage:  "Plots memory usage over time as keys are inserted in a set.",
			Flags:  memplotFlags,
			Action: memplotAction,
		},
	}

	return app
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetPrefix("[setbench] ")
	log.SetFlags(log.Lshortfile)

	err := NewApp().Run(os.Args)
	if err != nil {
		log.Fatalf("Running app: %v", err)
	}
}
