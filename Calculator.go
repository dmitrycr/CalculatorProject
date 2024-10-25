package main

func Calc(expression string) (string, error) {
	var value string
	var d int = len(expression)
	for d != 0 {

		for i := 0; i < len(expression); i++ {
			value += string(expression[i])
			d--
		}
	}
	return value, nil
}
