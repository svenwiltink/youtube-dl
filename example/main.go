package main

import (
	"flag"
	"fmt"
	"github.com/svenwiltink/youtube-dl"
	"log"
)

func main() {
	flag.Parse()

	data, err := youtubedl.GetMetaData("https://www.youtube.com/watch?v=BdJgwf-_HGY")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", data)
}
