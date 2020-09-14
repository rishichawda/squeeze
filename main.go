package main

import (
	"corelib"
	"flag"
	"log"
	"os"
)

var shouldExtract *bool
var source *string

func readFlags()  {
	shouldExtract = flag.Bool("extract", false, "Boolean to indicate if the file needs be extracted. By default, it will try to compress.")
	source = flag.String("src", "", "Source file path (Required)")
	flag.Parse()
	if *source == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main()  {
	log.SetFlags(0)
	readFlags()
	if *shouldExtract {
		corelib.ExtractZip(*source)
	} else {
		corelib.CompressToZip(*source)
	}
}