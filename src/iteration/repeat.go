package iteration

func Repeat(character string, repeatCount int) string {
	// := is shorthand for declaring and initializing variables
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
