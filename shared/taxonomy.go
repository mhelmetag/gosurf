package shared

import (
	"fmt"

	"github.com/mhelmetag/surflinef/v2"
)

type TaxonomyWrapper struct {
	Taxonomy surflinef.Taxonomy
}

func (t *TaxonomyWrapper) TreeName() string {
	switch t.Taxonomy.Type {
	case "spot":
		return fmt.Sprintf("%s (%s)", t.Taxonomy.Name, t.Taxonomy.Spot)
	case "subregion":
		return fmt.Sprintf("%s (%s)", t.Taxonomy.Name, t.Taxonomy.Subregion)
	default:
		return fmt.Sprintf("%s (%s)", t.Taxonomy.Name, t.Taxonomy.ID)
	}
}
