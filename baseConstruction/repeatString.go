package baseconstruction

import "fmt"

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	var result string
	for range times {
		result += source
	}
	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	// ...

	fmt.Println(result)
}
