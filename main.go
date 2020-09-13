package main

import (
	"archive/zip"
	"io"
	"log"
	"logger"
	"os"
	"path/filepath"
	"strings"
)

func main()  {
	if len(os.Args) == 1 {
		log.SetFlags(6)
		log.SetPrefix("Error: ")
		log.Fatal("Needs a file path!")
	}
	file_path := os.Args[1]
	reader, err := zip.OpenReader(file_path)
	logger.LogIfError(err, true, func ()  {})
	_, input_filename := filepath.Split(file_path)
	output_dirname:= strings.TrimSuffix(input_filename, filepath.Ext(input_filename))
	cleanup := func() {
		os.Remove(output_dirname)
	}
	logger.LogIfError(err, true, cleanup)
	logger.LogIfError(os.Mkdir(output_dirname, os.ModePerm), true, cleanup)
	for _, file := range reader.File {
		output_filepath := filepath.Join(output_dirname, file.Name)
		log.Println(output_filepath)
		log.Println(file.Name)
		if file.FileInfo().IsDir() {
			os.Mkdir(filepath.Join(output_dirname, file.Name), os.ModePerm)
		} else {
			file_reader, err := file.Open()
			logger.LogIfError(err, true, cleanup)
			file_writer, err := os.Create(output_filepath)
			logger.LogIfError(err, true, cleanup)
			io.Copy(file_writer, file_reader)
			logger.LogIfError(file_reader.Close(), true, cleanup)
		}
	}
}