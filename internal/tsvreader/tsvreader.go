package tsvreader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rlcache/internal/constants"
	"strings"
)

type TSVReaderConfig struct {
	NamesMap map[int]string
}

type TSVString map[string]string

var DefaultTSVReaderConfig = TSVReaderConfig{
	NamesMap: map[int]string{
		0: "timestamp",
		1: "data_id",
		2: "size",
		3: "read_bytes",
	},
}

type TSVReader struct {
	NamesMap map[int]string
}

func NewTSVReader(config *TSVReaderConfig) *TSVReader {
	return &TSVReader{
		NamesMap: config.NamesMap,
	}
}

func (r *TSVReader) Parse(filename string) (result []TSVString, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	linesRead := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		var newLine TSVString = make(map[string]string)
		for i, val := range splitLine {
			newLine[r.NamesMap[i]] = val
		}
		result = append(result, newLine)
		linesRead += 1
		if linesRead % 10000 == 0 {
			fmt.Printf("\rAt %s lines read %s", filename, constants.MakeBoldInt(linesRead))
		}
	}
	fmt.Printf("\rAt %s lines read %s\n", filename, constants.MakeBoldInt(linesRead))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}
