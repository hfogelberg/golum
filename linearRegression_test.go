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

	formula, err := TrainLinearModel(trainCsv, 2, "Salary", "YearsExperience")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(formula)

	os.Remove("test.csv")
	os.Remove("train.csv")
}

func TestLinearModelTest(t *testing.T) {
	file := "data/salary_data.csv"
	_, testCsv, err := TrainTestSplit(file, 0.3)
	if err != nil {
		t.Error(err)
	}

	formula, err := TestLinearModel(testCsv, 2, 0, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(formula)

	os.Remove("test.csv")
	os.Remove("train.csv")
}