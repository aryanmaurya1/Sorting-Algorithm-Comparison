package main

import (
	"fmt"
	"time"
)

const maxObservation = 3

func main() {
	t := time.Now().UnixNano()

	var nValues = []int{10, 100, 1000, 10000, 100000}

	for _, value := range nValues {
		fmt.Println("----------------------------------------------------------------")
		for i := 0; i < 4; i++ {

			arr1 := GeneratorAsc(value)
			arr2 := GeneratorDsc(value)
			arr3 := GeneratorRand(value)

			t1 := time.Now().UnixNano()

			InsertionSort(arr1)
			BubbleSort(arr1)
			QuickSort(arr1, 0, len(arr1)-1)
			MergeSort(arr1)

			t2 := time.Now().UnixNano()

			InsertionSort(arr2)
			BubbleSort(arr2)
			QuickSort(arr2, 0, len(arr2)-1)
			MergeSort(arr2)

			t3 := time.Now().UnixNano()

			InsertionSort(arr3)
			BubbleSort(arr3)
			QuickSort(arr3, 0, len(arr3)-1)
			MergeSort(arr3)

			t4 := time.Now().UnixNano()

			fmt.Println()
			fmt.Println("Observation No. ", i+1)
			fmt.Println("Value of N is ", value)
			fmt.Println("Time taken to sort asc : ", t2-t1, "nanosec")
			fmt.Println("Time taken to sort dsc : ", t3-t2, "nanosec")
			fmt.Println("Time taken to sort rand : ", t4-t3, "nanosec")
			fmt.Println()

		}
		fmt.Println("----------------------------------------------------------------")
	}

	fmt.Println(time.Now().UnixNano() - t)
}
