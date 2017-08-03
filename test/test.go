package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	scanLines()
	// scanAll()
}

func scanAll() {
	b, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(b)
}

func scanLines() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(n)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return
}
