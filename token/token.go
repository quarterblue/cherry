package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Illegal
	ILLEGAL = "ILLEGAL" //
	EOF     = "EOF"

	INDENT = "INDENT"
	INT    = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delim
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// func LookupIdent(ident string) TokenType {
// 	if tok, ok := keywords[ident]; ok {
// 		return tok
// 	}
// 	return IDENT
// }
