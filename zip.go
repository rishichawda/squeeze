package squeeze

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateArchive(file_path string, noCompress bool) {
	start := time.Now()
	_, input_filename := filepath.Split(file_path)
	output_file, err := os.Create(input_filename + ".zip")
	defer output_file.Close()
	cleanup := func() {
		os.Remove(output_file.Name())
	}
	LogIfError(err, true, func ()  {})
	file_writer := zip.NewWriter(output_file)
	defer file_writer.Close()
	filepath.Walk(file_path, func (path string, file_info os.FileInfo, err error) error {
		if !file_info.IsDir() {
			file_reader, err := os.Open(path)
			LogIfError(err, true, cleanup)
			defer file_reader.Close()
			file_info, err := file_reader.Stat()
			LogIfError(err, true, cleanup)
			file_header, err := zip.FileInfoHeader(file_info)
			LogIfError(err, true, cleanup)
			file_header.Name = strings.TrimPrefix(path, input_filename + "/")
			file_header.Method = map[bool]uint16{true: zip.Store, false: zip.Deflate}[noCompress]
			created_file, err := file_writer.CreateHeader(file_header)
			LogIfError(err, true, cleanup)
			_, err = io.Copy(created_file, file_reader)
			LogIfError(err, true, cleanup)
		}
		return nil
	})
	log.Printf("Done in %v!\n", time.Since(start))
}

func ExtractZip(file_path string)  {
	_, input_filename := filepath.Split(file_path)
	output_dirname:= strings.TrimSuffix(input_filename, filepath.Ext(input_filename))
	cleanup := func() {
		os.RemoveAll(output_dirname)
	}
	LogIfError(os.MkdirAll(output_dirname, os.ModePerm), true, cleanup)
	reader, err := zip.OpenReader(file_path)
	LogIfError(err, true, cleanup)
	defer reader.Close()
	for _, file := range reader.File {
		output_filepath := filepath.Join(output_dirname, file.Name)
		log.SetPrefix("Extracting: ")
		log.Println(output_filepath)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filepath.Join(output_dirname, file.Name), os.ModePerm)
		} else {
			// mkdirall was required to create directory before creating the file. Other zips were extracted earlier without any errors because they preserved order.
			// files created with this package don't preserve order, hence there are scenarios where directory is not present / already created before a contained file is written.
			os.MkdirAll(filepath.Dir(output_filepath), os.ModePerm)
			file_reader, err := file.Open()
			defer file_reader.Close()
			LogIfError(err, true, cleanup)
			file_writer, err := os.Create(output_filepath)
			defer file_writer.Close()
			LogIfError(err, true, cleanup)
			_, err = io.Copy(file_writer, file_reader)
			LogIfError(err, true, cleanup)
		}
	}
	log.SetPrefix("")
	log.Println("Done!")
}