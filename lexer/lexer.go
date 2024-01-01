package lexer

import (
  "errors"
  "fmt"
  "unicode"
)

type Token struct {
	startpos  int
	endpos    int
	Val       []byte
	Tokentype int
}

func Converter(aid Token) string {
	switch aid.Tokentype {
  case kw_end:
    return "end"
	case kw_program:
		return "program"
	}
	return "non-keyword"
}

const (
	// normal tokens
	identifier = iota
	number     = iota
	// mathematic operators
	plus               = iota
	minus              = iota
	asterisk           = iota
	slash              = iota
	leftbracket        = iota
	leftequal          = iota
	rightbracket       = iota
	rightequal         = iota
	notequal           = iota
	uparrow            = iota
	leftsquarebracket  = iota
	rightsquarebracket = iota

	semicolon  = iota
	leftparen  = iota
	rightparen = iota
	colon      = iota
	equals     = iota
	becomes    = iota
	comma      = iota
	period     = iota
	sstring    = iota
	// keywords
  kw_program = iota
  kw_end = iota
	)

func Lex(file []byte) ([]Token, error) {
	tokens := []Token{}
	for i := 0; i < len(file); i++ {
		switch k := file[i]; {
		case unicode.IsLetter(rune(k)) && !unicode.IsSpace(rune(k)):
			// might be identifier
			startpos := i
			for ; i < len(file); i++ {
				if !(unicode.IsLetter(rune(file[i])) || unicode.IsDigit(rune(file[i]))) {
					break
				}
			}
			ttype := identifier
			switch string(file[startpos:i]) {
      case "end":
        ttype = kw_end
      case "program":
        ttype = kw_program
			}
			tokens = append(tokens, Token{startpos, i, file[startpos:i], ttype})
			i--
		case unicode.IsNumber(rune(k)):
			// number
			startpos := i
			for ; i < len(file); i++ {
				if file[i] > '9' || file[i] < '0' {
					break
				}
			}
			tokens = append(tokens, Token{startpos, i, file[startpos:i], number})
			i--
		case k == '[':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], leftsquarebracket})
		case k == ']':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], rightsquarebracket})
		case k == '^':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], uparrow})
		case k == '<':
			if i+1 <= len(file) {
				if file[i+1] == '=' {
					tokens = append(tokens, Token{i, i + 2, file[i : i+2], leftequal})
					i++
				} else if file[i+1] == '>' {
					tokens = append(tokens, Token{i, i + 2, file[i : i+2], notequal})
					i++
				} else {
					tokens = append(tokens, Token{i, i + 1, file[i : i+1], leftbracket})
				}
			} else {
				tokens = append(tokens, Token{i, i + 1, file[i : i+1], leftbracket})
			}
		case k == '>':
			if i+1 <= len(file) {
				if file[i+1] == '=' {
					tokens = append(tokens, Token{i, i + 2, file[i : i+2], rightequal})
					i++
				} else {
					tokens = append(tokens, Token{i, i + 1, file[i : i+1], rightbracket})
				}
			} else {
				tokens = append(tokens, Token{i, i + 1, file[i : i+1], rightbracket})
			}
		case k == ' ' || k == '\n' || k == '\t':
			// whitespace
		case k == '+':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], plus})
		case k == '-':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], minus})
		case k == '*':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], asterisk})
		case k == '/':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], slash})
		case k == ';':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], semicolon})
		case k == '(':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], leftparen})
		case k == ')':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], rightparen})
		case k == '=':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], equals})
		case k == '\'':
			// string
			startpos := i
			i++
			for file[i] != '\'' {
				i++
				if i >= len(file) {
					return nil, errors.New(fmt.Sprintf("unexpected end of file at position %d", i))
				}
			}
			tokens = append(tokens, Token{startpos, i + 1, file[startpos : i+1], sstring})
		case k == ':':
			if i+1 >= len(file) {
				return nil, errors.New(fmt.Sprintf("unexpected end of file at position %d", i))
			}
      if file[i+1] == ':' {
				tokens = append(tokens, Token{i, i + 2, file[i : i+2], becomes})
				i++
			} else {
				tokens = append(tokens, Token{i, i + 1, file[i : i+1], colon})
			}
		case k == ',':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], comma})
		case k == '.':
			tokens = append(tokens, Token{i, i + 1, file[i : i+1], period})
		default:
			// unexpected character
			return nil, errors.New(fmt.Sprintf("unexpected character %c at position %d", k, i))
		}
	}
	return tokens, nil
}


