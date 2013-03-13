package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/arbovm/highlighter"
	"io"
	"io/ioutil"
	"os"
)

const (
	BUFFER_FILE = "/tmp/pipelight-%v"
)

var (
	html  = flag.Bool("html", false, "Use <b>..</b> instead of terminal escape chars for highlighting")
	start = flag.String("start", "\033[01;37m", "String to start highlighting")
	stop  = flag.String("stop", "\033[0m", "String stop highlighting")
	clear = flag.Bool("clear", false, "Clear buffer")
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Hightlights the differences of two subsequent command outputs\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample: \n")
		fmt.Fprintf(os.Stderr, "\tcurl -sI https://github.com | %s -clear && \\\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\tcurl -sI --compress https://github.com | %s\n", os.Args[0])
	}
	flag.Parse()

	if *html {
		*start, *stop = "<b>", "</b>"
	}

	filename := fmt.Sprintf(BUFFER_FILE, os.Getppid())
	if *clear {
		ioutil.WriteFile(filename, []byte{}, 0644)
	}
	file, oldInErr := os.Open(filename)
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
			fmt.Print(highlighter.DiffLines(oldLine, line, *start, *stop))
		} else {
			fmt.Print(line)
		}
	}
	file.Close()
	ioutil.WriteFile(filename, lines.Bytes(), 0644)
}

func readLine(reader *bufio.Reader) (string, error) {
	return reader.ReadString('\n')
}
