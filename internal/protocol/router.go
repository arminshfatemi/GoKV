package protocol

import (
	"bytes"
)

func Match(tokens [][]byte) (*MatchResult, error) {
	if len(tokens) == 0 {
		return nil, ErrEmptyCommand
	}

	var bestSpec CommandSpec
	bestDepth := 0

	var walk func(n *Node, idx int)
	walk = func(n *Node, idx int) {
		if n.Spec != nil {
			bestSpec = n.Spec
			bestDepth = idx
		}
		if idx >= len(tokens) {
			return
		}
		for _, c := range n.Children {
			if bytes.EqualFold(c.Key, tokens[idx]) {
				walk(c, idx+1)
			}
		}
	}

	walk(rootNode, 0)

	if bestSpec == nil {
		return nil, ErrInvalidCommand
	}

	return &MatchResult{
		Spec: bestSpec,
		Args: tokens[bestDepth:],
	}, nil
}
