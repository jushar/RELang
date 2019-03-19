package cpp

import (
	"testing"

	"github.com/Jusonex/RELang/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestTransformsInbuiltTypes(t *testing.T) {
	assert.Equal(t, "std::uint8_t", TransformInbuiltType("uint8")) // Transformation available
	assert.Equal(t, "Element*", TransformInbuiltType("Element*"))  // No transformation available
}

func TestVariableTypeTransformations(t *testing.T) {
	variable := model.NewVariable()
	variable.Type = "int8"

	ApplyVariableTypeTransformations(variable)
	assert.Equal(t, "std::int8_t", variable.Type)
}

func TestFunctionTypeTransformations(t *testing.T) {
	function := model.NewFunction("test")
	function.ReturnType = "float32"
	function.AddParameter("param1", "float64")

	ApplyFunctionTypeTransformations(function)
	assert.Equal(t, "float", function.ReturnType)
	assert.Equal(t, "double", function.Params[0].Type)
}
