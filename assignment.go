package golangexamples
import	"github.com/ehteshamz/greetings"


func ConcatSlice(sliceToConcat []byte) string {
	var str string
	for i := 0; i < len(sliceToConcat); i++ {
		str+=string(sliceToConcat[i]) + string('-')
	}
	return str
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] = sliceToEncrypt[i] + byte(ceaserCount-48)
	}
}


func EZGreetings(name string) string{
	return greetings.PrintGreetings(name)
}
