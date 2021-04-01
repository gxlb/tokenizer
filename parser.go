package tokenizer

type Parser struct {
	Path   string
	Buffer []byte
	P1     int
	P2     int
	Line   int
	Col    int
}
