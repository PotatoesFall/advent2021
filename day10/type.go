package main

// Line is a line of input
type Line []byte

// Scores contains either the Corruption or Completion score of a line
type Scores struct {
	Corruption int
	Completion int
}

// Chunk gives information about a type of chunk
type Chunk interface {
	OpenChar() byte
	CloseChar() byte
	CorruptionScore() int
	CompletionScore() int
}

func parseOpeningChar(char byte) (Chunk, bool) {
	switch char {
	case '(':
		return Parentheses{}, true
	case '[':
		return Brackets{}, true
	case '{':
		return Braces{}, true
	case '<':
		return Chevrons{}, true
	}

	return nil, false
}

func allChunkTypes() []Chunk {
	return []Chunk{
		Parentheses{},
		Brackets{},
		Braces{},
		Chevrons{},
	}
}

// Parentheses is a type of chunk
type Parentheses struct{}

// OpenChar implements Chunk.OpenChar
func (Parentheses) OpenChar() byte {
	return '('
}

// CloseChar implements Chunk.Parentheses)
func (Parentheses) CloseChar() byte {
	return ')'
}

// CorruptionScore implements Chunk.Parentheses)
func (Parentheses) CorruptionScore() int {
	return 3
}

// CompletionScore implements Chunk.Parentheses)
func (Parentheses) CompletionScore() int {
	return 1
}

// Brackets is a type of Chunk
type Brackets struct{}

// OpenChar implements Chunk.OpenChar
func (b Brackets) OpenChar() byte {
	return '['
}

// CloseChar implements Chunk.CloseChar
func (b Brackets) CloseChar() byte {
	return ']'
}

// CorruptionScore implements Chunk.CorruptionScore
func (b Brackets) CorruptionScore() int {
	return 57
}

// CompletionScore implements Chunk.CompletionScore
func (b Brackets) CompletionScore() int {
	return 2
}

// Braces is a type of Chunk
type Braces struct{}

// OpenChar implements Chunk.OpenChar
func (Braces) OpenChar() byte {
	return '{'
}

// CloseChar implements Chunk.CloseChar
func (Braces) CloseChar() byte {
	return '}'
}

// CorruptionScore implements Chunk.CorruptionScore
func (Braces) CorruptionScore() int {
	return 1197
}

// CompletionScore implements Chunk.CompletionScore
func (Braces) CompletionScore() int {
	return 3
}

// Chevrons is a type of Chunk
type Chevrons struct{}

// OpenChar implements Chunk.OpenChar
func (Chevrons) OpenChar() byte {
	return '<'
}

// CloseChar implements Chunk.CloseChar
func (Chevrons) CloseChar() byte {
	return '>'
}

// CorruptionScore implements Chunk.CorruptionScore
func (Chevrons) CorruptionScore() int {
	return 25137
}

// CompletionScore implements Chunk.CompletionScore
func (Chevrons) CompletionScore() int {
	return 4
}
