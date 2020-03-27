/*	Contains implementation of Bubble Sort.*/

/*
Author : Aryan Maurya
Date : 25 March 2020
*/

package main

func swap(input []int, index, nextIndex int) {
	temp := input[index]
	input[index] = input[nextIndex]
	input[nextIndex] = temp
}

// BubbleSort : Implementation of Bubble Sort
func BubbleSort(input []int) {
	length := len(input)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if input[j] > input[j+1] {
				swap(input, j, j+1)
			}
		}
	}
}
