package protocol

import (
	"bytes"
	"errors"
)

var (
	ErrEmptyCommand   = errors.New("empty command")
	ErrInvalidCommand = errors.New("invalid command")
	ErrWrongArgCount  = errors.New("wrong number of arguments")
)

type Parser struct {
	tokenizer *Tokenizer
}

func NewParser() *Parser {
	return &Parser{
		tokenizer: NewTokenizer(),
	}
}

func (p *Parser) Parse(line []byte) (*Command, error) {
	tokens := p.tokenizer.Tokenize(line)
	if len(tokens) == 0 {
		return nil, ErrEmptyCommand
	}

	switch {
	case bytes.EqualFold(tokens[0], []byte("CREATE")):
		return parseCreate(tokens)

	case bytes.EqualFold(tokens[0], []byte("DROP")):
		return parseDrop(tokens)

	case bytes.EqualFold(tokens[0], []byte("LIST")):
		return parseList(tokens)

		//case bytes.EqualFold(tokens[0], []byte("DESCRIBE")):
		//	return parseDescribe(tokens)
		//
	case bytes.EqualFold(tokens[0], []byte("SET")):
		return parseSet(tokens)

	case bytes.EqualFold(tokens[0], []byte("GET")):
		return parseGet(tokens)
		//
	case bytes.EqualFold(tokens[0], []byte("DEL")):
		return parseDel(tokens)
		//
		//case bytes.EqualFold(tokens[0], []byte("INCR")):
		//	return parseIncr(tokens)
		//
		//case bytes.EqualFold(tokens[0], []byte("STATS")):
		//	return parseStats(tokens)
	}

	return nil, ErrInvalidCommand
}
