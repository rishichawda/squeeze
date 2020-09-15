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
var useGZip *bool
var storageOnly *bool

func readFlags()  {
	shouldExtract = flag.Bool("extract", false, "Boolean to indicate if the file needs be extracted. By default, it will try to compress.")
	useGZip = flag.Bool("gzip", false, "Use GZip compression")
	source = flag.String("src", "", "Source file path (Required)")
	storageOnly = flag.Bool("store", false, "Don't compress zip file.")
	level = flag.Int("level", 9, "Compression level. 0 -> Storage, 9 -> Best Compression. (Defaults to 9)")
	flag.Parse()
	if *source == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if(*useGZip && *storageOnly) {
		log.Println("Invalid option -store. GZip doesn't support store flag.")
		os.Exit(1)
	}
	isLevelWithinThreshold := *level == 1 || *level == 9
	if !isLevelWithinThreshold {
		log.Println("Error: Invalid values for compression level. Should be 0 (storage) or 9 (best).")
		os.Exit(1)
	}
}

func main()  {
	log.SetFlags(0)
	readFlags()
	if *shouldExtract {
		if *useGZip {
			log.Println("here")
			corelib.ExtractGZip(*source)
		} else {
			corelib.ExtractZip(*source)
		}
	} else {
		if *useGZip {
			corelib.CreateGZip(*source, *level)
		} else {
			corelib.CreateArchive(*source, *storageOnly)
		}
	}
}