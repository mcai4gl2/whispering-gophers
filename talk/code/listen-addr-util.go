package main

import (
	"log"

	"github.com/mcai4gl2/whispering-gophers/util"
)

func main() {
	l, err := util.Listen()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on", l.Addr())
}
