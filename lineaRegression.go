package golum

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
)

type RegressionFormula struct {
	FormulaText string
	Constant    float64
	Coef        float64
}

func TrainLinearModel(file string, fieldsPerRecord int, xCol string, yCol string, xPos int, yPos int) (RegressionFormula, error) {
	var formula RegressionFormula
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening training file %s\n", err.Error())
		return formula, err
	}
	defer f.Close()

	// Read all training data
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = fieldsPerRecord
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error reading training data %s\n", err.Error())
		return formula, err
	}

	var r regression.Regression
	r.SetObserved(xCol)
	r.SetVar(0, yCol)
	for i, record := range trainingData {
		// Skip header
		if i == 0 {
			continue
		}

		yVal, err := strconv.ParseFloat(record[yPos], 64)
		if err != nil {
			log.Printf("Error parsing y data %s\n", err.Error())
			return formula, err
		}

		xVal, err := strconv.ParseFloat(record[xPos], 64)
		if err != nil {
			log.Printf("Error parsing x data %s\n", err.Error())
			return formula, err
		}

		// Add data ponts to training model and do the training
		r.Train(regression.DataPoint(yVal, []float64{xVal}))
		r.Run()
	}

	formula.FormulaText = r.Formula
	formula.Constant = r.Coeff(0)
	formula.Coef = r.Coeff(1)

	return formula, nil
}

func VisualizeRegression(df *dataframe.DataFrame, xCol string, yCol string, formulaConst float64, formulaCoeff float64) error {
	fmt.Println("Visaulize regression")
	// Extract traget columns
	yVals := df.Col(yCol).Float()

	// Pts will hold the observed values for plotting
	pts := make(plotter.XYs, df.Nrow())

	// ptsPred will hold the predicted values for plotting
	ptsPred := make(plotter.XYs, df.Nrow())

	// Fill plots with data
	for i, fv := range df.Col(xCol).Float() {
		pts[i].X = fv
		pts[i].Y = yVals[i]
		ptsPred[i].X = fv
		ptsPred[i].Y = predict(fv, formulaConst, formulaCoeff)
	}

	// Create plot
	p, err := plot.New()
	if err != nil {
		log.Printf("Error creating plot %s\n", err.Error())
		return err
	}

	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"

	p.Add(plotter.NewGrid())

	// Add scatter plot points for observations
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Printf("Error adding scatter plot %s\n", err.Error())
		return err
	}
	s.GlyphStyle.Radius = vg.Points(3)

	// Add line plot for predictions
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Printf("Error adding line plot %s\n", err.Error())
		return err
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	// Save to png
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "linear_regression.png"); err != nil {
		log.Printf("Error saving plot %s", err.Error())
		return err
	}

	return nil
}

func predict(fv float64, constVal float64, coeff float64) float64 {
	pred := (fv * coeff) + constVal
	return pred
}
