// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	for reader.Scan() {
		err := encoder.Encode(Message{reader.Text()})
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
}
