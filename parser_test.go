package goantlr

import (
	"fmt"
	"testing"
)

func TestParser_Parser(t *testing.T) {
	ps := NewParser()
	err := ps.Parser(`#BFDA# + #ABC#`)
	fmt.Println(err)
}
