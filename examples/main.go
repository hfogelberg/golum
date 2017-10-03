package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hfogelberg/golum"
)

func main() {
	file := "../data/labeled_iris.csv"
	df, err := golum.GetDFFromCSV(file, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := golum.CreateHistograms(&df, nil); err != nil {
		log.Println(err.Error())
		return
	}

	cols := []string{"sepal_length", "sepal_width", "petal_length", "petal_width"}
	for _, col := range cols {
		name := fmt.Sprintf("%s_histogram.png", col)
		os.Remove(name)
	}

	// file := "../data/advertising.csv"
	// // cols := []string{"TV", "Sales"}
	// df, err := golum.GetDFFromCSV(file, nil)
	// if err != nil {
	// 	return
	// }
	// if err := golum.CreateHistograms(&df, nil); err != nil {
	// 	fmt.Printf("ERR! %s\n", err.Error())
	// 	return
	// }

	// fmt.Println("OK!")

}
