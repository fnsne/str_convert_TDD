package TDD_string_convertor

import "strings"

func converter(input string) string {
	if len(input) == 2 {
		first := input[0]
		second := input[1]
		return string(first) + "-" + strings.ToUpper(string(second)) + string(second)
	} else {
		return strings.ToUpper(input)
	}
}
