package main

import (
	"fmt"

	"example.com/learngo/filemanager"
	"example.com/learngo/prices"
)

func priceCalc() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Job done")
		}
	}

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
