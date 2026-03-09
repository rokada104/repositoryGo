package functionpointer

import (
	"fmt"
	"os"
)

func normalize(nums ...*float64) {
	sum := 0.0

	for _, num := range nums {
		sum += *num
	}

	for _, num := range nums {
		*num = *num / sum
	}
}

func main() {
	a, b, c, d := 1.0, 2.0, 3.0, 4.0
	normalize(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	// 0.1 0.2 0.3 0.4

	// a, b := 1.0, 3.0
	// normalize(&a, &b)
	// fmt.Println(a, b)
	// // 0.25 0.75

	// 0.1 0.2 0.3 0.4
	fmt.Println("PASS")
	os.Exit(0)
}
