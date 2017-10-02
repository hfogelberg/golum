package golum

import (
	"fmt"
	"os"
	"testing"
)

func TestOneScatterplot(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_length"}
	df, err := GetDFFromCSV(file, nil)
	if err != nil {
		t.Error(err.Error())
	}

	if err := CreateScatterplots(&df, cols, "sepal_width"); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_scatter.png", col)
		os.Remove(name)
	}
}

func TestMultipleScatterplots(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_length", "sepal_width", "petal_length", "petal_width"}
	df, err := GetDFFromCSV(file, nil)
	if err != nil {
		t.Error(err.Error())
	}

	if err := CreateScatterplots(&df, cols, "sepal_length"); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_scatter.png", col)
		os.Remove(name)
	}
}
