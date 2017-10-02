package golum

import (
	"log"

	"github.com/gonum/floats"
	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
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
