package lexer

import (
	"testing"

	"github.com/regent-x/norm/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `:,=+`

	test := []struct {
		wantType    token.TokenType
		wantLiteral string
	}{
		{token.COLON, ":"},
		{token.COMMA, ","},
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
	}

	l := New(input)

	for tt, tl := range test {
		l.NextToken()
		if l.Type != tl.wantType {
			t.Errorf("test[%d] - Want: %q, Got:%q", tt, tl.wantType, l.Type)
		}

		if l.Literal != tl.wantLiteral {
			t.Errorf("test[%d] - Want: %q, Got:%q", tt, tl.wantLiteral, l.Literal)
		}
	}
}
