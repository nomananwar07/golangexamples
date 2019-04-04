package golangexamples

func ConccatSlice(sliceToConcat []byte) string {
	var str string
	for i := 0; i < len(sliceToConcat); i++ {
		str+=string(sliceToConcat[i]) + string('-')
	}

	return str
}


