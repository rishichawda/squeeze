package squeeze

import (
	"log"
	"os"
	"path/filepath"
)

func main()  {
	if len(os.Args) == 1 {
		log.SetFlags(6)
		log.SetPrefix("Error: ")
		log.Fatal("Needs a file path!")
	}
	file_path := os.Args[1]
	_, input_filename := filepath.Split(file_path)
	extension := filepath.Ext(input_filename)
	if extension == ".zip" {
		ExtractZip()
	}
}