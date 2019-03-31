package cmd

import (
  "fmt"

  "github.com/mhelmetag/gosurf/helpers"

  "github.com/mhelmetag/surflinef"
)

func Tide(aID string, rID string, srID string, d int) {
	c, err := surflinef.DefaultClient()
	if err != nil {
		fmt.Println("There was an error while building the SurflineF client")

		return
	}

	if !validDayAmount(d) {
		fmt.Println("The number of days to report can only be between 1 and 8")

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

	helpers.TideToTable(f.Tide)
}
