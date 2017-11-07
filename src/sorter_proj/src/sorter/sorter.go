package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
)

// 手抄代码
// flag, io, buffio, strconv 文档

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(filename string) (values []int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file:", filename)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("line is too long to parse.")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("failed to create the output file:", filename)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile name:", *infile)
		fmt.Println("outfile name", *outfile)
		fmt.Println("sorter name:", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
