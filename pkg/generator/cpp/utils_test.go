package cpp

import (
	"testing"

	"github.com/Jusonex/RELang/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestCppFunctionParametersToString(t *testing.T) {
	s := FunctionParametersToString([]model.Parameter{
		model.Parameter{Name: "param1", Type: "int8"},
		model.Parameter{Name: "param2", Type: "int16"},
	})

	assert.Equal(t, "int8 param1, int16 param2", s)
}

func TestCppFunctionParameterNamesToString(t *testing.T) {
	s := FunctionParameterNamesToString([]model.Parameter{
		model.Parameter{Name: "param1", Type: "int8"},
		model.Parameter{Name: "param2", Type: "int16"},
	}, false)

	assert.Equal(t, "param1, param2", s)
}

func TestCppFunctionParameterTypesToString(t *testing.T) {
	s := FunctionParameterTypesToString([]model.Parameter{
		model.Parameter{Name: "param1", Type: "int8"},
		model.Parameter{Name: "param2", Type: "int16"},
	})

	assert.Equal(t, "int8, int16", s)
}
