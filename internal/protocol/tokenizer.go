package protocol

type Tokenizer struct {
	tokens [][]byte
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		tokens: make([][]byte, 0, 8),
	}
}

// Tokenize splits a line into space-separated tokens.
// It reuses internal slice to avoid allocations.
func (t *Tokenizer) Tokenize(line []byte) [][]byte {
	t.tokens = t.tokens[:0]

	start := -1
	for i, b := range line {
		if b == ' ' || b == '\t' {
			if start != -1 {
				t.tokens = append(t.tokens, line[start:i])
				start = -1
			}
			continue
		}
		if start == -1 {
			start = i
		}
	}

	if start != -1 {
		t.tokens = append(t.tokens, line[start:])
	}

	return t.tokens
}
