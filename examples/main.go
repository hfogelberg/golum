package main

import (
	"fmt"
	"log"

	"github.com/hfogelberg/golum"
)

func main() {
	file := "../data/advertising.csv"
	trainCsv, _, err := golum.TrainTestSplit(file, 0.3)
	if err != nil {
		log.Println(err.Error())
	}

	f, err := golum.TrainLinearModel(trainCsv, 4, "TV", "Sales", 0, 3)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Formula: %s\n", f.FormulaText)
	fmt.Printf("Const: %0.2f, Coeff: %0.2f\n", f.Constant, f.Coef)

	cols := []string{"TV", "Sales"}
	df, err := golum.GetDFFromCSV(file, cols)
	if err := golum.VisualizeRegression(&df, "TV", "Sales", f.Constant, f.Coef); err != nil {
		log.Println(err.Error())
	}
}
