package token

type TokenType string

type Token struct {
    Type  TokenType
    Literal  string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // IDENTIFIERS
    NAME = "NAME"
    
    // TYPES
    STRING = "STRING"
    INTEGER = "INT"
    BOOLEAN = "BOOL"

    // DELIMITERS
    COLON = ":"
    COMMA = ","

    // OPERATORS
    PLUS = "+"
    ASSIGN = "="
    MINUS = "-"
    MULTIPLY = "*"
    DIVIDE = "/"

    // KEYWORDS
    FUNC = "func"
)
