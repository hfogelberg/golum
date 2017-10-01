package golum

import (
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/kniren/gota/dataframe"
)

func CreateScatterplots(file string, cols []string, target string) error {
	iris, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening CSV file %s\n", err.Error())
		return err
	}
	defer iris.Close()

	dfAll := dataframe.ReadCSV(iris)
	df := dfAll.Select(cols)
	// Extract the target column
	yVals := dfAll.Col(target).Float()
	for _, name := range df.Names() {
		// pts holds data to be plotted
		pts := make(plotter.XYs, df.Nrow())
		for i, fVal := range df.Col(name).Float() {
			pts[i].X = fVal
			pts[i].Y = yVals[i]
		}

		// Create plot
		p, err := plot.New()
		if err != nil {
			log.Printf("Error plotting %s %s\n", name, err.Error())
			return err
		}

		p.X.Label.Text = name
		p.Y.Label.Text = target
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Printf("Error creating new scatter plot for %s %s\n", name, err.Error())
			return err
		}

		s.GlyphStyle.Radius = vg.Points(3)
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, name+"_scatter.png"); err != nil {
			log.Printf("Error saving scatter plot %s\n", err.Error())
			return err
		}
	}

	return nil
}

func CreatePairplots(file string, colsA []string, colsB []string) error {
	for _, colA := range colsA {
		var cols []string
		for _, colB := range colsB {
			if colA != colB {
				cols = append(cols, colB)
			}
		}
		if err := CreateScatterplots(file, cols, colA); err != nil {
			return err
		}
	}

	return nil
}
