package cmd

import (
	"fmt"
	"net/url"
	"os"
	"sort"

	"github.com/mhelmetag/gosurf/shared"

	"github.com/mhelmetag/surflinef/v2"
	"github.com/olekukonko/tablewriter"
)

// Search gathers the taxonomy tree for an ID and prints it
func Search(tID string, md int) {
	bu, err := url.Parse(surflinef.TaxonomyBaseURL)
	if err != nil {
		fmt.Println("An unexpected error occured")

		return
	}

	c := surflinef.Client{BaseURL: bu}

	q := surflinef.TaxonomyQuery{
		ID:       tID,
		MaxDepth: md,
		Type:     "taxonomy",
	}

	t, err := c.GetTaxonomy(q)
	if err != nil {
		fmt.Println("An error occured while fetching the taxonomy tree from Surfline")

		return
	}

	ts := t.Contains

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Type", "TypeID", "Name"})

	sort.Sort(shared.TaxonomySlice(ts))

	for i := range ts {
		t := ts[i]

		var tID string
		switch t.Type {
		case "spot":
			tID = t.Spot
		case "subregion":
			tID = t.Subregion
		default:
			tID = "N/A"
		}

		table.Append([]string{t.ID, t.Type, tID, t.Name})
	}

	table.Render()
}
