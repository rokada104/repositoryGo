package interfaces

import (
	"fmt"
)

// element - интерфейс элемента последовательности
// (пустой, потому что элемент может быть любым).
type element interface{}

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	next() bool
	val() element
	// определите методы итератора
	// чтобы понять сигнатуры методов - посмотрите,
	// как они используются в функции iterate() ниже
	// `next()` переходит к следующему элементу и возвращает `true`. Либо возвращает `false`, если последовательность закончилась.
	// `val()` возвращает текущий элемент последовательности.
}

// iterate обходит последовательность
// и печатает каждый элемент
func iterate(it iterator) {
	for it.next() {
		curr := it.val()
		fmt.Println(curr)
	}
}

// в этом задании функция main() определена "за кадром",
// не добавляйте ее

//////////////////полное решение
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	next() bool
	val() element
	// чтобы понять сигнатуры методов - посмотрите,
	// как они используются в функции max() ниже
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	srs   []int
	index int
}

func (i *intIterator) next() bool {
	return i.index < len(i.srs)
}

func (i *intIterator) val() element {
	value := i.srs[i.index]
	i.index++
	return value

}

// методы intIterator, которые реализуют интерфейс iterator

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
	return &intIterator{
		srs:   src,
		index: -1,
	}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
	nums := readInput()
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
