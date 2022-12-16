package utils

// Panic upon error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
