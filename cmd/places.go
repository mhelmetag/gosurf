package cmd

import (
  "fmt"

  "github.com/mhelmetag/gosurf/helpers"

  "github.com/mhelmetag/surfliner"
)

func Places(pType string, aID string, rID string) {
  var ps []surfliner.Place
  var err error

	switch pType {
	case "areas":
		ps, err = helpers.Areas()
	case "regions":
		ps, err = helpers.Regions(aID)
	case "subregions":
		ps, err = helpers.SubRegions(aID, rID)
	default:

	}

  if err != nil {
    fmt.Println("An error occured while searching spots...")
  }

  helpers.PlaceToTable(ps)
}
