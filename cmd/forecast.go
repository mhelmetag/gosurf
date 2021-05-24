package cmd

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/mhelmetag/surflinef/v2"
	"github.com/olekukonko/tablewriter"
)

// Forecast gathers forecast data for a subregion and prints it
func Forecast(srID string, d int) {
	bu, err := url.Parse(surflinef.ConditionsBaseURL)
	if err != nil {
		fmt.Println("An unexpected error occured")

		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.ConditionsQuery{
		Days:        d,
		SubregionID: srID,
	}

	cr, err := c.GetConditions(q)
	if err != nil {
		fmt.Println("An error occured while fetching the conditions from Surfline")

		return
	}

	cs := cr.Data.Conditions

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Time of Day", "Rating", "Range", "Forecast"})

	for i := range cs {
		t := time.Unix(int64(cs[i].Timestamp), 0)
		ts := fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())

		cAM := cs[i].AM
		rangeAM := fmt.Sprintf("%.1f-%.1fft", cAM.MinHeight, cAM.MaxHeight)
		table.Append([]string{ts, "AM", convertRating(cAM.Rating), rangeAM, cAM.HumanRelation})

		cPM := cs[i].PM
		rangePM := fmt.Sprintf("%.1f-%.1fft", cAM.MinHeight, cAM.MaxHeight)
		table.Append([]string{ts, "PM", convertRating(cPM.Rating), rangePM, cPM.HumanRelation})
	}

	table.Render()
}

func convertRating(rating string) string {
	switch rating {
	case "FLAT":
		return "Flat"
	case "VERY_POOR":
		return "Very Poor"
	case "POOR":
		return "Poor"
	case "POOR_TO_FAIR":
		return "Poor to Fair"
	case "FAIR":
		return "Fair"
	case "FAIR_TO_GOOD":
		return "Fair to Good"
	case "GOOD":
		return "Good"
	case "VERY_GOOD":
		return "Very Good"
	case "GOOD_TO_EPIC":
		return "Good to Epic"
	case "EPIC":
		return "Epic"
	default:
		return "Unkown"
	}
}
