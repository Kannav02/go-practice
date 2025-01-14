package slices

func Sum(numbers [5]int) (sum int) {
	
	// gives the index and the value from the array 
	for _, number := range numbers {

		sum += number
	}

	return

}
