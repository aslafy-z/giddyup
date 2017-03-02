package main

import (
	"os"

	giddyupApp "github.com/aslafy-z/giddyup/app"
	"github.com/aslafy-z/giddyup/version"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = version.VERSION
	app.Name = "giddyup"
	app.Usage = "Entrypoint functions for Rancher"

	app.Commands = []cli.Command{
		giddyupApp.IPCommand(),
		giddyupApp.LeaderCommand(),
		giddyupApp.ServiceCommand(),
		giddyupApp.HealthCommand(),
		giddyupApp.ProbeCommand(),
	}

	app.Run(os.Args)
}
