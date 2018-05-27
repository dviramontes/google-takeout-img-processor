package main

import (
	"path/filepath"
	"io/ioutil"
	"fmt"
	"log"
	"os"
)

func isDir(file string) (bool, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

func isJPG(file os.FileInfo) bool {
	ext := filepath.Ext(file.Name())
	return ext == ".jpg"
}

func walk(files []os.FileInfo) {
	for _, file := range files {
		path := os.Getenv("root") + "/" + file.Name()
		if _, err := isDir(path); err == nil {
			files, _ := ioutil.ReadDir(path)
			walk(files)
		} else if isJPG(file) {
			fmt.Println(file.Name())
		}
	}
}

func main()  {
	fmt.Printf("Processing files beginning at: %s \n", os.Getenv("root"))
	files, err := ioutil.ReadDir(os.Getenv("root"))
	if err != nil {
		log.Panic(err)
	}

	walk(files)
}