package TDD_string_convertor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSuiteInit(t *testing.T) {
	suite.Run(t, new(StrConverterTests))
}

type StrConverterTests struct {
	suite.Suite
}

func (suite *StrConverterTests) SetupTest() {

}

func (suite *StrConverterTests) Test_lower_to_upper() {
	assert.Equal(suite.T(), "A", converter("a"), "")
}

func (suite *StrConverterTests) Test_upper_to_upper() {
	assert.Equal(suite.T(), "A", converter("A"), "")
}

func (suite *StrConverterTests) Test_number_not_changed() {
	assert.Equal(suite.T(), "1", converter("1"), "")
}

func (suite *StrConverterTests) Test_second_char_duplicate_twice() {
	assert.Equal(suite.T(), "1-22", converter("12"), "")
}

func (suite *StrConverterTests) Test_char_duplicate_by_its_index() {
	assert.Equal(suite.T(), "1-22-333", converter("123"), "")
}

func (suite *StrConverterTests) Test_second_char_first_upper_second_lower() {
	assert.Equal(suite.T(), "A-Bb", converter("ab"), "")
}
