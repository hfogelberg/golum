package golum

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateOneHistogram(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_length"}
	if err := CreateHistograms(file, cols); err != nil {
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
	if err := CreateHistograms(file, cols); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_histogram.png", col)
		os.Remove(name)
	}
}
func TestHistogramsSpeciesFirst(t *testing.T) {
	file := "data/iris_species_first.csv"
	cols := []string{"sepal_length", "sepal_width", "petal_length", "petal_width"}
	if err := CreateHistograms(file, cols); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_histogram.png", col)
		os.Remove(name)
	}
}
