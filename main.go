package main

import (
	"corelib"
	"flag"
	"log"
	"os"
)

var shouldExtract *bool
var source *string
var level *int

func readFlags()  {
	shouldExtract = flag.Bool("extract", false, "Boolean to indicate if the file needs be extracted. By default, it will try to compress.")
	source = flag.String("src", "", "Source file path (Required)")
	level = flag.Int("level", 9, "Compression level [0-9]. 0 -> Storage, 9 -> Best Compression. (Defaults to 9)")
	flag.Parse()
	if *source == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	isLevelWithinThreshold := *level >= 1 && *level <= 9
	if !isLevelWithinThreshold {
		log.Println("Error: Invalid values for compression level. Should be between 0 (storage) and 9 (best).")
		os.Exit(1)
	}
}

func main()  {
	log.SetFlags(0)
	readFlags()
	if *shouldExtract {
		corelib.ExtractZip(*source)
	} else {
		corelib.CreateArchive(*source, *level)
	}
}