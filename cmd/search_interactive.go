package cmd

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/mhelmetag/gosurf/shared"
	"github.com/mhelmetag/gosurf/workarounds"

	"github.com/manifoldco/promptui"
	"github.com/mhelmetag/surflinef/v2"
)

const EARTH = "58f7ed51dadb30820bb38782"

// SearchInteractive opens interactive selects to navigate the taxonomy tree
func SearchInteractive(t string) {
	switch t {
	case "subregion":
		searchSubregions()
	case "spot":
		searchSpots()
	default:
		fmt.Println("Incorrect search type. Must be one of: subregion or spot")
		return
	}
}

func searchSpots() {
	bu, err := url.Parse(surflinef.TaxonomyBaseURL)
	if err != nil {
		fmt.Println("An unexpected error occured")

		return
	}

	c := surflinef.Client{BaseURL: bu}

	ts, err := getTaxonomy(c, EARTH, 0)

	if err != nil {
		fmt.Println("An error occured while fetching the taxonomy tree from Surfline")

		return
	}

	step := 0

	promptOrBailSpots(c, ts, step)
}

func searchSubregions() {
	bu, err := url.Parse(surflinef.TaxonomyBaseURL)
	if err != nil {
		fmt.Println("An unexpected error occured")

		return
	}

	c := surflinef.Client{BaseURL: bu}

	ts, err := getTaxonomy(c, EARTH, 0)

	if err != nil {
		fmt.Println("An error occured while fetching the taxonomy tree from Surfline")

		return
	}

	promptOrBailSubregions(c, ts, false)
}

func getTaxonomy(c surflinef.Client, id string, maxDepth int) ([]surflinef.Taxonomy, error) {
	q := surflinef.TaxonomyQuery{
		ID:       id,
		MaxDepth: maxDepth,
		Type:     "taxonomy",
	}

	t, err := c.GetTaxonomy(q)
	if err != nil {
		return []surflinef.Taxonomy{}, err
	}

	return t.Contains, nil
}

func promptOrBailSpots(c surflinef.Client, ts []surflinef.Taxonomy, step int) error {
	// step == 0 Earth (depth 0)
	// step == 1 Continent (depth 0)
	// step == 2 Country (depth 0)
	// step == 3 Region (depth 1)
	// step == 4 Area & Spots (stop)
	// step == 5 Nothing

	if step == 5 {
		return nil
	}

	depth := 0

	if step == 3 {
		depth = 1
	}

	id, err := deliverPrompt(ts)
	if err != nil {
		return err
	}

	nts, err := getTaxonomy(c, id, depth)
	if err != nil {
		return err
	}

	fts := []surflinef.Taxonomy{}
	for i := range nts {
		t := nts[i]

		if step == 3 {
			if t.Type == "spot" {
				fts = append(fts, t)
			}
		} else {
			if t.Type == "geoname" && t.HasSpots {
				fts = append(fts, t)
			}
		}
	}

	step++

	promptOrBailSpots(c, fts, step)

	return nil
}

func promptOrBailSubregions(c surflinef.Client, ts []surflinef.Taxonomy, subregionsExist bool) error {
	// Keep searching for subregions
	// Stop once found

	if subregionsExist {
		return nil
	} else {
		srs := []surflinef.Taxonomy{}
		for i := range ts {
			t := ts[i]

			if t.Type == "subregion" {
				srs = append(srs, t)
			}
		}

		var id string
		var err error
		subregionsExist := len(srs) > 0
		if subregionsExist {
			id, err = deliverPrompt(srs)
		} else {
			id, err = deliverPrompt(ts)
		}
		if err != nil {
			return err
		}

		nts, err := getTaxonomy(c, id, 0)
		if err != nil {
			return err
		}

		promptOrBailSubregions(c, nts, subregionsExist)
	}

	return nil
}

func deliverPrompt(ts []surflinef.Taxonomy) (string, error) {
	sort.Sort(shared.TaxonomySlice(ts))

	var names []string
	for i := range ts {
		t := ts[i]
		n := treeName(t)

		names = append(names, n)
	}

	searcher := func(in string, i int) bool {
		t := ts[i]
		n := strings.Replace(strings.ToLower(t.Name), " ", "", -1)
		in = strings.Replace(strings.ToLower(in), " ", "", -1)

		return strings.Contains(n, in)
	}

	prompt := promptui.Select{
		Label:    "Select Taxonomy",
		Items:    names,
		Searcher: searcher,
		Stdout:   &workarounds.BellSkipper{},
	}

	i, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	t := ts[i]
	id := t.ID

	return id, nil
}

func treeName(t surflinef.Taxonomy) string {
	if t.Type == "subregion" {
		return fmt.Sprintf("%s (%s)", t.Name, t.Subregion)
	} else if t.Type == "spot" {
		return fmt.Sprintf("%s (%s)", t.Name, t.Spot)
	} else {
		return fmt.Sprintf("%s (%s)", t.Name, t.ID)
	}
}
