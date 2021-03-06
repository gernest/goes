package lexer

import (
	"bytes"
	"fmt"
)

type nullLexer struct{}

func (nullLexer) Name() string {
	return "null"
}

func (nullLexer) Accept(s scanner) bool {
	n, _, err := s.Peek()
	if err != nil {
		return false
	}
	if n == 'n' {
		var b bytes.Buffer
		b.WriteRune(n)
		for i := 2; i < 5; i++ {
			n, _, err := s.PeekAt(i)
			if err != nil {
				return false
			}
			b.WriteRune(n)
		}
		return b.String() == "null"
	}
	return false
}

func (n nullLexer) Lex(s scanner, ctx *context) (*token, error) {
	var start, end position
	if ctx.lastToken != nil {
		start, end = ctx.lastToken.End, start
	}
	nx, w, err := s.Next()
	if err != nil {
		return nil, err
	}
	chrs := string(nx)
	end.Column += w
	for i := 0; i < 3; i++ {
		nx, w, err = s.Next()
		if err != nil {
			return nil, err
		}
		end.Column += w
		chrs += string(nx)
	}
	if chrs == "null" {
		return &token{
			Text:  chrs,
			Kind:  NULL,
			Start: start,
			End:   end,
		}, nil
	}
	return nil, fmt.Errorf(unexpectedTkn, n.Name(), end)
}
