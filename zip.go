package squeeze

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ExtractZip()  {
	file_path := os.Args[1]
	_, input_filename := filepath.Split(file_path)
	output_dirname:= strings.TrimSuffix(input_filename, filepath.Ext(input_filename))
	cleanup := func() {
		os.Remove(output_dirname)
	}
	LogIfError(os.Mkdir(output_dirname, os.ModePerm), true, cleanup)
	reader, err := zip.OpenReader(file_path)
	LogIfError(err, true, cleanup)
	for _, file := range reader.File {
		output_filepath := filepath.Join(output_dirname, file.Name)
		log.Println(output_filepath)
		log.Println(file.Name)
		if file.FileInfo().IsDir() {
			os.Mkdir(filepath.Join(output_dirname, file.Name), os.ModePerm)
		} else {
			file_reader, err := file.Open()
			LogIfError(err, true, cleanup)
			file_writer, err := os.Create(output_filepath)
			LogIfError(err, true, cleanup)
			io.Copy(file_writer, file_reader)
			LogIfError(file_reader.Close(), true, cleanup)
		}
	}
}