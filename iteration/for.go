package iteration

func Repeat(character string, stopPoint int) (repeatedString string) {

	if stopPoint == 0 {
		stopPoint = 5
	}
	for i := 0; i < stopPoint; i++ {
		repeatedString += character
	}
	return
}
