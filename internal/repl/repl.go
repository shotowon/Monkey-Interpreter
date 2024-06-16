package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/internal/lexer"
	"monkey/internal/parser"
	"strings"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "quit" {
			return
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			p.PrintErrors(out)
			continue
		}

		io.WriteString(out, program.String()+"\n")
	}
}
