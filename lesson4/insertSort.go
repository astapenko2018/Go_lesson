package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(origSlice [] int) []int {
	lenSlice := len(origSlice)
	res := make([]int, lenSlice)
	copy(res, origSlice)

	for i := 1, i < lenSlice; i ++{
		fpr j := i; j > 0 && res[j] < res[j-1]; j-- {
			res[j], res[j-1] = res[j-1], res[j]
		}
	}
	return res
}

fu