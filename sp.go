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
}

type abs_node interface {
	get_type() int
	get_val() string
	construct(tokens []lexer.Token) (abs_node, error)
}

const (
	nt_program  = iota
	nt_block    = iota
	nt_var      = iota
	nt_funccall = iota
	nt_int      = iota
	nt_string   = iota
)

type node_program struct {
	name  string
	block abs_node
}

func (p node_program) get_type() int {
	return nt_program
}

func (p node_program) get_val() string {
	return p.name
}

func (p node_program) construct(tokens []lexer.Token) (abs_node, error) {
	return nil, nil
}

type node_int struct {
	name  string
	value int
}

func (p node_int) get_type() int {
	return nt_int
}

func (p node_int) get_val() string {
	return p.name
}

func (p node_int) construct(tokens []lexer.Token) (abs_node, error) {
	if len(tokens) != 1 {
		return nil, errors.New("expected one token")
	}
	switch tokens[0].Tokentype {
	case 1:
		return node_int{string(tokens[0].Val), 0}, nil
	}
	return nil, errors.New("expected number")
}
