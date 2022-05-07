package main

import (
	"fmt"

	xgb "github.com/Elvenson/xgboost-go"
	"github.com/Elvenson/xgboost-go/activation"
	"github.com/Elvenson/xgboost-go/mat"
)

func main() {
	ensemble, err := xgb.LoadXGBoostFromJSON("your model path",
		"", 1, 4, &activation.Logistic{})
	if err != nil {
		panic(err)
	}

	input, err := mat.ReadLibsvmFileToSparseMatrix("your libsvm input path")
	if err != nil {
		panic(err)
	}
	predictions, err := ensemble.PredictProba(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", predictions)
}