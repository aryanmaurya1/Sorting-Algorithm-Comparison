/*	Contains implementation of Quick Sort.*/

/*
Author : Aryan Maurya
Date : 25 March 2020
*/
package main

func pivot(arr []int, si, ei int) int {
	var p = ei // selecting as pivot
	var i = si - 1
	var j int
	for j = si; j <= p; j++ {
		if arr[j] <= arr[p] {
			i++
			swap(arr, i, j)
		}
	}
	return i // correct pivot
}

// QuickSort : Implementation of Quick Sort
func QuickSort(arr []int, si, ei int) {
	if si >= ei {
		return
	}
	var p = pivot(arr, si, ei)
	QuickSort(arr, p+1, ei)
	QuickSort(arr, si, p-1)
}
