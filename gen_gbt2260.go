package gbt2260

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CreateGBT2260Table() {
	file, err := os.Open("./data/GBT2260-201804.csv")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var line string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break //文件读完了就结束
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		code := record[0]
		name := record[1]
		line = line + "{\"" + code + "\",\"" + name + "\"},"
	}
	content := "package gbt2260\nvar gbt2260Table = [][]string{" + line + "}"
	ioutil.WriteFile("gbt2260_table.go", []byte(content), 0644)
}
