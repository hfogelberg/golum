package golum

import (
	"log"
	"os"
	"testing"

	"github.com/kniren/gota/dataframe"
)

func TestHistogram(t *testing.T) {
	f, err := os.Open("data/labeled_iris.csv")
	if err != nil {
		log.Printf("Error opening file %s\n", err.Error())
		t.Error(err.Error())
		return
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	if err := CreateHistograms(df); err != nil {
		t.Error(err.Error())
	}

	for _, name := range []string{"histogram_sepal_length.png", "histogram_sepal_width.png", "histogram_petal_length.png", "histogram_petal_width.png"} {
		os.Remove(name)
	}
}

func TestScatterplot(t *testing.T) {
	f, err := os.Open("data/labeled_iris.csv")
	if err != nil {
		log.Printf("Error opening file %s\n", err.Error())
		t.Error(err.Error())
		return
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	cols := []string{"sepal_length", "sepal_width", "petal_length", "petal_width"}
	if err := CreatScatterplots(df, cols); err != nil {
		t.Error(err.Error())
	}

	for _, name := range []string{"sepal_length_scatterplot.png", "sepal_width_scatterplot.png", "petal_length_scatterplot.png", "petal_width_scatterplot.png"} {
		os.Remove(name)
	}
}
