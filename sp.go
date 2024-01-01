// simple standard pascal compiler to assembly
// info based off from https://wiki.freepascal.org/Standard_Pascal and places idk

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sp/lexer"
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
		f, err := lexer.Lex(filetext)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("[")
		for _, tok := range f {
			fmt.Printf(fmt.Sprintf("type: %12.12s with %2.2d length: %s\n", lexer.Converter(tok), len(tok.Val), tok.Val))
		}
		fmt.Println("]")
	}
	// try parsing example expr
	expr := "1 + 2 * 3"
	tok, err := lexer.Lex([]byte(expr))
	if err != nil {
		fmt.Println(err)
	}
	p := parser{tok, -1, ""}
	node, err := p.parse_expr()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(node.get_val())
}

type abs_node interface {
	get_type() int
	get_val() string
}

const (
	nt_program  = iota
	nt_block    = iota
	nt_var      = iota
	nt_funccall = iota
	nt_int      = iota
	nt_string   = iota
	nt_expr     = iota
)

type parser struct {
	tokens   []lexer.Token
	pos      int
	filename string
}

func (p *parser) next() (lexer.Token, error) {
	p.pos++
	if p.pos >= len(p.tokens) {
		return lexer.Token{Tokentype: -1}, errors.New("out of tokens")
	}
	return p.tokens[p.pos], nil
}

func (p *parser) peek() lexer.Token {
	return p.tokens[p.pos+1]
}

func (p *parser) parse() (abs_node, error) {
	return nil, nil
}

func (p *parser) parse_expr() (abs_node, error) {
	return p.actually_parse_expr(0)
}

func (p *parser) actually_parse_expr(min_bp int) (abs_node, error) {
	lhs, err := p.parse_atom()
	if err != nil {
		return nil, err
	}

	for p.pos < len(p.tokens) {
		op, err := p.next()
		if op.Tokentype == -1 {
			break
		}
		if err != nil {
			return nil, err
		}
		if op.Tokentype != 2 && op.Tokentype != 3 && op.Tokentype != 4 && op.Tokentype != 5 {
			break
		}
		lbp, rbp := calc_bp(op)
		if lbp < min_bp {
			break
		}
		rhs, err := p.actually_parse_expr(rbp)
		if err != nil {
			return nil, err
		}
		lhs = node_expr{op, lhs, rhs}
	}

	return lhs, nil
}

type node_expr struct {
	op  lexer.Token
	lhs abs_node
	rhs abs_node
}

func (n node_expr) get_type() int {
	return nt_expr
}

func (n node_expr) get_val() string {
	str := "( " + string(n.op.Val) + " " + n.lhs.get_val() + " " + n.rhs.get_val() + " )"
	return str
}

func calc_bp(op lexer.Token) (int, int) {
	switch op.Tokentype {
	case 2:
		return 1, 2
	case 3:
		return 1, 2
	case 4:
		return 3, 4
	case 5:
		return 3, 4
	}
	return 0, 0
}

func (p *parser) parse_atom() (abs_node, error) {
	t, err := p.next()
	if err != nil {
		return nil, err
	}
	if t.Tokentype == 1 {
		return node_int{string(p.tokens[p.pos].Val)}, nil
	}
	return nil, errors.New("not implemented, " + string(t.Val))
}

type node_int struct {
	val string
}

func (n node_int) get_type() int {
	return nt_int
}

func (n node_int) get_val() string {
	return n.val
}
