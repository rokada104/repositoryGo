package errore

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// не удаляйте, они нужны для проверки
var _ = errors.As
var _ = reflect.Append
var _ = runtime.Gosched

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	var err error
	// defer func() {
	// 	fmt.Print("-> ")
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(acc, trans)
	// }()
	acc, trans, err = parseInput()
	if err != nil {
		fmt.Print("-> ")
		fmt.Println(err)

	} else {
		fmt.Print("-> ")
		fmt.Println(acc, trans)
	}

}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	var errs error
	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)           // ,,
	trans, err1 := parseTransactions(transSrc) // ,,
	if err != nil {
		errs = err
	} else if err1 != nil {
		errs = err1
	}
	return acc, trans, errs
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

//errors
// 80/-10 10 -20 30
// -> expect overdraft >= 0

// x/10 10 -20 30

// parseAccount парсит счет из строки
// в формате balance/overdraft.
// func parseAccount(src string) account {
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0]) // ,,
	if err != nil {                        // ..
		return account{}, err
	}
	overdraft, err := strconv.Atoi(parts[1]) // ,,
	if err != nil {                          // ..
		return account{}, err
	}
	if overdraft < 0 {
		// panic("expect overdraft >= 0")
		return account{}, errors.New("expect overdraft >= 0")
	}
	if balance < -overdraft {
		// panic("balance cannot exceed overdraft")
		return account{}, errors.New("balance cannot exceed overdraft")
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
// func parseTransactions(src []string) []int {
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s) // ,,
		if err != nil {           // ..
			return trans, err
		}
		trans[idx] = t
	}
	return trans, nil
}
