package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("1 argument required")
		os.Exit(1)
	}

	file := os.Args[1]
	_, err := os.Open(file)
	if err == nil {
		fmt.Println("file already exist")
		os.Exit(1)
	}

	output, err := os.Create(file)
	defer output.Close()

	if err != nil {
		fmt.Println("error craeting file")
		os.Exit(1)
	}

	inputdata := [][]string{{"A", "B", "C."}, {"D", "E", "F."}, {"G", "H", "I."}}
	writer := csv.NewWriter(output)

	for _, record := range inputdata {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("error writing to file")
			os.Exit(-1)
		}
	}
	writer.Flush()

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, rec := range allRecords {
		fmt.Printf("%s:%s:%s\n", rec[0], rec[1], rec[2])
	}

}
