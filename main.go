package main

import (
	"corelib"
	"flag"
	"log"
	"os"
)

var shouldExtract *bool

func readFlags()  {
	shouldExtract = flag.Bool("extract", false, "Boolean to indicate if the file needs be extracted. By default, it will try to compress.")
	flag.Parse()
}

func main()  {
	if len(os.Args) == 1 {
		log.SetFlags(6)
		log.SetPrefix("Error: ")
		log.Fatal("Needs a file path!")
	}
	readFlags()
	if *shouldExtract {
		corelib.ExtractZip()
	} else {
		corelib.CompressToZip()
	}
}