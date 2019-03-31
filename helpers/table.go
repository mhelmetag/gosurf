package helpers

import(
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mhelmetag/surflinef"
	"github.com/mhelmetag/surfliner"
	"github.com/olekukonko/tablewriter"
)

func PlaceToTable(ps []surfliner.Place) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})

	for i := range ps {
		p := ps[i]
		table.Append([]string{p.ID, p.Name})
	}

	table.Render()
}

func AnalysisToTable(a surflinef.Analysis) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Condition", "Report"})
	rs := analysisReports(a)
	t := time.Now()

	for i := range a.GeneralCondition {
		ts := fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())
		table.Append([]string{ts, a.GeneralCondition[i], rs[i]})
		t = t.AddDate(0, 0, 1)
	}

	table.Render()
}

func analysisReports(a surflinef.Analysis) []string {
	var rs []string

	for i := range a.GeneralCondition {
		r := fmt.Sprintf("%s-%sft. - %s", a.SurfMin[i], a.SurfMax[i], a.SurfText[i])

		rs = append(rs, r)
	}

	return rs
}

func TideToTable(t surflinef.Tide) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Time", "Description", "Height"})
	table.SetAutoMergeCells(true)

	filtered := filterPoints(t.DataPoints)

	tf := "2006-01-02 15:04:05"

	for i := range filtered {
		t := filtered[i]

		tdt, err := time.Parse(tf, t.Localtime)
		if err != nil {
			fmt.Println("Error while parsing date for tides")

			return
		}

		td := fmt.Sprintf("%d/%d/%d", tdt.Month(), tdt.Day(), tdt.Year())
		ttt := fmt.Sprintf("%02d:%02d", tdt.Hour(), tdt.Minute())
		h := strconv.FormatFloat(float64(t.Height), 'f', 2, 32)
		table.Append([]string{td, ttt, t.Type, h})
	}

	table.Render()
}

func filterPoints(ps []surflinef.DataPoint) []surflinef.DataPoint {
	vps := []surflinef.DataPoint{}
	for i := range ps {
		p := ps[i]

		if validPoint(p) {
			vps = append(vps, p)
		}
	}

	return vps
}

func validPoint(p surflinef.DataPoint) bool {
	return p.Type == "Low" || p.Type == "High"
}
