package main

import (
	"fmt"
	"github.com/michcioperz/raven_util"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-current-ref"
	app.Usage = "extract current commit"
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
		},
		cli.StringFlag{
			Name:  "path",
			Value: pwd,
		},
	}
	app.Action = func(c *cli.Context) error {
		var l *log.Logger
		if c.Bool("verbose") {
			l = log.New(os.Stderr, "researching: ", log.LstdFlags)
		} else {
			l = nil
		}
		fmt.Println(raven_util.ExtractCurrentRelease(c.String("path"), l))
		return nil
	}
	app.Run(os.Args)
}
