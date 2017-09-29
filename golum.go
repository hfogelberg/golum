package golum

import (
	"fmt"
	"log"
	"math"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Create histogram for each column in the dataframe
func CreateHistograms(df dataframe.DataFrame) error {
	log.Println("Creating histograms")
	numElems := len(df.Names())
	for i := 0; i < len(df.Names()); i++ {
		plotVals := make(plotter.Values, df.Nrow())
		name := df.Names()[i]
		for j, fVal := range df.Col(name).Float() {
			if math.IsNaN(fVal) {
				// Does not contain numeric data, so skip
				i++
				if i >= numElems-1 {
					return nil
				}
				break
			}
			plotVals[j] = fVal
		}

		// Make plot
		p, err := plot.New()
		if err != nil {
			log.Printf("Error creating plot %s\n", err.Error())
			return err
		}

		// Set title
		p.Title.Text = fmt.Sprintf("Histogram of a %s", name)

		// Create histogram
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Printf("Error plotting %s %s\n", name, err.Error())
			return err
		}
		h.Normalize(1)

		// Add histogram to plot
		p.Add(h)
		fileName := fmt.Sprintf("histogram_%s.png", name)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, fileName); err != nil {
			log.Printf("Error saving plot %s\n", err.Error())
			return err
		}
		log.Printf("%s created", name)
	}

	log.Println("Done creating histograms")
	return nil
}

// Create scatterplot for each column in the CSV file
func CreatScatterplots(df dataframe.DataFrame, names []string) error {
	// Create histogram for each feature column in the dataset
	for _, colName := range names {
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
			p.Add(h)

			// Save plot to png file
			name := fmt.Sprintf("%s_scatterplot.png", colName)
			if err := p.Save(4*vg.Inch, 4*vg.Inch, name); err != nil {
				log.Printf("Error saving file %s\n", err.Error())
				return err
			}
		}
	}

	return nil
}
