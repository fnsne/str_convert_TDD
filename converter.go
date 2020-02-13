package TDD_string_convertor

import "strings"

func converter(input string) string {
	if len(input) == 2 {
		var generatedStrings []string
		for index, c := range input {
			var genString string
			if index == 0 {
				genString = string(c)
			}
			if index == 1 {
				genString = string(c) + string(c)
			}
			generatedStrings = append(generatedStrings, genString)
		}
		return strings.Join(generatedStrings, "-")
	} else if len(input) == 3 {
		first := input[0]
		second := input[1]
		third := input[2]
		return string(first) + "-" +
			string(second) + string(second) + "-" +
			string(third) + string(third) + string(third)
	} else {
		return strings.ToUpper(input)
	}
}
