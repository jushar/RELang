package model

// Represents a raw block
type RawBlock struct {
	Content string
}

// Constructs a new raw block with the given content
func NewRawBlock(content string) *RawBlock {
	return &RawBlock{
		Content: content,
	}
}
