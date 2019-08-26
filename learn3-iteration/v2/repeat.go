package iteration

// Repeat01 func
func Repeat01(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}

// Repeat02 func
func Repeat02(character string) string {
	var repeated string
	i := 0
	for i < 5 {
		repeated = repeated + character
		i++
	}
	return repeated
}
