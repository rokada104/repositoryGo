package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Eyjafjallajokull 6
	// hello 6
	// Eyjafj...
	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

	// var res string
	// res := []rune(text)
	// fmt.Println(res)
	res := []string{}
	res1 := []string{}
	var res2 string
	var res3 string

	if len(text) <= width {
		for i := 0; i < len(text); i++ {
			res3 += string(text[i])
		}
	} else {
		for i := 0; i < width; i++ {
			res3 += string(text[i])
		}
		res3 += "..."
	}

	for _, v := range text {
		res = append(res, string(v))
	}

	if len(res) <= width {
		res1 = append(res1, res...)
	} else {
		for i := 0; i < width; i++ {
			res1 = append(res1, string(text[i]))
			res2 += string(text[i])
		}

		res1 = append(res1, "...")
		res2 += "..."
	}

	// if len(res) < width {
	// 	res = res[0]
	// }
	// for i := 0; i < width; i++ {

	// }
	// if len(res) < width {
	// 	res = append(res)
	// } else if len(text) >= width {
	// 	for i := 0; i < width; i++ {
	// 		res += string(text[i])
	// 	}
	// 	res += "..."
	// }

	// res1 := []byte(text)
	// res := text[width]

	// fmt.Println(string(res1))
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
