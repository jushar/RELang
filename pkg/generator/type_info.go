package generator

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

const POINTER_SIZE uint8 = 4
