package golum

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateOneHistogram(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_length"}
	df, err := GetDFFromCSV(file, cols)
	if err != nil {
		t.Error(err.Error())
	}
	if err := CreateHistograms(&df); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_histogram.png", col)
		os.Remove(name)
	}
}

func TestMultipleHistograms(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_length", "sepal_width", "petal_length", "petal_width"}
	df, err := GetDFFromCSV(file, cols)
	if err != nil {
		t.Error(err.Error())
	}
	if err := CreateHistograms(&df); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_histogram.png", col)
		os.Remove(name)
	}
}
