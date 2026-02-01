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

var (
	cmdCREATE   = []byte("CREATE")
	cmdDROP     = []byte("DROP")
	cmdLIST     = []byte("LIST")
	cmdDESCRIBE = []byte("DESCRIBE")
	cmdSET      = []byte("SET")
	cmdGET      = []byte("GET")
	cmdDEL      = []byte("DEL")
	cmdINCR     = []byte("INCR")
	cmdSTATS    = []byte("STATS")
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
	case bytes.EqualFold(tokens[0], cmdCREATE):
		return parseCreate(tokens)

	case bytes.EqualFold(tokens[0], cmdDROP):
		return parseDrop(tokens)

	case bytes.EqualFold(tokens[0], cmdLIST):
		return parseList(tokens)

	case bytes.EqualFold(tokens[0], cmdDESCRIBE):
		return parseDescribe(tokens)

	case bytes.EqualFold(tokens[0], cmdSET):
		return parseSet(tokens)

	case bytes.EqualFold(tokens[0], cmdGET):
		return parseGet(tokens)
		//
	case bytes.EqualFold(tokens[0], cmdDEL):
		return parseDel(tokens)

	case bytes.EqualFold(tokens[0], cmdINCR):
		return parseIncr(tokens)

	case bytes.EqualFold(tokens[0], cmdSTATS):
		return parseStats(tokens)
	}

	return nil, ErrInvalidCommand
}
