package golum

import (
	"errors"
	"testing"
)

func TestGetDFFromCSV(t *testing.T) {
	cols := []string{}
	_, err := GetDFFromCSV("data/labeled_iris.csv", cols)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateDFFiltered(t *testing.T) {
	cols := []string{"sepal_length", "sepal_width"}
	df, err := GetDFFromCSV("data/labeled_iris.csv", cols)
	if err != nil {
		t.Error(err.Error())
	}

	if df.Ncol() != 2 {
		t.Error(errors.New("Wrong number of columns returned"))
	}
}

// func TestUnlabeledCSVShouldFail(t *testing.T) {
// 	cols := []string{}
// 	_, err := GetDFFromCSV("data/iris.csv", cols)
// 	if err == nil {
// 		t.Error(err.Error())
// 	}
// }
