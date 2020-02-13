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
