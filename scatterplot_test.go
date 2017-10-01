package golum

import (
	"fmt"
	"os"
	"testing"
)

func TestOneScatterplot(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_length"}
	if err := CreateScatterplots(file, cols, "sepal_width"); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_scatter.png", col)
		os.Remove(name)
	}
}

func TestMultipleScatterplots(t *testing.T) {
	file := "data/labeled_iris.csv"
	cols := []string{"sepal_width", "petal_length", "petal_width"}
	if err := CreateScatterplots(file, cols, "sepal_length"); err != nil {
		t.Error(err.Error())
	}

	for _, col := range cols {
		name := fmt.Sprintf("%s_scatter.png", col)
		os.Remove(name)
	}
}

func TestPairPlot(t *testing.T) {
	cols := []string{"sepal_width", "petal_length", "petal_width"}
	if err := CreatePairplots("data/labeled_iris.csv", cols, cols); err != nil {
		t.Error(err.Error())
	}

	for _, f := range []string{"petal_length_scatter.png", "petal_width_scatter.png", "sepal_width_scatter.png"} {
		os.Remove(f)
	}

}
