package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/arbovm/highlighter"
	"io"
	"io/ioutil"
	"os"
)

func readLine(reader *bufio.Reader) (string, error) {
	return reader.ReadString('\n')
}

func main() {
	filename := fmt.Sprintf("/tmp/pipelight-%v", os.Getppid())
	file, oldInErr := os.Open(filename)
	defer file.Close()
	oldIn := bufio.NewReader(file)
	in := bufio.NewReader(os.Stdin)
	var oldLine string
	var lines bytes.Buffer

	for line, err := readLine(in); err != io.EOF; line, err = readLine(in) {
		lines.WriteString(line)
		if oldInErr == nil {
			oldLine, oldInErr = readLine(oldIn)
		}
		if oldInErr == nil {
			fmt.Print(highlighter.DiffLines(oldLine, line))
		} else {
			fmt.Print(line)
		}
	}
	ioutil.WriteFile(filename, lines.Bytes(), 0644)
}
