package shared

import "github.com/mhelmetag/surflinef/v2"

// TaxonomySlice attaches the methods of Interface to []surflinef.Taxonomy, sorting in increasing order
type TaxonomySlice []surflinef.Taxonomy

func (x TaxonomySlice) Len() int           { return len(x) }
func (x TaxonomySlice) Less(i, j int) bool { return x[i].Name < x[j].Name }
func (x TaxonomySlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
