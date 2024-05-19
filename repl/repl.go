package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		i := lexer.New(line)


		for tok := i.NextToken(); tok.Type != token.EOF; tok = i.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}