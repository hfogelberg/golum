package golum

import (
	"errors"
	"os"
	"testing"
)

func TestTrainTestSplit(t *testing.T) {
	file := "data/labeled_iris.csv"

	_, _, err := TrainTestSplit(file, 0.3)
	if err != nil {
		t.Error(err)
	}

	os.Remove("test.csv")
	os.Remove("train.csv")

}

func TestTrainTestSplitNoFile(t *testing.T) {
	file := "data/foo.csv"

	_, _, err := TrainTestSplit(file, 0.3)
	if err == nil {
		t.Error(errors.New("Should detect that file does not exist"))
	}
}

func TestTrainTestSplitWrongSize(t *testing.T) {
	file := "data/foo.csv"

	_, _, err := TrainTestSplit(file, 0)
	if err == nil {
		t.Error(errors.New("Wrong split size"))
	}
}
