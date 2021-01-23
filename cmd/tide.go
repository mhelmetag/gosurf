package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/mhelmetag/surflinef/v2"
	"github.com/olekukonko/tablewriter"
)

// Tide gathers tide data for a spot and prints it
func Tide(sID string, d int) {
	bu, err := url.Parse(surflinef.TidesBaseURL)
	if err != nil {
		fmt.Println("An unexpected error occured")

		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.TidesQuery{
		Days:   d,
		SpotID: sID,
	}

	tr, err := c.GetTides(q)
	if err != nil {
		fmt.Println("An error occured while fetching the tides from Surfline")

		return
	}

	ts := tr.Data.Tides

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Time", "Description", "Height"})
	table.SetAutoMergeCells(true)

	ft := filterPoints(ts)

	for i := range ft {
		t := ft[i]

		tt := time.Unix(int64(t.Timestamp), 0)
		td := fmt.Sprintf("%d/%d/%d", tt.Month(), tt.Day(), tt.Year())
		ttt := fmt.Sprintf("%02d:%02d", tt.Hour(), tt.Minute())
		h := strconv.FormatFloat(float64(t.Height), 'f', 2, 32)

		table.Append([]string{td, ttt, t.Type, h})
	}

	table.Render()
}

func filterPoints(ts []surflinef.Tide) []surflinef.Tide {
	fts := []surflinef.Tide{}
	for i := range ts {
		t := ts[i]

		if validPoint(t) {
			fts = append(fts, t)
		}
	}

	return fts
}

func validPoint(t surflinef.Tide) bool {
	return t.Type == "LOW" || t.Type == "HIGH"
}
