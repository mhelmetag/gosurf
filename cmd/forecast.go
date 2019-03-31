package cmd

import (
	"fmt"

	"github.com/mhelmetag/gosurf/helpers"

	"github.com/mhelmetag/surflinef"
)

func Forecast(aID string, rID string, srID string, d int) {
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
		Resources:		[]string{"analysis"},
		Days:				 d,
		Units:				"e",
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

	helpers.AnalysisToTable(f.Analysis)
}

func validDayAmount(d int) bool {
	// I'd like to support more days (it can be up to 17)
	// but I have to update SurflineF to hardcode the query param
	// "callback" which seems to allow more than 8 days
	// Or allow a token to be configured (not sure if the callback matters anymore)
	return d > 1 || d < 8
}
