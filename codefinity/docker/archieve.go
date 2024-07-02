package main

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	filepath := "test"
	test := "test.tar"
	createFileWithContent(filepath, "test")
	addFileToArchive(filepath, "tarfilename.tar")
	createEmptyArchive(test)
	addFileToArchive(filepath, test)
}

func createFileWithContent(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)

}

func rmFile(filename string) error {
	return os.Remove(filename)
}

func createRandomStr() string {
	return "random"
}

func createEmptyArchive(path string) {

	os.Create(path)

}

func addFileToArchive(path string, pathToArchive string) {
	w, _ := os.OpenFile(pathToArchive, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	tarWriter := tar.NewWriter(w)
	defer tarWriter.Close()
	r, _ := os.Open(path)
	defer r.Close()

	stat, _ := r.Stat()
	///fmt.Println(stat.Mode())
	header := &tar.Header{
		Name: path,
		// 	Size:    stat.Size(),
		Mode: int64(stat.Mode()),

		// 	ModTime: stat.ModTime(),
	}

	tarWriter.WriteHeader(header)

	io.Copy(tarWriter, r)

}
