package golum

import (
	"fmt"
	"os"
	"testing"
)

func TestGetStatistics(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_length", "sepal_width", "petal_length", "petal_width"}
	df, err := GetDFFromCSV(file, cols)
	if err != nil {
		t.Error(err.Error())
	}

	stats, err := GetStatistics(&df, cols)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(stats)
}

func TestCalculateMAE(t *testing.T) {
	file := "data/salary_data.csv"
	_, testCsv, err := TrainTestSplit(file, 0.3)
	if err != nil {
		t.Error(err)
	}

	mae, err := CalculateMAE(testCsv, 2, 0, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(mae)

	os.Remove("test.csv")
	os.Remove("train.csv")
}

func TestCalculateMSE(t *testing.T) {
	file := "data/salary_data.csv"
	_, testCsv, err := TrainTestSplit(file, 0.3)
	if err != nil {
		t.Error(err)
	}

	mse, err := CalculateMSE(testCsv, 2, 0, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(mse)

	os.Remove("test.csv")
	os.Remove("train.csv")
}
