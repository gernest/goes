package lexer

import (
	"bytes"
	"fmt"
)

type boolLexer struct{}

func (boolLexer) Name() string {
	return "boolean"
}

func (boolLexer) Accept(s scanner) bool {
	n, _, err := s.Peek()
	if err != nil {
		return false
	}
	var b bytes.Buffer
	switch n {
	case 't':
		b.WriteRune(n)
		for i := 2; i < 5; i++ {
			n, _, err := s.PeekAt(i)
			if err != nil {
				return false
			}
			b.WriteRune(n)
		}
		return b.String() == "true"
	case 'f':
		b.WriteRune(n)
		for i := 2; i < 6; i++ {
			n, _, err := s.PeekAt(i)
			if err != nil {
				return false
			}
			b.WriteRune(n)
		}
		return b.String() == "false"
	default:
		return false
	}
}

func (b boolLexer) Lex(s scanner, ctx *context) (*token, error) {
	var start, end position
	if ctx.lastToken != nil {
		start, end = ctx.lastToken.End, start
	}
	nx, w, err := s.Next()
	if err != nil {
		return nil, err
	}
	end.Column += w
	if nx == 't' || nx == 'f' {
		limit := 3
		if nx == 'f' {
			limit = 4
		}
		var b bytes.Buffer
		b.WriteRune(nx)
		for i := 0; i < limit; i++ {
			nx, w, err = s.Next()
			if err != nil {
				return nil, err
			}
			end.Column += w
			b.WriteRune(nx)
		}
		tk := &token{Text: b.String(), Start: start, End: end}
		switch b.String() {
		case "true":
			tk.Kind = TRUE
			return tk, nil
		case "false":
			tk.Kind = FALSE
			return tk, nil
		}
	}
	return nil, fmt.Errorf(unexpectedTkn, b.Name(), end)
}
