// simple standard pascal compiler to assembly
// info based off from https://wiki.freepascal.org/Standard_Pascal and places idk

package main

import (
	"errors"
	"flag"
	"fmt"
	"unicode"
	"os"
)

func main() {
	flag.Parse()
	files := flag.Args()
	for _, file := range files {
		fmt.Println(file)
		filetext, err := os.ReadFile(file)

		if err != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println(err)
		}
		f, err := lex(filetext)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("[")
		for _, tok := range f {
      fmt.Printf(fmt.Sprintf("type: %12.12s with %2.2d length: %s\n", converter(tok.tokentype), len(tok.val), tok.val))
		}
		fmt.Println("]")
	}
}

type token struct {
	startpos  int
	endpos    int
	val       []byte
	tokentype int
}

func converter(aid int) string {
	switch aid {
	case kw_and:
		return "and"
	case kw_array:
		return "array"
	case kw_begin:
		return "begin"
	case kw_case:
		return "case"
	case kw_const:
		return "const"
	case kw_div:
		return "div"
	case kw_do:
		return "do"
	case kw_downto:
		return "downto"
	case kw_else:
		return "else"
	case kw_end:
		return "end"
	case kw_file:
		return "file"
	case kw_for:
		return "for"
	case kw_function:
		return "function"
	case kw_goto:
		return "goto"
	case kw_if:
		return "if"
	case kw_in:
		return "in"
	case kw_label:
		return "label"
	case kw_mod:
		return "mod"
	case kw_nil:
		return "nil"
	case kw_not:
		return "not"
	case kw_of:
		return "of"
	case kw_or:
		return "or"
	case kw_packed:
		return "packed"
	case kw_procedure:
		return "procedure"
	case kw_program:
		return "program"
	case kw_record:
		return "record"
	case kw_repeat:
		return "repeat"
	case kw_set:
		return "set"
	case kw_then:
		return "then"
	case kw_to:
		return "to"
	case kw_type:
		return "type"
	case kw_until:
		return "until"
	case kw_var:
		return "var"
	case kw_while:
		return "while"
	case kw_with:
		return "with"
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
	kw_and       = iota
	kw_array     = iota
	kw_begin     = iota
	kw_case      = iota
	kw_const     = iota
	kw_div       = iota
	kw_do        = iota
	kw_downto    = iota
	kw_else      = iota
	kw_end       = iota
	kw_file      = iota
	kw_for       = iota
	kw_function  = iota
	kw_goto      = iota
	kw_if        = iota
	kw_in        = iota
	kw_label     = iota
	kw_mod       = iota
	kw_nil       = iota
	kw_not       = iota
	kw_of        = iota
	kw_or        = iota
	kw_packed    = iota
	kw_procedure = iota
	kw_program   = iota
	kw_record    = iota
	kw_repeat    = iota
	kw_set       = iota
	kw_then      = iota
	kw_to        = iota
	kw_type      = iota
	kw_until     = iota
	kw_var       = iota
	kw_while     = iota
	kw_with      = iota
)

func lex(file []byte) ([]token, error) {
	tokens := []token{}
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
			// hell generated from shitty snake language
			case "and":
				ttype = kw_and
			case "array":
				ttype = kw_array
			case "begin":
				ttype = kw_begin
			case "case":
				ttype = kw_case
			case "const":
				ttype = kw_const
			case "div":
				ttype = kw_div
			case "do":
				ttype = kw_do
			case "downto":
				ttype = kw_downto
			case "else":
				ttype = kw_else
			case "end":
				ttype = kw_end
			case "file":
				ttype = kw_file
			case "for":
				ttype = kw_for
			case "function":
				ttype = kw_function
			case "goto":
				ttype = kw_goto
			case "if":
				ttype = kw_if
			case "in":
				ttype = kw_in
			case "label":
				ttype = kw_label
			case "mod":
				ttype = kw_mod
			case "nil":
				ttype = kw_nil
			case "not":
				ttype = kw_not
			case "of":
				ttype = kw_of
			case "or":
				ttype = kw_or
			case "packed":
				ttype = kw_packed
			case "procedure":
				ttype = kw_procedure
			case "program":
				ttype = kw_program
			case "record":
				ttype = kw_record
			case "repeat":
				ttype = kw_repeat
			case "set":
				ttype = kw_set
			case "then":
				ttype = kw_then
			case "to":
				ttype = kw_to
			case "type":
				ttype = kw_type
			case "until":
				ttype = kw_until
			case "var":
				ttype = kw_var
			case "while":
				ttype = kw_while
			case "with":
				ttype = kw_with
			}
			tokens = append(tokens, token{startpos, i, file[startpos:i], ttype})
			i--
		case unicode.IsNumber(rune(k)):
			// number
			startpos := i
			for ; i < len(file); i++ {
				if file[i] > '9' || file[i] < '0' {
					break
				}
			}
			tokens = append(tokens, token{startpos, i, file[startpos:i], number})
			i--
		case k == '[':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], leftsquarebracket})
		case k == ']':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], rightsquarebracket})
		case k == '^':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], uparrow})
		case k == '<':
			if i+1 <= len(file) {
				if file[i+1] == '=' {
					tokens = append(tokens, token{i, i + 2, file[i : i+2], leftequal})
					i++
				} else if file[i+1] == '>' {
					tokens = append(tokens, token{i, i + 2, file[i : i+2], notequal})
					i++
				} else {
					tokens = append(tokens, token{i, i + 1, file[i : i+1], leftbracket})
				}
			} else {
				tokens = append(tokens, token{i, i + 1, file[i : i+1], leftbracket})
			}
		case k == '>':
			if i+1 <= len(file) {
				if file[i+1] == '=' {
					tokens = append(tokens, token{i, i + 2, file[i : i+2], rightequal})
					i++
				} else {
					tokens = append(tokens, token{i, i + 1, file[i : i+1], rightbracket})
				}
			} else {
				tokens = append(tokens, token{i, i + 1, file[i : i+1], rightbracket})
			}
		case k == ' ' || k == '\n' || k == '\t':
			// whitespace
		case k == '+':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], plus})
		case k == '-':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], minus})
		case k == '*':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], asterisk})
		case k == '/':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], slash})
		case k == ';':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], semicolon})
		case k == '(':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], leftparen})
		case k == ')':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], rightparen})
		case k == '=':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], equals})
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
			tokens = append(tokens, token{startpos, i + 1, file[startpos : i+1], sstring})
		case k == ':':
			if i+1 >= len(file) {
				return nil, errors.New(fmt.Sprintf("unexpected end of file at position %d", i))
			}
			if file[i+1] == '=' {
				tokens = append(tokens, token{i, i + 2, file[i : i+2], becomes})
				i++
			} else {
				tokens = append(tokens, token{i, i + 1, file[i : i+1], colon})
			}
		case k == ',':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], comma})
		case k == '.':
			tokens = append(tokens, token{i, i + 1, file[i : i+1], period})
		default:
			// unexpected character
			return nil, errors.New(fmt.Sprintf("unexpected character %c at position %d", k, i))
		}
	}
	return tokens, nil
}

type abs_node interface {
  get_type() int
  get_val() string
  construct(tokens []token) (abs_node, error)
}

const (
  nt_program = iota
  nt_block = iota
  nt_var = iota
  nt_funccall = iota
  nt_int = iota
  nt_string = iota
)

type node_program struct {
  name string
  block abs_node
}

func (p node_program) get_type() int {
  return nt_program
}

func (p node_program) get_val() string {
  return p.name
}

func (p node_program) construct(tokens []token) (abs_node, error) {
  return nil, nil
}

type node_int struct {
  name string
  value int
}

func (p node_int) get_type() int {
  return nt_int
}

func (p node_int) get_val() string {
  return p.name
}

func (p node_int) construct(tokens []token) (abs_node, error) {
  if len(tokens) != 1 {
    return nil, errors.New("expected one token")
  }
  switch tokens[0].tokentype {
  case number:
    return node_int{string(tokens[0].val), 0}, nil
  }
  return nil, errors.New("expected number")
}


