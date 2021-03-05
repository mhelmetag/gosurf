package shared

import "github.com/mhelmetag/surflinef/v2"

// TaxonomySlice allows me to add sort related funcs to []surflinef.Taxonomy
type TaxonomySlice []surflinef.Taxonomy

func (x TaxonomySlice) Len() int           { return len(x) }
func (x TaxonomySlice) Less(i, j int) bool { return x[i].Name < x[j].Name }
func (x TaxonomySlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
