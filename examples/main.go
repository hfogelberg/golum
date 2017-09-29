package main

import (
	"log"
	"os"

	"github.com/hfogelberg/golum"
	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("../data/labeled_iris.csv")
	if err != nil {
		log.Printf("Error opening file %s\n", err.Error())
		return
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	cols := []string{"sepal_length", "sepal_width", "petal_length", "petal_width"}
	if err := golum.CreatScatterplots(df, cols); err != nil {
		log.Println(err.Error())
	}
	log.Println("OK!")
}
