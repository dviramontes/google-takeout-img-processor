package main

import (
	"path/filepath"
	"fmt"
	"log"
	"os"
	"io"
)

func main()  {
	root := os.Getenv("root")
	fmt.Printf("Processing files beginning at: %s \n", root)

	list := walk(root)
	fmt.Printf("Processed: %d jpgs\n", len(list))
}

func walk(root string) []string {
	var list []string
	dest := os.Getenv("dest")

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".jpg" {
			list = append(list, info.Name())
			copyErr := Copy(path, dest, info.Name())
			if copyErr != nil {
				log.Panic(copyErr)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return list
}

func Copy(path, dest, filename string) (err error) {
	if _, statErr := os.Stat(path); !os.IsNotExist(statErr) {
		fmt.Println(path)
		from, fromErr := os.Open(path)
		if fromErr != nil {
			log.Panic(fromErr)
		}
		defer from.Close()

		to, toErr := os.Create(dest + "/" + filename)

		if toErr != nil {
			log.Panic(toErr)
		}
		defer to.Close()

		if _, copyErr := io.Copy(to, from); copyErr != nil {
			log.Panic(copyErr)
			to.Close()
			return
		}
		syncErr := from.Sync()
		if syncErr != nil {
			log.Panic(syncErr)
		}
		return nil
	} else {
		return statErr
	}
}