package main

import (
	"os"

	"github.com/mhelmetag/gosurf/cmd"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

const version = "2.1.0"

func main() {
	var srID string
	var sID string
	var d int

	var tID string
	var md int

	cfgFilepath, _ := homedir.Expand("~/.gosurf.yml")
	cfgFilePathFlag := &cli.StringFlag{
		Name:    "configfile",
		Aliases: []string{"c"},
		Value:   cfgFilepath,
		Usage:   "application config filepath",
	}

	globalFlags := []cli.Flag{
		cfgFilePathFlag,
	}

	subregionFlag := altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:        "subregion",
			Aliases:     []string{"sr"},
			Required:    true,
			Usage:       "subregion ID",
			Destination: &srID,
		},
	)
	spotFlag := altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:        "spot",
			Aliases:     []string{"s"},
			Required:    true,
			Usage:       "spot ID",
			Destination: &sID,
		},
	)
	daysFlag := altsrc.NewIntFlag(
		&cli.IntFlag{
			Name:        "days",
			Aliases:     []string{"d"},
			Value:       6,
			Usage:       "number of days to report (between 1 and 6)",
			Destination: &d,
		},
	)

	allFlags := []cli.Flag{
		cfgFilePathFlag,
		subregionFlag,
		spotFlag,
		daysFlag,
	}

	// TODO - fix this before release
	var beforeFunc cli.BeforeFunc
	_, err := os.Stat(cfgFilepath)
	if err == nil {
		beforeFunc = altsrc.InitInputSourceWithContext(allFlags, altsrc.NewYamlSourceFromFlagFunc("configfile"))
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
		Flags:   globalFlags,
		Commands: []*cli.Command{
			{
				Name:    "forecast",
				Aliases: []string{"f"},
				Usage:   "get a forecast for a subregion",
				Flags: []cli.Flag{
					subregionFlag,
					daysFlag,
				},
				Action: func(c *cli.Context) error {
					cmd.Forecast(srID, d)

					return nil
				},
			},
			{
				Name:    "tide",
				Aliases: []string{"t"},
				Usage:   "get a tide for a spot",
				Flags: []cli.Flag{
					spotFlag,
					daysFlag,
				},
				Action: func(c *cli.Context) error {
					cmd.Tide(sID, d)

					return nil
				},
			},
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "search through the taxonomy tree",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "taxonomy",
						Aliases:     []string{"t"},
						Required:    true,
						Usage:       "taxonomy ID",
						Destination: &tID,
					},
					&cli.IntFlag{
						Name:        "maxdepth",
						Aliases:     []string{"md"},
						Value:       0,
						Usage:       "max depth for the tree search",
						Destination: &md,
					},
				},
				Action: func(c *cli.Context) error {
					cmd.Search(tID, md)

					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
