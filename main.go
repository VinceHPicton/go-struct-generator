package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go-struct-generator/buildstructs"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var file string

	flag.StringVar(&file, "file", "", "Name of the input json file for the program to consume")

	flag.Parse()

	jsonBytes := readFile(file)

	var fieldSlice []buildstructs.StructField

	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&fieldSlice)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// lesson: range makes a copy of values if you use i, v -> v is a copy not a reference
	for i := range fieldSlice {
		fieldSlice[i].DataType = "struct"
	}

	fmt.Println(fieldSlice)

	// err := json.Unmarshal(jsonBytes, &fieldSlice)
	// if err != nil {
	// 	log.Fatalf("Failed to unmarshal JSON")
	// }

	structDefinitions := buildstructs.CreateStructsForFieldSlice(fieldSlice)

	printGoStructs(structDefinitions)

	writeGoStructsToOutFile(structDefinitions)

}

func writeGoStructsToOutFile(structDefinitions []string) {
	file, err := os.Create("out.go")
	if err != nil {
		log.Fatal("Failed to create out file")
	}
	defer file.Close()

	for _, s := range structDefinitions {

		_, err := file.WriteString(s)
		if err != nil {
			log.Fatalf("Failed to write line: %v", s)
		}
	}
}

func printGoStructs(fieldSlice []string) {

	for _, s := range fieldSlice {
		fmt.Print(s)
	}
}

func readFile(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", fileName)
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", fileName)
	}
	defer jsonFile.Close()

	return jsonBytes
}
