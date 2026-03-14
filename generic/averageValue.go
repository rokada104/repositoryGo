package generic

import (
	"fmt"
)

// начало решения

// Avg - накопительное среднее значение.
type Avg[T int | float64] struct { // чтобы хранить значение для прменения цепочного вызова
	slice []T
}

// Add пересчитывает среднее значение с учетом val.
func (a *Avg[T]) Add(val T) *Avg[T] {
	// s := Avg[T]{}
	a.slice = append(a.slice, val)
	return a
}

// Val возвращает текущее среднее значение.
func (a *Avg[T]) Val() T {
	if len(a.slice) == 0 {
		return 0
	}
	var num T
	for _, v := range a.slice {
		num += v
	}
	return num / T(len(a.slice))
}

// конец решения

func main() {
	intAvg := Avg[int]{}
	intAvg.Add(4).Add(3).Add(2)
	fmt.Println(intAvg.slice, len(intAvg.slice))
	fmt.Println(intAvg.Val()) // 3

	floatAvg := Avg[float64]{}
	floatAvg.Add(4.0).Add(3.0)
	floatAvg.Add(2.0).Add(1.0)
	fmt.Println(floatAvg.slice, len(floatAvg.slice))
	fmt.Println(floatAvg.Val()) // 2.5
}

import "fmt"


// func Reverse[T any](s []T) {
// 	for i := 0; i < len(s)/2; i++ {
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 	}
// }

// type IntSlice []int

// func (s IntSlice) Reverse() {
// 	for i := 0; i < len(s)/2; i++ {
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 	}
// }

// type StringSlice []string

// func (s StringSlice) Reverse() {
// 	for i := 0; i < len(s)/2; i++ {
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 	}
// }

type Slice[T any] []T

func (s Slice[T]) Reverse() {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

type Person struct {
	Name string
}

type Pair[T any] struct {
	first  T
	second T
}

func (p *Pair[T]) Swap() {
	p.first, p.second = p.second, p.first
}

func main() {
	intPair := Pair[int]{5, 3}
	intPair.Swap()
	fmt.Println(intPair)

	strPair := Pair[string]{"a", "b"}
	strPair.Swap()
	fmt.Println(strPair)

	personSlice := Slice[Person]{
		{Name: "Anton"}, {Name: "Aleksandr"}, {Name: "Ekaterina"},
	}

	personSlice.Reverse()
	fmt.Println(personSlice)

	intSlice := Slice[int]{1, 2, 3}
	intSlice.Reverse()
	fmt.Println(intSlice)

	strSlice := Slice[string]{"a", "b", "c"}
	strSlice.Reverse()
	fmt.Println(strSlice)

}
