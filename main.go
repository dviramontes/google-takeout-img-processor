package main

import (
	"path/filepath"
	"io/ioutil"
	"fmt"
	"log"
	"os"
	"io"
)

func main()  {
	fmt.Printf("Processing files beginning at: %s \n", os.Getenv("root"))
	files, err := ioutil.ReadDir(os.Getenv("root"))
	if err != nil {
		log.Panic(err)
	}

	//dest := os.Getenv("dest")
	//mkdirErr := os.Mkdir(dest, 0777)
	//if mkdirErr != nil {
	//	log.Println(mkdirErr)
	//}


	walk(files, os.Getenv("dest"))
}

func isJPG(file os.FileInfo) bool {
	ext := filepath.Ext(file.Name())
	return ext == ".jpg"
}

func walk(files []os.FileInfo, dest string) {
	root := os.Getenv("root")
	var ls []os.FileInfo
	for _, file := range files {
		absPath := root + "/" + file.Name()
		log.Println(absPath)
		if file.IsDir() {
			if isEmptyDir, _ := IsEmpty(absPath); isEmptyDir == true {
				continue
			}
			ls, _ = ioutil.ReadDir(absPath)
			log.Println(ls)
		}
	}
	if len(ls) > 0 {
		walk(ls, dest)
	}
}

// Copy the src file to dest
// Any existing file will be overwritten and will not copy file attributes.
func Copy(src, dest string) error {
	fmt.Println(src)
	fmt.Println(dest)
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}