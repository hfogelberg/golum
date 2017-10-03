package golum

import (
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/kniren/gota/dataframe"
)

func CreateHistograms(df *dataframe.DataFrame, cols []string) error {
	if len(cols) > 0 {
		log.Println("Multiple columns")
		for _, col := range cols {
			dfSel := df.Select([]string{col})
			if err := drawHistogram(&dfSel); err != nil {
				return err
			}
		}
	} else {
		if err := drawHistogram(df); err != nil {
			return err
		}
	}
	return nil
}

func drawHistogram(df *dataframe.DataFrame) error {
	for _, colName := range df.Names() {
		log.Printf("Drawing histogram for %s\n", colName)
		v := make(plotter.Values, df.Nrow())
		for i, val := range df.Col(colName).Float() {
			v[i] = val
		}

		p, err := plot.New()
		if err != nil {
			log.Printf("Error making plot %s\n", err.Error())
			return err
		}

		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)
		h, err := plotter.NewHist(v, 16)
		if err != nil {
			log.Printf("Error creating histogram %s\n", err.Error())
			return err
		}

		h.Normalize(1)
		p.Add(h)

		name := fmt.Sprintf("%s_histogram.png", colName)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, name); err != nil {
			log.Printf("Error saving file %s\n", err.Error())
			return err
		}
	}

	return nil
}
