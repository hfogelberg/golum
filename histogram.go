package golum

import (
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/kniren/gota/dataframe"
)

func CreateHistograms(df *dataframe.DataFrame) error {
	for _, colName := range df.Names() {
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
