package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"rlcache/internal/constants"
	"rlcache/internal/tsvreader"
)

func main() {
	dirname := "/Users/durrdurr/Downloads/219.76.10.215_logs/"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	reader := tsvreader.NewTSVReader(&tsvreader.DefaultTSVReaderConfig)

	for i, f := range files {
		_, err := reader.Parse(path.Join(dirname, f.Name()))
		if err != nil {
			return
		}
		fmt.Printf("Done %s out of %s files\n", constants.MakeBoldInt(i), constants.MakeBoldInt(len(files)))
	}

	//for _, line := range lines {
	//	fmt.Println(line)
	//}
}
