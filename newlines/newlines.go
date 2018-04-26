package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Read data from stdin (caution, buffers all currently)
	data, err := ioutil.ReadAll(os.Stdin)

	// Write error to stderr and exit, if any.
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	// Use strings.replacer to flip \n to newline literals and vice-versa
	replacer := strings.NewReplacer(`\n`, "\n", "\n", `\n`)
	fmt.Println(replacer.Replace(string(data)))
}
