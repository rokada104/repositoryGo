package structmethod

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func digits(s string) bool {
	for _, v := range s {
		if unicode.IsDigit(v) {
			return true
		}
	}
	return false
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func letters(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func minlen(length int) validator {
	return func(s string) bool { return utf8.RuneCountInString(s) >= length }

}

// and возвращает валидатор, который принимает строку и проверяет,
// что все funcs вернули true для этой строки
func and(funcs ...validator) validator {
	// return func(s string) bool { return digits(s) && letters(s) }
	return func(s string) bool {
		for _, v := range funcs {
			if v(s) == false {
				return false
			}
		}
		return true
	}
}

// or возвращает валидатор, который принимает строку и проверяет,
// что хотя бы одна из funcs вернула true для этой строки
func or(funcs ...validator) validator {
	return func(s string) bool {
		for _, v := range funcs {
			if v(s) == true {
				return true
			}
		}
		return false
	}
}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	return p.validator(p.value)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	var s string
	fmt.Scan(&s)
	// валидатор, который проверяет, что пароль содержит буквы и цифры,
	// либо его длина не менее 10 символов
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
