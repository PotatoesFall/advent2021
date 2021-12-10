package main

import "fmt"

// Parser reads a line of chunks
type Parser struct {
	line       Line
	cursor     int
	openChunks []Chunk
}

// NewParser reads a line
func NewParser(line Line) *Parser {
	return &Parser{
		line: line,
	}
}

// Scores returns the corruption and completion score
// only either is ever not zero
func (lp *Parser) Scores() Scores {
	for _, char := range lp.line {
		// check if new chunk is starting
		chunk, ok := parseOpeningChar(char)
		if ok {
			lp.Push(chunk)
			lp.cursor++
			continue
		}

		// check if chunk is closing
		openChunk := lp.Pop()
		if openChunk.CloseChar() == char {
			lp.cursor++
			continue
		}

		// invalid --> get corruption score
		corruptScore := corruptionScore(char)
		return Scores{Corruption: corruptScore}
	}

	// incomplete --> get completion score
	completionScore := lp.completionScore()
	return Scores{Completion: completionScore}
}

func corruptionScore(char byte) int {
	for _, chunk := range allChunkTypes() {
		if chunk.CloseChar() == char {
			return chunk.CorruptionScore()
		}
	}

	panic(fmt.Sprintf(`invalid closing char: %c`, char))
}

func (lp Parser) completionScore() int {
	score := 0

	for len(lp.openChunks) != 0 {
		score *= 5
		score += lp.Pop().CompletionScore()
	}

	return score
}

// Pop pops a chunk, it should be used to check if that chunk is being closed correctly
func (lp *Parser) Pop() Chunk {
	top := lp.openChunks[len(lp.openChunks)-1]

	lp.openChunks = lp.openChunks[:len(lp.openChunks)-1]

	return top
}

// Push adds a new open chunk inside the previous chunk
func (lp *Parser) Push(chunk Chunk) {
	lp.openChunks = append(lp.openChunks, chunk)
}
