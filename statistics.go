package golum

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/gonum/floats"
	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
	"github.com/sajari/regression"
	"gonum.org/v1/gonum/stat"
)

type MLStats struct {
	Mean       float64
	Mode       float64
	ModeCount  float64
	Median     float64
	Min        float64
	Max        float64
	Range      float64
	Variance   float64
	StDev      float64
	Quantile25 float64
	Quantile50 float64
	Quantile75 float64
}

func GetStatistics(df *dataframe.DataFrame, cols []string) ([]MLStats, error) {
	var statisticData []MLStats
	var err error

	for _, colName := range cols {
		var s MLStats
		col := df.Col(colName).Float()
		s.Mean = stat.Mean(col, nil)
		s.Mode, s.ModeCount = stat.Mode(col, nil)
		s.Median, err = stats.Median(col)
		if err != nil {
			log.Printf("Error calculating median %s\n", err.Error())
			return statisticData, err
		}
		s.Min = floats.Min(col)
		s.Max = floats.Max(col)
		s.Range = s.Max - s.Min
		s.Variance = stat.Variance(col, nil)
		s.StDev = stat.StdDev(col, nil)

		// Sort the values
		inds := make([]int, len(col))
		floats.Argsort(col, inds)

		s.Quantile25 =
			stat.Quantile(0.25, stat.Empirical, col, nil)
		s.Quantile50 =
			stat.Quantile(0.50, stat.Empirical, col, nil)
		s.Quantile75 =
			stat.Quantile(0.75, stat.Empirical, col, nil)

		statisticData = append(statisticData, s)
	}

	return statisticData, nil
}

func CalculateMAE(file string, fieldsPerRecord int, xCol int, yCol int) (float64, error) {
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

		// Parse X
		xVal, err := strconv.ParseFloat(record[xCol], 64)
		if err != nil {
			log.Printf("Error parsing X %s\n", err.Error())
			return 0, err
		}

		// Predict Y
		var r regression.Regression
		yPred, err := r.Predict([]float64{xVal})

		// Add to MAE
		mAE += math.Abs(yObs-yPred) / float64(len(testData))
	}
	fmt.Printf("MAE = %f0.2\n\n", mAE)

	return mAE, nil
}

func CalculateMSE(file string, fieldsPerRecord int, xCol int, yCol int) (float64, error) {
	var mSE float64
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

		// Parse X
		xVal, err := strconv.ParseFloat(record[xCol], 64)
		if err != nil {
			log.Printf("Error parsing X %s\n", err.Error())
			return 0, err
		}

		// Predict Y
		var r regression.Regression
		yPred, err := r.Predict([]float64{xVal})

		// Add to MAE
		mSE += math.Pow((yObs-yPred), 2) / float64(len(testData))
	}
	fmt.Printf("MSE = %f0.2\n\n", mSE)

	return mSE, nil
}
