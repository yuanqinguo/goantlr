package goantlr

import (
	"context"
	"fmt"
	"testing"
)

func TestParser_Parser(t *testing.T) {
	ps, _ := NewParser(context.TODO())
	err := ps.Parser(`#人事信息.职务# + #ABC#`)
	fmt.Println(err)
}
