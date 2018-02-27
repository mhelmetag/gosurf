package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mhelmetag/surflinef"
	"github.com/mhelmetag/surfliner"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

const version = "0.0.2"

func main() {
	var pType string
	var aID string
	var rID string
	var srID string

	app := &cli.App{
		Name:    "gosurf",
		Usage:   "is there surf?",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "a",
				Value:       "4716",
				Usage:       "area ID for a region or subregion",
				Destination: &aID,
			},
			&cli.StringFlag{
				Name:        "r",
				Value:       "2081",
				Usage:       "region ID for a subregion",
				Destination: &rID,
			},
			&cli.StringFlag{
				Name:        "sr",
				Value:       "2141",
				Usage:       "subregion ID",
				Destination: &srID,
			},
		},
		Commands: []cli.Command{
			{
				Name:    "places",
				Aliases: []string{"p"},
				Usage:   "search for places that surfline supports",
				Action: func(c *cli.Context) error {
					search(pType, aID, rID)

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "pt",
						Value:       "areas",
						Usage:       "which place type to search fo (areas, regions, subregions)",
						Destination: &pType,
					},
				},
			},
			{
				Name:    "forecast",
				Aliases: []string{"f"},
				Usage:   "get a forecast for a subregion",
				Action: func(c *cli.Context) error {
					err := validateSubRegion(aID, rID, srID)
					if err != nil {
						fmt.Println("Couldn't find a subregion with the given ID")

						return nil
					}

					forecast(aID, rID, srID)

					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}

func search(pType string, aID string, rID string) {
	switch pType {
	case "areas":
		areas()
	case "regions":
		regions(aID)
	case "subregions":
		subRegions(aID, rID)
	default:
		fmt.Println("No valid type selected. Choose from areas, regions or subregions")
	}
}

func forecast(aID string, rID string, srID string) {
	c, err := surflinef.DefaultClient()
	if err != nil {
		fmt.Println("Error while building the SurflineF client")

		return
	}

	q := surflinef.Query{
		Resources:    []string{"analysis"},
		Days:         7,
		Units:        "e",
		FullAnalysis: true,
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Println("Error building query for Forecast")

		return
	}

	f, err := c.GetForecast(srID, qs)
	if err != nil {
		fmt.Println("Error fetching Forecast")

		return
	}

	analysisToTable(f.Analysis)
}

func areas() {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("Error while building the SurflineR client")

		return
	}

	as, err := c.Areas()
	if err != nil {
		fmt.Println("Error while fetching Areas")

		return
	}

	placeToTable(as)
}

func regions(aID string) {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("Error while building the SurflineR client")
	}

	rs, err := c.Regions(aID)
	if err != nil {
		fmt.Println("Error while fetching Regions")
	}

	placeToTable(rs)
}

func subRegions(aID string, rID string) {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("Error while building the SurflineR client")
	}

	rs, err := c.SubRegions(aID, rID)
	if err != nil {
		fmt.Println("Error while fetching Regions")
	}

	placeToTable(rs)
}

func validateSubRegion(aID string, rID string, srID string) error {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("Error while building the SurflineR client")
	}

	_, err = c.SubRegion(aID, rID, srID)

	return err
}

func placeToTable(ps []surfliner.Place) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})

	for i := range ps {
		p := ps[i]
		table.Append([]string{p.ID, p.Name})
	}

	table.Render()
}

func analysisToTable(a surflinef.Analysis) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Condition", "Report"})
	rs := analysisReports(a)
	t := time.Now()

	for i := range a.GeneralCondition {
		ts := fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())
		table.Append([]string{ts, a.GeneralCondition[i], rs[i]})
		t = t.AddDate(0, 0, 1)
	}

	table.Render()
}

func analysisReports(a surflinef.Analysis) []string {
	var rs []string

	for i := range a.GeneralCondition {
		r := fmt.Sprintf("%s-%sft. - %s", a.SurfMin[i], a.SurfMax[i], a.SurfText[i])

		rs = append(rs, r)
	}

	return rs
}
