package golum

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/kniren/gota/dataframe"
)

func CreateScatterplots(dfAll *dataframe.DataFrame, cols []string, target string) error {
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
