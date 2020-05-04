package main

import (
	"fmt"
	"foo/internal/commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "foo"
	app.Version = "0.0.1"
	app.Usage = "to write fast and distributable command line applications in an expressive way."
	app.Commands = []cli.Command{
		commands.StartCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
