package iteration

func Repeat(character string, num int) string {
	// zero value for a string is ""
	var repeated string
	for i := 0; i < num; i++ {
		repeated += character
	}
	return repeated
}
