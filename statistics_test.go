package golum

import (
	"fmt"
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
