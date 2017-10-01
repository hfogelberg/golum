package golum

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func GetDFFromCSV(file string, cols []string) (dataframe.DataFrame, error) {
	var df dataframe.DataFrame

	fData, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening CSV file %s\n", err.Error())
		return df, err
	}
	defer fData.Close()

	df = dataframe.ReadCSV(fData)

	if len(cols) > 0 {
		df = df.Select(cols)
	}

	return df, nil
}
