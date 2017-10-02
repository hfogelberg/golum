package golum

import (
	"bufio"
	"errors"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func TrainTestSplit(filename string, test float64) (string, string, error) {
	if test > 1 || test <= 0 {
		err := errors.New("Invalid value for test")
		return "", "", err
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file %s\n", err.Error())
		return "", "", err
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	fRows := float64(df.Nrow())
	trainNum := int(((1.0 - test) * fRows))
	testNum := int(test * fRows)
	if trainNum+testNum < df.Nrow() {
		trainNum++
	}

	// Create subset of indices
	trainIdx := make([]int, trainNum)
	testIdx := make([]int, testNum)

	// Enumerate training and testing indices
	for i := 0; i < trainNum; i++ {
		trainIdx[i] = i
	}

	for i := 0; i < testNum; i++ {
		testIdx[i] = i
	}

	// Create subset dataframes
	trainDF := df.Subset(testIdx)
	testDF := df.Subset(trainIdx)

	// Create map that will be used to write data files
	setMap := map[int]dataframe.DataFrame{
		0: trainDF,
		1: testDF,
	}

	// Create files
	for idx, setName := range []string{"train.csv", "test.csv"} {
		f, err := os.Create(setName)
		if err != nil {
			log.Printf("Error creating file %s\n", err.Error())
			return "", "", err
		}
		w := bufio.NewWriter(f)
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Printf("Error writing to file %s\n", err.Error())
			return "", "", err
		}

	}

	return "training.csv", "test.csv", nil
}
