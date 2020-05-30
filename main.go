package main

import (
	"os"

	"github.com/mhelmetag/gosurf/cmd"

	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/urfave/cli.v2"
	"gopkg.in/urfave/cli.v2/altsrc"
)

const version = "2.0.0"

func main() {
	var srID string
	var d int

	cfgFilepath, _ := homedir.Expand("~/.gosurf.yml")
	flags := []cli.Flag{
		altsrc.NewStringFlag(
			&cli.StringFlag{
				Name:        "subregion",
				Aliases:     []string{"s"},
				Value:       "58581a836630e24c44878fd4",
				Usage:       "subregion ID",
				Destination: &srID,
			},
		),
		altsrc.NewIntFlag(
			&cli.IntFlag{
				Name:        "days",
				Aliases:     []string{"d"},
				Value:       6,
				Usage:       "number of days to report (between 1 and 6)",
				Destination: &d,
			},
		),
		&cli.StringFlag{
			Name:    "configfile",
			Aliases: []string{"c"},
			Value:   cfgFilepath,
			Usage:   "application config filepath",
		},
	}

	var beforeFunc cli.BeforeFunc
	_, err := os.Stat(cfgFilepath)
	if err == nil {
		beforeFunc = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("configfile"))
	} else {
		beforeFunc = func(c *cli.Context) error {
			return nil
		}
	}

	app := &cli.App{
		Name:    "gosurf",
		Usage:   "is there surf?",
		Version: version,
		Before:  beforeFunc,
		Flags:   flags,
		Commands: []*cli.Command{
			{
				Name:    "forecast",
				Aliases: []string{"f"},
				Usage:   "get a forecast for a subregion",
				Action: func(c *cli.Context) error {
					cmd.Forecast(srID, d)

					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
