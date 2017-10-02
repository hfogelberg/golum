package golum

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
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

func TestLinearModel(file string, fieldsPerRecord int, xCol int, yCol int) (float64, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening test file %s\n", err.Error())
		return 0.0, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = fieldsPerRecord
	testData, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error reading test data %s\n", err.Error())
		return 0, err
	}

	// Loop over data and validate with Mean Absolute Error
	var mAE float64
	for i, record := range testData {
		// Skip header
		if i == 0 {
			continue
		}

		// Parse observed y
		yObs, err := strconv.ParseFloat(record[yCol], 64)
		if err != nil {
			log.Printf("Error parsing y %s\n", err.Error())
			return 0, err
		}

		// Parse X (TV value)
		tvVal, err := strconv.ParseFloat(record[xCol], 64)
		if err != nil {
			log.Printf("Error parsing X %s\n", err.Error())
			return 0, err
		}

		// Predict Y
		var r regression.Regression
		yPred, err := r.Predict([]float64{tvVal})

		// Add to MAE
		mAE += math.Abs(yObs-yPred) / float64(len(testData))
	}
	fmt.Printf("MAE = %f0.2\n\n", mAE)

	return mAE, nil
}
