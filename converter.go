package TDD_string_convertor

import "strings"

func converter(input string) string {
	var generatedStrings []string
	for index, c := range input {
		genString := duplicateByIndex(index, c)
		genString = firstCharToUpper(genString)
		generatedStrings = append(generatedStrings, genString)
	}
	return strings.Join(generatedStrings, "-")
}

func firstCharToUpper(genString string) string {
	genString = strings.ToUpper(string(genString[0])) + genString[1:]
	return genString
}

func duplicateByIndex(index int, c int32) string {
	var genString string
	for i := 0; i < index+1; i++ {
		genString += string(c)
	}
	return genString
}
