package golum

import (
	"fmt"
	"os"
	"testing"
)

func TrainLinearModelTest(t *testing.T) {
	file := "data/salary_data.csv"

	trainCsv, _, err := TrainTestSplit(file, 0.3)
	if err != nil {
		t.Error(err)
	}

	f, err := TrainLinearModel(trainCsv, 2, "Salary", "YearsExperience")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Text %s, cons: %0.2f, coeff: %0.2f\n", f.FormulaText, f.Constant, f.Coef)

	os.Remove("test.csv")
	os.Remove("train.csv")
}
