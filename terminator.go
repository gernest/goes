package goes

import "fmt"

type terminatorLexer struct{}

func (terminatorLexer) name() string {
	return "terminator"
}

func (terminatorLexer) accept(s scanner) bool {
	n, _, err := s.peek()
	if err != nil {
		return false
	}
	return isLineTerminator(n)
}

func isLineTerminator(ch rune) bool {
	switch ch {
	case 0x0000A, 0x000D, 0x02028, 0x2029:
		return true
	default:
		return false
	}
}

func (t terminatorLexer) lex(s scanner, ctx *context) (*token, error) {
	var start, end position
	if ctx.lastToken != nil {
		start, end = ctx.lastToken.end, start
	}
	n, _, err := s.next()
	if err != nil {
		return nil, err
	}
	if isLineTerminator(n) {
		end.line++
		end.column = 0
		tk := &token{
			kind:  lineTerminator,
			text:  string(n),
			start: start,
			end:   end}
		switch n {
		case 0x0000A:
			tk.kind = LF
		case 0x000D:
			tk.kind = CR
		case 0x02028:
			tk.kind = LS
		case 0x2029:
			tk.kind = PS
		}
		return tk, nil
	}
	return nil, fmt.Errorf(unexpectedTkn, t.name(), end)
}
