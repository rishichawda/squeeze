package corelib

import (
	"archive/zip"
	"compress/flate"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateArchive(file_path string, compressionLevel int) {
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
	file_writer.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, compressionLevel)
	})
	filepath.Walk(file_path, func (path string, file_info os.FileInfo, err error) error {
		if !file_info.IsDir() {
			file_reader, err := os.Open(path)
			defer file_reader.Close()
			LogIfError(err, true, cleanup)
			created_file, err := file_writer.Create(path)
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
	LogIfError(os.Mkdir(output_dirname, os.ModePerm), true, cleanup)
	reader, err := zip.OpenReader(file_path)
	defer reader.Close()
	LogIfError(err, true, cleanup)
	for _, file := range reader.File {
		output_filepath := filepath.Join(output_dirname, file.Name)
		log.SetPrefix("Extracting: ")
		log.Println(output_filepath)
		if file.FileInfo().IsDir() {
			os.Mkdir(filepath.Join(output_dirname, file.Name), os.ModePerm)
		} else {
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