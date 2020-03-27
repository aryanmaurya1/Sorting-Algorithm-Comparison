/* Implementation of Generator functions.*/
/*
Author : Aryan Maurya
Date : 25 March 2020
*/

package main

import (
	"math/rand"
	"time"
)

func GeneratorAsc(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	return arr
}

func GeneratorDsc(n int) []int {
	var arr []int
	for i := n; i > 0; i-- {
		arr = append(arr, i)
	}
	return arr
}

func GeneratorRand(n int) []int {
	rand.Seed(time.Now().UnixNano())
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Int())
	}
	return arr
}
