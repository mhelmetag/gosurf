package cmd

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/mhelmetag/surflinef/v2"
	"github.com/olekukonko/tablewriter"
)

// Forecast gathers and prints the forecast table
func Forecast(sID string, d int) {
	bu, err := url.Parse("https://services.surfline.com/kbyg/regions/forecasts/conditions")
	if err != nil {
		fmt.Printf("An unexpected error occured")
		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.Query{
		Days:        d,
		SubregionID: sID,
	}

	qs, err := q.QueryString()
	if err != nil {
		fmt.Printf("An error occured while building the query to Surfline")
		return
	}

	cr, err := c.GetConditions(qs)
	if err != nil {
		fmt.Printf("An error occured while fetching the conditions from Surfline")
		return
	}

	cs := cr.Data.Conditions

	t := time.Now()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Time of Day", "Rating", "Range", "Forecast"})

	for i := range cs {
		ts := fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())

		cAM := cs[i].AM
		rangeAM := fmt.Sprintf("%.1f - %.1f", cAM.MinHeight, cAM.MaxHeight)
		table.Append([]string{ts, "AM", cAM.Rating, rangeAM, cAM.HumanRelation})

		cPM := cs[i].PM
		rangePM := fmt.Sprintf("%.1f - %.1f", cAM.MinHeight, cAM.MaxHeight)
		table.Append([]string{ts, "PM", cPM.Rating, rangePM, cPM.HumanRelation})

		t = t.AddDate(0, 0, 1)
	}

	table.Render()
}
