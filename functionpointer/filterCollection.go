package functionpointer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func filter(iterable []int, predicate func(i int) bool) []int { // 1 variant
func filter(predicate func(int) bool, iterable []int) []int { // 2 variant
	iter := make([]int, 0)
	for _, v := range iterable {
		if predicate(v) {
			iter = append(iter, v)
		}
	}
	return iter
	// проверка на четность
	// добавление в слайс
	// return
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
}

func main() {

	src := readInput()
	// src := []int{1, 2, 3, 4, 5, 6}

	// res := filter(src, func(i int) bool { return i%2 == 0 }) // 1 variant
	res := filter(func(i int) bool { return i%2 == 0 }, src) // 2 variant

	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	// res := filter(...)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
