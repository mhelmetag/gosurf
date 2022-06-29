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
	srName, err := getSubregionName(srID)
	if err != nil {
		fmt.Println("An error occured while fetching the subregion from Surfline")

		return
	}

	if srName == "" {
		fmt.Printf("The subregion with id %s doesn't exist\n", srID)

		return
	}

	fmt.Printf("Fetching %d day(s) of forecasts for %s...\n", d, srName)

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
		fmt.Println("An error occured while fetching the forecasts from Surfline")

		return
	}

	cs := cr.Data.Conditions

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Time of Day", "Rating", "Range", "Forecast"})
	table.SetAutoMergeCellsByColumnIndex([]int{0})

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
		return "Unknown"
	}
}

func getSubregionName(srID string) (string, error) {
	bu, err := url.Parse(surflinef.TaxonomyBaseURL)
	if err != nil {
		return "", err
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.TaxonomyQuery{
		ID:       srID,
		MaxDepth: 0,
		Type:     "subregion",
	}

	t, err := c.GetTaxonomy(q)
	if err != nil {
		return "", err
	}

	return t.Name, nil
}
