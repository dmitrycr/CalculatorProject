package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	// Удаляем пробелы
	expression = strings.ReplaceAll(expression, " ", "")

	// Проверка на корректность скобок
	if !isValidParentheses(expression) {
		return 0, fmt.Errorf("некорректное выражение: скобки")
	}

	// Преобразуем выражение в постфиксную нотацию
	postfixExpression := infixToPostfix(expression)

	// Вычисляем выражение в постфиксной нотации
	return evaluatePostfix(postfixExpression)
}

// Проверяет корректность скобок в выражении
func isValidParentheses(expression string) bool {
	stack := []rune{}
	for _, char := range expression {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// Преобразует инфиксное выражение в постфиксную нотацию
func infixToPostfix(expression string) string {
	stack := []rune{}
	postfix := ""
	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	for _, char := range expression {
		if isDigit(char) {
			postfix += string(char)
		} else if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				postfix += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else {
			for len(stack) > 0 && precedence[char] <= precedence[stack[len(stack)-1]] {
				postfix += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, char)
		}
	}

	for len(stack) > 0 {
		postfix += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

// Вычисляет выражение в постфиксной нотации
func evaluatePostfix(postfix string) (float64, error) {
	stack := []float64{}

	for _, char := range postfix {
		if isDigit(char) {
			num, err := strconv.ParseFloat(string(char), 64)
			if err != nil {
				return 0, fmt.Errorf("некорректное выражение: число")
			}
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("некорректное выражение: недостаточно операндов")
			}

			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := 0.0
			switch char {
			case '+':
				result = op1 + op2
			case '-':
				result = op1 - op2
			case '*':
				result = op1 * op2
			case '/':
				if op2 == 0 {
					return 0, fmt.Errorf("некорректное выражение: деление на ноль")
				}
				result = op1 / op2
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("некорректное выражение: ошибка вычисления")
	}

	return stack[0], nil
}

// Проверка на цифру
func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
