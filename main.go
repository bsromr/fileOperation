package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var htmlFiles []string
var path string
var pureHTML []string

func main() {
	path = "/home/omrbsr/websites/deckkompozit.com/deckkompozit.com" //TODO: Get value from input
	changeDirectory(path)
	getHTMLFiles()
	readFile()
	fmt.Println(pureHTML) //TODO: Firstly, write the codes in file.
}

func readFile() {
	for _, f := range htmlFiles {
		filerc, err := os.Open(f)
		logErr(err)
		defer filerc.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(filerc)
		contents := buf.String()
		pureHTML = append(pureHTML, contents)
	}
}

func getHTMLFiles() []string {
	files, err := ioutil.ReadDir(path)
	logErr(err)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".html" {
			htmlFiles = append(htmlFiles, f.Name())
		}
	}
	return htmlFiles
}

func changeDirectory(path string) {
	os.Chdir(path)
	_, err := os.Getwd()
	logErr(err)
}

func logErr(err error) {
	fmt.Println(err)
}
