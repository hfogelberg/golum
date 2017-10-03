# Golum

Helper functions for machine learning in Golang. Nothing is unique - it uses standard machine learning packages under the hood, but make them easier to work with and reduces the amount of boiler plate code that need to be written

Use with care! Very much work in production. The apis will most likely change

## GetDFFromCSV(string, []string) (dataframe.DataFrame, error)
Input the path to a CSV file and and a slice of colums to be parsed. If nil is passed to the function all columns are read.
Returns a dataframe (github.com/kniren/gota/dataframe)
````
  df, err := golum.GETDFFromCSV("iris.csv", {"sepal_length", "sepal_width"})
````

## func CreateHistograms(*dataframe.DataFrame) error
Create histograms for all columns in the dataframe

