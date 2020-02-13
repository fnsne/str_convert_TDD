package TDD_string_convertor

import "strings"

func converter(input string) string {
	if len(input) == 2 {
		first := input[0]
		second := input[1]
		firstStatement := string(first)
		secondStatement := strings.ToUpper(string(second)) + string(second)
		generatedStrings := []string{firstStatement, secondStatement}
		return strings.Join(generatedStrings, "-")
	} else if len(input) == 3 {
		first := input[0]
		second := input[1]
		third := input[2]
		return string(first) + "-" +
			strings.ToUpper(string(second)) + string(second) + "-" +
			strings.ToUpper(string(third)) + string(third) + string(third)
	} else {
		return strings.ToUpper(input)
	}
}
