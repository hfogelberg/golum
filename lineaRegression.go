package golum

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func TrainLinearModel(file string, fieldsPerRecord int, input string, observed string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening training file %s\n", err.Error())
		return "", err
	}
	defer f.Close()

	// Read all training data
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = fieldsPerRecord
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error reading training data %s\n", err.Error())
		return "", err
	}

	var r regression.Regression
	r.SetObserved(observed)
	r.SetVar(0, input)
	for i, record := range trainingData {
		// Skip header
		if i == 0 {
			continue
		}

		// Parse sales val
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Printf("Error parsing sales data %s\n", err.Error())
			return "", err
		}

		// Parse TV val
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Error parsing sales data %s\n", err.Error())
			return "", err
		}

		// Add data ponts to training model and do the training
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
		r.Run()
	}
	fmt.Printf("\nRegression formula: \n%v\n", r.Formula)
	return r.Formula, nil
}
