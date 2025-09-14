package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/huavcjj/monkey-playground/lexer"
	"github.com/huavcjj/monkey-playground/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprint(out, PROMPT); err != nil {
			return
		}
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if _, err := fmt.Fprintf(out, "%+v\n", tok); err != nil {
				return
			}
		}
	}
}
