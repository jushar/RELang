package generator

import "github.com/Jusonex/RELang/pkg/model"

// Transformer defines an interface for chunk transform operations
type Transformer interface {
	Transform(chunk *model.Chunk)
}
