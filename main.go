package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mhelmetag/surflinef"
	"github.com/mhelmetag/surfliner"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
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
				Usage:       "number of days to report (between 1 and 15)",
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
				Action: func(c *cli.Context) error {
					search(pType, aID, rID)

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "placetype",
						Aliases:     []string{"t"},
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

					forecast(aID, rID, srID, d)

					return nil
				},
			},
			{
				Name:    "tide",
				Aliases: []string{"t"},
				Usage:   "get the tides for a subregion",
				Action: func(c *cli.Context) error {
					err := validateSubRegion(aID, rID, srID)
					if err != nil {
						fmt.Println("Couldn't find a subregion with the given ID")

						return nil
					}

					tide(aID, rID, srID, d)

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

func forecast(aID string, rID string, srID string, d int) {
	c, err := surflinef.DefaultClient()
	if err != nil {
		fmt.Println("There was an error while building the SurflineF client")

		return
	}

	if !validDayAmount(d) {
		fmt.Println("The number of days to report can only be between 1 and 15")

		return
	}

	q := surflinef.Query{
		Resources:    []string{"analysis"},
		Days:         d,
		Units:        "e",
		FullAnalysis: true,
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Println("There was an error while building a query for the forecast")

		return
	}

	f, err := c.GetForecast(srID, qs)
	if err != nil {
		fmt.Println("There was an error while fetching the forecast")

		return
	}

	analysisToTable(f.Analysis)
}

func tide(aID string, rID string, srID string, d int) {
	c, err := surflinef.DefaultClient()
	if err != nil {
		fmt.Println("There was an error while building the SurflineF client")

		return
	}

	if !validDayAmount(d) {
		fmt.Println("The number of days to report can only be between 1 and 15")

		return
	}

	q := surflinef.Query{
		Resources:    []string{"tide"},
		Days:         d,
		Units:        "e",
		FullAnalysis: true,
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Println("There was an error while building a query for the tide")

		return
	}

	f, err := c.GetForecast(srID, qs)
	if err != nil {
		fmt.Println("There was an error while fetching the tide")

		return
	}

	tideToTable(f.Tide)
}

func areas() {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("There was an error while building the SurflineR client")

		return
	}

	as, err := c.Areas()
	if err != nil {
		fmt.Println("There was an error while fetching the areas")

		return
	}

	placeToTable(as)
}

func regions(aID string) {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("There was an error while building the SurflineR client")
	}

	rs, err := c.Regions(aID)
	if err != nil {
		fmt.Println("There was an error while fetching the regions")
	}

	placeToTable(rs)
}

func subRegions(aID string, rID string) {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("There was an error while building the SurflineR client")
	}

	rs, err := c.SubRegions(aID, rID)
	if err != nil {
		fmt.Println("There was an error while fetching the subregions")
	}

	placeToTable(rs)
}

func validateSubRegion(aID string, rID string, srID string) error {
	c, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Println("There was an error while building the SurflineR client")
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

// use merging on date https://github.com/olekukonko/tablewriter#example-6----identical-cells-merging
func tideToTable(t surflinef.Tide) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Time", "Description", "Height"})
	table.SetAutoMergeCells(true)

	filtered := filterPoints(t.DataPoints)

	tf := "2006-01-02 15:04:05"

	for i := range filtered {
		t := filtered[i]

		tt, err := time.Parse(tf, t.Localtime)
		if err != nil {
			fmt.Println("Error while parsing date for tides")

			return
		}

		td := fmt.Sprintf("%d/%d/%d", tt.Month(), tt.Day(), tt.Year())
		ttt := fmt.Sprintf("%02d:%02d", tt.Hour(), tt.Minute())
		h := strconv.FormatFloat(float64(t.Height), 'f', 2, 32)
		table.Append([]string{td, ttt, t.Type, h})
	}

	table.Render()
}

func filterPoints(ps []surflinef.DataPoint) []surflinef.DataPoint {
	vps := []surflinef.DataPoint{}
	for i := range ps {
		p := ps[i]

		if validPoint(p) {
			vps = append(vps, p)
		}
	}

	return vps
}

func validPoint(p surflinef.DataPoint) bool {
	return p.Type == "Low" || p.Type == "High"
}

func validDayAmount(d int) bool {
	if d < 1 {
		return false
	} else if d > 15 {
		return false
	} else {
		return true
	}
}
