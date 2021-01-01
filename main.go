package main

import (
	"os"

	"github.com/mhelmetag/gosurf/cmd"

	"github.com/urfave/cli/v2"
)

const version = "2.1.0"

func main() {
	var srID string
	var sID string
	var d int

	var tID string
	var md int

	app := &cli.App{
		Name:    "gosurf",
		Usage:   "is there surf?",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:    "forecast",
				Aliases: []string{"f"},
				Usage:   "get a forecast for a subregion",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "subregion",
						Aliases:     []string{"sr"},
						Required:    true,
						Usage:       "subregion ID",
						Destination: &srID,
					},
					&cli.IntFlag{
						Name:        "days",
						Aliases:     []string{"d"},
						Value:       6,
						Usage:       "number of days to report (between 1 and 6)",
						Destination: &d,
					},
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
					&cli.StringFlag{
						Name:        "spot",
						Aliases:     []string{"s"},
						Required:    true,
						Usage:       "spot ID",
						Destination: &sID,
					},
					&cli.IntFlag{
						Name:        "days",
						Aliases:     []string{"d"},
						Value:       6,
						Usage:       "number of days to report (between 1 and 6)",
						Destination: &d,
					},
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
						Value:       "58f7ed51dadb30820bb38782", // default is Earth
						Usage:       "taxonomy ID",
						Destination: &tID,
					},
					&cli.IntFlag{
						Name:        "maxdepth",
						Aliases:     []string{"md"},
						Value:       0, // default to depth 0 for most searches
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
