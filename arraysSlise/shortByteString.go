package arraysslise

import "fmt"

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

	// res := text // короткое решение
	// if len(text) > width {
	// 	res = text[:width] + "..."
	// }

	var res string

	if len(text) <= width {
		for i := 0; i < len(text); i++ {
			res += string(text[i])
		}
	} else {
		for i := 0; i < width; i++ {
			res += string(text[i])
		}
		res += "..."
	}

	fmt.Println(res)
}
