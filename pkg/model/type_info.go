package model

import "errors"

// Sizes of inbuilt types
var TYPES_SIZE = map[string]uint8{
	"bool":    1,
	"int8":    1,
	"int16":   2,
	"int32":   4,
	"int64":   8,
	"uint8":   1,
	"uint16":  2,
	"uint32":  4,
	"uint64":  8,
	"float32": 4,
	"float64": 8,
}

// Size of a pointer for the IA-32 arch
const POINTER_SIZE uint8 = 4

// Returns the size of an inbuilt type
func GetInbuiltTypeSize(typeName string) (uint8, error) {
	if IsTypePointer(typeName) {
		return POINTER_SIZE, nil
	}

	if size, ok := TYPES_SIZE[typeName]; ok {
		return size, nil
	}

	return 0, errors.New("not an inbuilt type")
}

// IsTypePointer checks whether or not a type is a pointer
func IsTypePointer(typeName string) bool {
	return typeName[len(typeName)-1] == '*'
}

// GetTypeFromPointer returns the type name without the pointer asterisks
func GetTypeFromPointer(typeName string) string {
	if !IsTypePointer(typeName) {
		panic("not a pointer")
	}

	return typeName[:len(typeName)-1]
}
