package tokenizer

type TokenType int

type Token struct {
	Path  string
	Type  TokenType
	Value string
	Line  int
	Col   int
}

type TokenKind int

const (
	KindKeyWord = iota
	KindIdentify
	KindDigits
	KindOperator
	KindSplit
)

type TokenPattern struct {
}

// TokenSet to predefine token patterns
type TokenSet struct {
}
