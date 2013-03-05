package main

import (
	"fmt"
	"github.com/arbovm/highlighter"
	"bufio"
	"os"
	"io/ioutil"
	"io"
	"bytes"
)

func main() {
	filename := fmt.Sprintf("/tmp/pipelight-%v", os.Getppid())
	file, readError := os.Open(filename)
	stdinReader     := bufio.NewReader(os.Stdin)
	reader          := bufio.NewReader(file)
	var oldLine string
	var lines bytes.Buffer
	for {
		line,  err := stdinReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if readError == nil {
			oldLine, readError = reader.ReadString('\n')
		}

		if readError == nil {
			fmt.Print(highlighter.DiffLines(oldLine, line))
		} else {
			fmt.Print(line)
		}
		lines.WriteString(line)
	}
	if readError == nil {
		file.Close()
	}
	ioutil.WriteFile(filename, lines.Bytes(), 0644)
	
}

