package main

import (
	"github.com/svenwiltink/youtube-dl"
	"vitess.io/vitess/go/vt/log"
	"fmt"
	"flag"
)

func main() {
	flag.Parse()

	data, err := youtubedl.GetMetaData("https://www.youtube.com/watch?v=BdJgwf-_HGY")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", data)
}
