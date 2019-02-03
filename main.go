package main

import (
	"os"

	"github.com/mhelmetag/gosurf/cmd"

	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/urfave/cli.v2"
	"gopkg.in/urfave/cli.v2/altsrc"
)

const version = "0.0.5"

func main() {
	var aID string
	var rID string
	var srID string
	var pType string
	var d int

	cfgFilepath, _ := homedir.Expand("~/.gosurf.yml")
	flags := []cli.Flag{
		altsrc.NewStringFlag(
			&cli.StringFlag{
				Name:        "area",
				Aliases:     []string{"a"},
				Value:       "4716",
				Usage:       "area ID for a region or subregion",
				Destination: &aID,
			},
		),
		altsrc.NewStringFlag(
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"r"},
				Value:       "2081",
				Usage:       "region ID for a subregion",
				Destination: &rID,
			},
		),
		altsrc.NewStringFlag(
			&cli.StringFlag{
				Name:        "subregion",
				Aliases:     []string{"s"},
				Value:       "2141",
				Usage:       "subregion ID",
				Destination: &srID,
			},
		),
		altsrc.NewIntFlag(
			&cli.IntFlag{
				Name:        "days",
				Aliases:     []string{"d"},
				Value:       7,
				Usage:       "number of days to report (between 1 and 8)",
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
				Name:    "places",
				Aliases: []string{"p"},
				Usage:   "search for places that surfline supports",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "placetype",
						Aliases:     []string{"t"},
						Value:       "areas",
						Usage:       "which place type to search fo (areas, regions, subregions)",
						Destination: &pType,
					},
				},
				Action: func(c *cli.Context) error {
					cmd.Places(pType, aID, rID)

					return nil
				},
			},
			{
				Name:    "forecast",
				Aliases: []string{"f"},
				Usage:   "get a forecast for a subregion",
				Action: func(c *cli.Context) error {
					cmd.Forecast(aID, rID, srID, d)

					return nil
				},
			},
			{
				Name:    "tide",
				Aliases: []string{"t"},
				Usage:   "get the tides for a subregion",
				Action: func(c *cli.Context) error {
					cmd.Tide(aID, rID, srID, d)

					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
