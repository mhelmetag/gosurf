package cmd

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/mhelmetag/gosurf/shared"
	"github.com/mhelmetag/gosurf/workarounds"

	"github.com/manifoldco/promptui"
	"github.com/mhelmetag/surflinef/v2"
)

const earth = "58f7ed51dadb30820bb38782"
const initialStep = 0
const subregionStep = 3
const spotStep = 5

// SearchInteractive opens interactive selects to navigate the taxonomy tree
func SearchInteractive() {
	bu, err := url.Parse(surflinef.TaxonomyBaseURL)
	if err != nil {
		fmt.Println("An unexpected error occured")

		return
	}

	c := surflinef.Client{BaseURL: bu}

	ts, err := getTaxonomy(c, earth)

	if err != nil {
		fmt.Println("An error occured while fetching the taxonomy tree from Surfline")

		return
	}

	promptOrBail(c, ts, initialStep)
}

func getTaxonomy(c surflinef.Client, id string) ([]surflinef.Taxonomy, error) {
	q := surflinef.TaxonomyQuery{
		ID:       id,
		MaxDepth: 0,
		Type:     "taxonomy",
	}

	t, err := c.GetTaxonomy(q)
	if err != nil {
		return []surflinef.Taxonomy{}, err
	}

	return t.Contains, nil
}

func promptOrBail(c surflinef.Client, ts []surflinef.Taxonomy, step int) error {
	if step >= spotStep {
		// Last step
		return nil
	} else if step == subregionStep {
		fts := []surflinef.Taxonomy{}
		for i := range ts {
			t := ts[i]

			if validSubregion(t) {
				fts = append(fts, t)
			}
		}

		id, err := deliverPrompt(fts)
		if err != nil {
			return err
		}

		nts, err := getTaxonomy(c, id)

		step++

		promptOrBail(c, nts, step)
	} else {
		id, err := deliverPrompt(ts)
		if err != nil {
			return err
		}

		nts, err := getTaxonomy(c, id)

		step++

		promptOrBail(c, nts, step)
	}

	return nil
}

func validSubregion(t surflinef.Taxonomy) bool {
	return t.Type == "subregion"
}

func deliverPrompt(ts []surflinef.Taxonomy) (string, error) {
	sort.Sort(shared.TaxonomySlice(ts))

	var names []string
	for i := range ts {
		t := ts[i]
		// Need to present different IDs for normal geonames (0, 1, 2), subregions (3) and spots (4)
		w := shared.TaxonomyWrapper{Taxonomy: t}
		n := w.TreeName()

		names = append(names, n)
	}

	prompt := promptui.Select{
		Label:  "Select Taxonomy",
		Items:  names,
		Stdout: &workarounds.BellSkipper{},
	}

	i, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	t := ts[i]
	id := t.ID

	return id, nil
}
