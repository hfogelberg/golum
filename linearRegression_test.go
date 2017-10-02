package golum

import (
	"fmt"
	"os"
	"testing"
)

func TestTrainLinearModel(t *testing.T) {
	file := "../data/advertising.csv"
	trainCsv, _, err := TrainTestSplit(file, 0.3)
	if err != nil {
		t.Error(err.Error())
	}

	f, err := TrainLinearModel(trainCsv, 4, "TV", "Sale", 0, 3)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("Formula: %s\n", f.FormulaText)
	os.Remove("test.csv")
	os.Remove("train.csv")
}

func TestVisualizeRegression(t *testing.T) {
	file := "data/advertising.csv"
	cols := []string{"TV", "Sales"}
	df, err := GetDFFromCSV(file, cols)
	if err != nil {
		t.Error(err.Error())
	}
	if err := VisualizeRegression(&df, "TV", "Sales", 7.98, 0.06); err != nil {
		t.Error(err.Error)
	}
	os.Remove("linear_regression.png")
}
