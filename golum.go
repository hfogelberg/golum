package golum

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/kniren/gota/dataframe"
)

func CreateHistograms(file string, cols []string) error {
	iris, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening CSV file %s\n", err.Error())
		return err
	}
	defer iris.Close()

	df := dataframe.ReadCSV(iris)
	sel := df.Select(cols)

	// fmt.Println(sel)
	for _, colName := range sel.Names() {
		if colName != "species" {
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

			// Create histogram of drawn values from the standard normal
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Printf("Error creating histogram %s\n", err.Error())
				return err
			}

			// Normalize the diagram.
			// This makes it possible to compare different distributions side by side
			h.Normalize(1)

			// Add histogram to plot
			p.Add(h)

			// Save plot to png file
			name := fmt.Sprintf("%s_histogram.png", colName)
			if err := p.Save(4*vg.Inch, 4*vg.Inch, name); err != nil {
				log.Printf("Error saving file %s\n", err.Error())
				return err
			}
		}
	}
	return nil
}
