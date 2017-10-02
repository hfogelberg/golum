package golum

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

type RegressionFormula struct {
	FormulaText string
	Constant    float64
	Coef        float64
}

func TrainLinearModel(file string, fieldsPerRecord int, input string, observed string) (RegressionFormula, error) {
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
	r.SetObserved(observed)
	r.SetVar(0, input)
	for i, record := range trainingData {
		// Skip header
		if i == 0 {
			continue
		}

		// Parse sales val
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Printf("Error parsing sales data %s\n", err.Error())
			return formula, err
		}

		// Parse TV val
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Error parsing sales data %s\n", err.Error())
			return formula, err
		}

		// Add data ponts to training model and do the training
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
		r.Run()
	}
	fmt.Printf("\nRegression formula: \n%v\n", r.Formula)

	formula.FormulaText = r.Formula
	formula.Constant = r.Coeff(0)
	formula.Coef = r.Coeff(1)
	return formula, nil
}
